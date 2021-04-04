package user

import (
	"fmt"
	"github.com/sniddunc/bitperms"
	"github.com/sniddunc/refractor/internal/params"
	"github.com/sniddunc/refractor/pkg/config"
	"github.com/sniddunc/refractor/pkg/log"
	"github.com/sniddunc/refractor/pkg/perms"
	"github.com/sniddunc/refractor/refractor"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"net/url"
)

type userService struct {
	repo refractor.UserRepository
	log  log.Logger
}

func NewUserService(userRepo refractor.UserRepository, logger log.Logger) refractor.UserService {
	return &userService{
		repo: userRepo,
		log:  logger,
	}
}

func (s *userService) CreateUser(body params.CreateUserParams) (*refractor.User, *refractor.ServiceResponse) {
	// Make sure username isn't already taken
	exists, err := s.repo.Exists(refractor.FindArgs{
		"Username": body.Username,
	})

	if err != nil {
		s.log.Error("Could not check existence of user. Error: %v", err)
		return nil, refractor.InternalErrorResponse
	}

	if exists {
		return nil, &refractor.ServiceResponse{
			Success:    false,
			StatusCode: http.StatusBadRequest,
			ValidationErrors: url.Values{
				"username": []string{"That username is already in use"},
			},
		}
	}

	// Make sure email isn't already taken
	exists, err = s.repo.Exists(refractor.FindArgs{
		"Email": body.Email,
	})
	if err != nil {
		s.log.Error("Could not check existence of user. Error: %v", err)
		return nil, refractor.InternalErrorResponse
	}

	if exists {
		return nil, &refractor.ServiceResponse{
			Success:    false,
			StatusCode: http.StatusBadRequest,
			ValidationErrors: url.Values{
				"email": []string{"That email is already in use"},
			},
		}
	}

	// If new user credentials are valid, hash the password
	hashAndSalt, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		s.log.Error("Could not generate hash and salt for new user's password. Error: %v", err)
		return nil, refractor.InternalErrorResponse
	}

	// Create the new user
	newUser := &refractor.User{
		Email:       body.Email,
		Username:    body.Username,
		Password:    string(hashAndSalt),
		Permissions: perms.DEFAULT_PERMS,
	}

	if err := s.repo.Create(newUser); err != nil {
		s.log.Error("Could not insert new user into repository. Error: %v", err)
		return nil, refractor.InternalErrorResponse
	}

	s.log.Info("A new user with the username: %s has been created", body.Username)

	return newUser, &refractor.ServiceResponse{
		Success:    true,
		StatusCode: http.StatusOK,
		Message:    "User created",
	}
}

func (s *userService) GetUserByID(id int64) (*refractor.User, *refractor.ServiceResponse) {
	foundUser, err := s.repo.FindByID(id)
	if err != nil {
		if err == refractor.ErrNotFound {
			return nil, &refractor.ServiceResponse{
				Success:    false,
				StatusCode: http.StatusBadRequest,
				Message:    config.MessageInvalidIDProvided,
			}
		}

		s.log.Error("Could not retrieve user by ID: %d. Error: %v", id, err)
		return nil, refractor.InternalErrorResponse
	}

	return foundUser, nil
}

func (s *userService) GetUserInfo(id int64) (*refractor.UserInfo, *refractor.ServiceResponse) {
	foundUser, err := s.repo.FindByID(id)
	if err != nil {
		if err == refractor.ErrNotFound {
			return nil, &refractor.ServiceResponse{
				Success:    false,
				StatusCode: http.StatusBadRequest,
				Message:    config.MessageInvalidIDProvided,
			}
		}

		s.log.Error("Could not retrieve user by ID: %d. Error: %v", id, err)
		return nil, refractor.InternalErrorResponse
	}

	userInfo := &refractor.UserInfo{
		ID:                  foundUser.UserID,
		Email:               foundUser.Email,
		Username:            foundUser.Username,
		Activated:           foundUser.Activated,
		Permissions:         foundUser.Permissions,
		NeedsPasswordChange: foundUser.NeedsPasswordChange,
	}

	return userInfo, &refractor.ServiceResponse{
		Success:    true,
		StatusCode: http.StatusOK,
		Message:    "User info retrieved",
	}
}

