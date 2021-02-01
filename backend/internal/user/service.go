package user

import (
	"github.com/sniddunc/refractor/internal/params"
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

func (s *userService) GetUserInfo(id int64) (*refractor.User, *refractor.ServiceResponse) {
	panic("implement me")
}
