package auth

import (
	"github.com/sniddunc/refractor/internal/params"
	"github.com/sniddunc/refractor/pkg/config"
	"github.com/sniddunc/refractor/pkg/jwt"
	"github.com/sniddunc/refractor/pkg/log"
	"github.com/sniddunc/refractor/refractor"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type authService struct {
	repo      refractor.UserRepository
	log       log.Logger
	jwtSecret string
}

func NewAuthService(userRepo refractor.UserRepository, logger log.Logger, jwtSecret string) refractor.AuthService {
	return &authService{
		repo:      userRepo,
		log:       logger,
		jwtSecret: jwtSecret,
	}
}

func (s *authService) LogInUser(body params.LoginParams) (*refractor.TokenPair, *refractor.ServiceResponse) {
	// Check if an account with the provided username exists
	args := refractor.FindArgs{
		"Username": body.Username,
	}

	foundUser, err := s.repo.FindOne(args)
	if err != nil {
		if err == refractor.ErrNotFound {
			return nil, &refractor.ServiceResponse{
				Success:    false,
				StatusCode: http.StatusBadRequest,
				Message:    config.MessageInvalidCredentials,
			}
		}

		s.log.Error("Could not FindOne user. Error: %v", err)
		return nil, refractor.InternalErrorResponse
	}

	// Make sure user account is activated. If it isn't, we don't want to let them log in.
	if !foundUser.Activated {
		s.log.Warn("Attempted login of deactivated user account. ID: %d Username: %s", foundUser.UserID, foundUser.Username)
		return nil, &refractor.ServiceResponse{
			Success:    false,
			StatusCode: http.StatusBadRequest,
			Message:    config.MessageDeactivatedAccount,
		}
	}

	// Compare password hashes
	hashedPassword := []byte(foundUser.Password)

	if err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(body.Password)); err != nil {
		s.log.Info("Failed login attempt for user: %s. Error: %v", foundUser.Username, err)
		return nil, &refractor.ServiceResponse{
			Success:    false,
			StatusCode: http.StatusBadRequest,
			Message:    config.MessageInvalidCredentials,
		}
	}

	// Generate and return user JWTs
	tokenPair, err := getAuthRefreshTokenPair(foundUser, s.jwtSecret)
	if err != nil {
		s.log.Error("Could not generate JWT pair. Error: %v", err)
		return nil, &refractor.ServiceResponse{
			Success:    false,
			StatusCode: http.StatusInternalServerError,
			Message:    config.MessageInternalError,
		}
	}

	s.log.Info("User %s (%d) logged in", foundUser.Username, foundUser.UserID)

	// All ok. Send back success message and tokens
	return &refractor.TokenPair{
			AuthToken:    tokenPair.AuthToken,
			RefreshToken: tokenPair.RefreshToken,
		}, &refractor.ServiceResponse{
			Success:    true,
			StatusCode: http.StatusOK,
			Message:    "Successfully logged in",
		}
}

func (s *authService) RefreshUser(refreshToken string) (*refractor.TokenPair, *refractor.ServiceResponse) {
	// Read refresh token from cookie
	if refreshToken == "" {
		return nil, &refractor.ServiceResponse{
			Success:    true,
			StatusCode: http.StatusBadRequest,
			Message:    config.MessageUnableRefreshCreds,
		}
	}

	// Parse token within cookie
	claims, err := jwt.ExtractRefreshClaims(refreshToken, s.jwtSecret)
	if err != nil {
		return nil, &refractor.ServiceResponse{
			Success:    false,
			StatusCode: http.StatusBadRequest,
			Message:    config.MessageUnableRefreshCreds,
		}
	}

	// Retrieve claimed user to make sure they exist and that their account is still activated
	user, err := s.repo.FindByID(claims.UserID)
	if err != nil {
		if err == refractor.ErrNotFound {
			return nil, &refractor.ServiceResponse{
				Success:    false,
				StatusCode: http.StatusBadRequest,
				Message:    config.MessageUnableRefreshCreds,
			}
		}

		s.log.Error("Could not find user by ID: %d. Error: %v", claims.UserID, err)
		return nil, refractor.InternalErrorResponse
	}

	if !user.Activated {
		s.log.Info("Failed refresh attempt on deactivated user account: %s (%d)", user.Username, user.UserID)
		return nil, &refractor.ServiceResponse{
			Success:    false,
			StatusCode: http.StatusBadRequest,
			Message:    config.MessageUnableRefreshCreds,
		}
	}

	tokenPair, err := getAuthRefreshTokenPair(user, s.jwtSecret)
	if err != nil {
		s.log.Error("Could not generate token pair. Error: %v", err)
		return nil, refractor.InternalErrorResponse
	}

	s.log.Info("New auth and refresh tokens have been generated for user %s (%d) (refresh)", user.Username, user.UserID)

	// Send back response and token pair
	return &refractor.TokenPair{
			AuthToken:    tokenPair.AuthToken,
			RefreshToken: tokenPair.RefreshToken,
		}, &refractor.ServiceResponse{
			Success:    true,
			StatusCode: http.StatusOK,
		}
}

func getAuthRefreshTokenPair(user *refractor.User, jwtSecret string) (*jwt.TokenPair, error) {
	// Generate tokens
	jwtSecretStr := []byte(jwtSecret)

	authToken, refreshToken, err := jwt.GenerateAuthTokens(user, jwtSecretStr)
	if err != nil {
		return nil, err
	}

	return &jwt.TokenPair{
		AuthToken:    authToken,
		RefreshToken: refreshToken,
	}, nil
}
