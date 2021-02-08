package user

import (
	"github.com/sniddunc/refractor/internal/params"
	"github.com/sniddunc/refractor/pkg/config"
	"github.com/sniddunc/refractor/pkg/log"
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
		Email:    body.Email,
		Username: body.Username,
		Password: string(hashAndSalt),
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
		AccessLevel:         foundUser.AccessLevel,
		NeedsPasswordChange: foundUser.NeedsPasswordChange,
	}

	return userInfo, &refractor.ServiceResponse{
		Success:    true,
		StatusCode: http.StatusOK,
		Message:    "User info retrieved",
	}
}

// SetUserAccessLevel sets a user's access level. The user kicking off this interaction must be at least an admin, and
// must have a higher access level than the user whose access level they're updating.
func (s *userService) SetUserAccessLevel(body params.SetUserAccessLevelParams) (*refractor.User, *refractor.ServiceResponse) {
	// Make sure the setter user is an admin or higher
	if body.UserMeta.AccessLevel < config.AL_ADMIN || body.AccessLevel >= body.UserMeta.AccessLevel {
		s.log.Warn("Non-admin user with ID: %d and Access Level: %d tried to set the access level of user ID: %d to: %d",
			body.UserMeta.UserID, body.UserMeta.AccessLevel, body.UserID, body.AccessLevel)

		return nil, &refractor.ServiceResponse{
			Success:    false,
			StatusCode: http.StatusBadRequest,
			Message:    "You do not have permission to set the access level of this user",
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

	// Ensure that the updating user has a higher access level than the original user
	if body.UserMeta.AccessLevel <= foundUser.AccessLevel {
		s.log.Warn("User with ID: %d tried to set the access level of user ID: %d without having a higher access level",
			body.UserMeta.UserID, body.UserID)
		return nil, &refractor.ServiceResponse{
			Success:    false,
			StatusCode: http.StatusBadRequest,
			Message:    "You do not have permission to set the access level of this user",
		}
	}

	// Update the access level of the target user
	args := refractor.UpdateArgs{
		"AccessLevel": body.AccessLevel,
	}

	updatedUser, err := s.repo.Update(body.UserID, args)
	if err != nil {
		s.log.Error("Could not set the new access level of %d for user with ID: %d. Error: %v", body.AccessLevel, body.UserID, err)
		return nil, refractor.InternalErrorResponse
	}

	return updatedUser, &refractor.ServiceResponse{
		Success:    true,
		StatusCode: http.StatusOK,
		Message:    "Access level set. Any new access rights will come into effect next time the user logs in",
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

	return updatedUser, &refractor.ServiceResponse{
		Success:    true,
		StatusCode: http.StatusOK,
		Message:    "Password changed",
	}
}