// SetUserPermissions sets a user's permissions. The user kicking off this interaction must be at least an admin, and
// must have a higher level of access than the user whose access level they're updating.
func (s *userService) SetUserPermissions(body params.SetUserPermissionsParams) (*refractor.User, *refractor.ServiceResponse) {
	setterPerms := bitperms.PermissionValue(body.UserMeta.Permissions)

	if !perms.UserHasFullAccess(setterPerms) {
		s.log.Warn("Non-admin user with ID: %d and permissions value: %d tried to set the permissions of user ID: %d to: %d",
			body.UserMeta.UserID, body.UserMeta.Permissions, body.UserID, body.Permissions)

		return nil, &refractor.ServiceResponse{
			Success:    false,
			StatusCode: http.StatusBadRequest,
			Message:    config.MessageNoPermission,
		}
	}

	// Find the target user in storage
	foundUser, err := s.repo.FindByID(body.UserID)
	if err != nil {
		if err == refractor.ErrNotFound {
			return nil, &refractor.ServiceResponse{
				Success:    false,
				StatusCode: http.StatusBadRequest,
				Message:    config.MessageInvalidIDProvided,
			}
		}

		s.log.Error("Could not get user by ID from repo. UserID: %d Error: %v", body.UserID, err)
		return nil, refractor.InternalErrorResponse
	}

	targetPerms := bitperms.PermissionValue(foundUser.Permissions)

	// Ensure that the updating user has a higher access level than the original user
	if !perms.HasHigherAccess(setterPerms, targetPerms) {
		s.log.Warn("User with ID: %d tried to set the permissions of user ID: %d without having a higher access level",
			body.UserMeta.UserID, body.UserID)
		return nil, &refractor.ServiceResponse{
			Success:    false,
			StatusCode: http.StatusBadRequest,
			Message:    config.MessageNoPermission,
		}
	}

	// Update the permissions value of the target user
	args := refractor.UpdateArgs{
		"Permissions": body.Permissions,
	}

	updatedUser, err := s.repo.Update(body.UserID, args)
	if err != nil {
		s.log.Error("Could not set the new permissions value of %d for user with ID: %d. Error: %v", body.Permissions, body.UserID, err)
		return nil, refractor.InternalErrorResponse
	}

	s.log.Info("User ID %d set the permissions for User %s (%d)", body.UserMeta.UserID, foundUser.Username, foundUser.UserID)

	return updatedUser, &refractor.ServiceResponse{
		Success:    true,
		StatusCode: http.StatusOK,
		Message:    "Permissions set. Any new access rights will come into effect next time the reloads.",
	}
}

func (s *userService) ForceUserPasswordChange(id int64, userMeta *params.UserMeta) *refractor.ServiceResponse {
	user1Perms := bitperms.PermissionValue(userMeta.Permissions)

	if !perms.UserHasFullAccess(user1Perms) {
		return &refractor.ServiceResponse{
			Success:    true,
			StatusCode: http.StatusBadRequest,
			Message:    config.MessageNoPermission,
		}
	}

	foundUser, err := s.repo.FindByID(id)
	if err != nil {
		if err == refractor.ErrNotFound {
			return &refractor.ServiceResponse{
				Success:    false,
				StatusCode: http.StatusBadRequest,
				Message:    config.MessageInvalidIDProvided,
			}
		}

		s.log.Error("Could not find user by ID: %d. Error: %v", id, err)
		return refractor.InternalErrorResponse
	}

	user2Perms := bitperms.PermissionValue(foundUser.Permissions)

	if !perms.HasHigherAccess(user1Perms, user2Perms) {
		return &refractor.ServiceResponse{
			Success:    false,
			StatusCode: http.StatusBadRequest,
			Message:    config.MessageNoPermission,
		}
	}

	// Update the user's password flag
	_, err = s.repo.Update(foundUser.UserID, refractor.UpdateArgs{
		"NeedsPasswordChange": true,
	})
	if err != nil {
		s.log.Error("Could not update user ID: %d in repository. Error: %v", foundUser.UserID, err)
		return refractor.InternalErrorResponse
	}

	return &refractor.ServiceResponse{
		Success:    true,
		StatusCode: http.StatusOK,
		Message:    "User will be forced to change their password on their next visit.",
	}
}

func (s *userService) ChangeUserPassword(id int64, body params.ChangeUserPassword) (*refractor.User, *refractor.ServiceResponse) {
	foundUser, err := s.repo.FindByID(id)
	if err != nil {
		if err == refractor.ErrNotFound {
			return nil, &refractor.ServiceResponse{
				Success:    false,
				StatusCode: http.StatusBadRequest,
				Message:    config.MessageInvalidIDProvided,
			}
		}

		s.log.Error("Could not retrieve user by ID: %d. Error: %v", id, err)
		return nil, refractor.InternalErrorResponse
	}

	// Make sure the current password provided by the user matches their current password
	hashedPassword := []byte(foundUser.Password)

	if err := bcrypt.CompareHashAndPassword(hashedPassword, []byte(body.CurrentPassword)); err != nil {
		return nil, &refractor.ServiceResponse{
			Success:    false,
			StatusCode: http.StatusBadRequest,
			ValidationErrors: url.Values{
				"currentPassword": []string{"Incorrect password"},
			},
		}
	}

	// Make sure their new password doesn't match their current password
	if err := bcrypt.CompareHashAndPassword(hashedPassword, []byte(body.NewPassword)); err == nil {
		return nil, &refractor.ServiceResponse{
			Success:    false,
			StatusCode: http.StatusBadRequest,
			ValidationErrors: url.Values{
				"newPassword": []string{"You can't re-use your current password"},
			},
		}
	}

	// Hash the new password
	hashAndSalt, err := bcrypt.GenerateFromPassword([]byte(body.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		s.log.Error("Could not new password hash for user with ID: %d. Error: %v", foundUser.UserID, err)
		return nil, refractor.InternalErrorResponse
	}

	// Update the user in the repository and make sure their NeedsPasswordChange flag is set to false
	args := refractor.UpdateArgs{
		"Password":            string(hashAndSalt),
		"NeedsPasswordChange": false,
	}

	updatedUser, err := s.repo.Update(foundUser.UserID, args)
	if err != nil {
		s.log.Error("Could not update user ID: %d in repository. Error: %v", foundUser.UserID, err)
		return nil, refractor.InternalErrorResponse
	}

	s.log.Info("User %s (%d) has changed their password", foundUser.Username, foundUser.UserID)

	return updatedUser, &refractor.ServiceResponse{
		Success:    true,
		StatusCode: http.StatusOK,
		Message:    "Password changed",
	}
}

func (s *userService) GetAllUsers() ([]*refractor.User, *refractor.ServiceResponse) {
	users, err := s.repo.FindAll()
	if err != nil {
		if err == refractor.ErrNotFound {
			return []*refractor.User{}, &refractor.ServiceResponse{
				Success:    true,
				StatusCode: http.StatusOK,
				Message:    "Fetched 0 users",
			}
		}

		s.log.Error("Could not FindAll users in storage. Error: %v", err)
		return nil, refractor.InternalErrorResponse
	}

	return users, &refractor.ServiceResponse{
		Success:    true,
		StatusCode: http.StatusOK,
		Message:    fmt.Sprintf("Fetched %d users", len(users)),
	}
}

func (s *userService) UpdateUser(id int64, args refractor.UpdateArgs) (*refractor.User, *refractor.ServiceResponse) {
	updatedUser, err := s.repo.Update(id, args)
	if err != nil {
		if err == refractor.ErrNotFound {
			return nil, &refractor.ServiceResponse{
				Success:    false,
				StatusCode: http.StatusBadRequest,
				Message:    config.MessageInvalidIDProvided,
			}
		}

		s.log.Error("Could not update user with ID: %d. Error: %v", id, err)
		return nil, refractor.InternalErrorResponse
	}

	return updatedUser, &refractor.ServiceResponse{
		Success:    true,
		StatusCode: http.StatusOK,
		Message:    "User updated",
	}
}

func (s *userService) SetUserPassword(body params.SetUserPasswordParams) (*refractor.User, *refractor.ServiceResponse) {
	foundUser, err := s.repo.FindByID(body.UserID)
	if err != nil {
		if err == refractor.ErrNotFound {
			return nil, &refractor.ServiceResponse{
				Success:    false,
				StatusCode: http.StatusBadRequest,
				Message:    config.MessageInvalidIDProvided,
			}
		}

		s.log.Error("Could not get user by ID from repository. User ID: %d Error: %v", body.UserID, err)
		return nil, refractor.InternalErrorResponse
	}

	setterPerms := bitperms.PermissionValue(body.UserMeta.Permissions)
	targetPerms := bitperms.PermissionValue(foundUser.Permissions)

	// Ensure that the updating user has a level of access than the user that they're updating
	if !perms.HasHigherAccess(setterPerms, targetPerms) {
		s.log.Warn("User with ID: %d tried to set a new password on a user with ID: %d without having a higher level of access", body.UserMeta.UserID, body.UserID)
		return nil, &refractor.ServiceResponse{
			Success:    false,
			StatusCode: http.StatusBadRequest,
			Message:    "You do not have permission to set a new password for this user. This incident was recorded.",
		}
	}

	// Hash and salt the new password
	hashAndSalt, err := bcrypt.GenerateFromPassword([]byte(body.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		s.log.Error("Could not generate salt and hash for new password. Error: %v", err)
		return nil, refractor.InternalErrorResponse
	}

	// Update user to use the new password and set their NeedsPasswordChange flag to true
	args := refractor.UpdateArgs{
		"Password":            string(hashAndSalt),
		"NeedsPasswordChange": true,
	}

	updatedUser, err := s.repo.Update(foundUser.UserID, args)
	if err != nil {
		s.log.Error("Could not set new password for user with ID: %d. Error: %v", body.UserID, err)
		return nil, refractor.InternalErrorResponse
	}

	s.log.Info("User ID %d set a new password for User %s (%d)", body.UserMeta.UserID, foundUser.Username, foundUser.UserID)

	return updatedUser, &refractor.ServiceResponse{
		Success:    true,
		StatusCode: http.StatusOK,
		Message:    "New password set",
	}
}
