package api

import (
	"encoding/json"
	"github.com/gorilla/schema"
	"github.com/labstack/echo/v4"
	"github.com/sniddunc/refractor/internal/auth"
	"github.com/sniddunc/refractor/internal/mock"
	"github.com/sniddunc/refractor/internal/params"
	"github.com/sniddunc/refractor/pkg/config"
	"github.com/sniddunc/refractor/pkg/jwt"
	"github.com/sniddunc/refractor/pkg/log"
	"github.com/sniddunc/refractor/refractor"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"
)

func Test_authHandler_LogInUser(t *testing.T) {
	encoder := schema.NewEncoder()
	logger, _ := log.NewLogger(true, false)
	echoApp := echo.New()
	const testJWTSecret = "djswg2%T2tst^!@%!bte72tSB^"

	type fields struct {
		user   *mock.MockUser
		secure bool
	}
	tests := []struct {
		name        string
		fields      fields
		body        params.LoginParams
		wantSuccess bool
	}{
		{
			name: "http.api.auth.loginuser.1",
			fields: fields{
				user: &mock.MockUser{
					UnhashedPassword: "password",
					User: &refractor.User{
						UserID:              1,
						Email:               "test@test.com",
						Username:            "testuser.1",
						Password:            mock.HashPassword("password"),
						AccessLevel:         config.AL_USER,
						Activated:           true,
						NeedsPasswordChange: false,
					},
				},
				secure: false,
			},
			body: params.LoginParams{
				Username: "testuser.1",
				Password: "password",
			},
			wantSuccess: true,
		},
		{
			name: "http.api.auth.loginuser.2",
			fields: fields{
				user: &mock.MockUser{
					UnhashedPassword: "password",
					User: &refractor.User{
						UserID:              1,
						Email:               "test@test.com",
						Username:            "testuser.1",
						Password:            mock.HashPassword("password"),
						AccessLevel:         config.AL_USER,
						Activated:           true,
						NeedsPasswordChange: false,
					},
				},
				secure: false,
			},
			body: params.LoginParams{
				Username: "testuser.1",
				Password: "invalidpassword",
			},
			wantSuccess: false,
		},
	}
	for _, tt := range tests {
		mockUserRepo := mock.NewMockUserRepository(map[int64]*mock.MockUser{
			1: tt.fields.user,
		})
		authService := auth.NewAuthService(mockUserRepo, logger, testJWTSecret)
		authHandler := NewAuthHandler(authService, tt.fields.secure)

		form := url.Values{}
		if err := encoder.Encode(tt.body, form); err != nil {
			t.Fatalf("Could not encode test form data for test %s. Error: %v", tt.name, err)
		}

		req := httptest.NewRequest(http.MethodPost, "/api/v1/auth/login", strings.NewReader(form.Encode()))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
		rec := httptest.NewRecorder()
		c := echoApp.NewContext(req, rec)

		t.Run(tt.name, func(t *testing.T) {
			assert.NoError(t, authHandler.LogInUser(c), "LogInUser returned an error")

			var response *Response
			assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &response), "Could not unmarshal response")

			// Verify claims of generated token if we're expect a successful login
			responsePayload := response.Payload
			if responsePayload == nil && tt.wantSuccess {
				t.Fatalf("Could not extract claims from JWT. response.Payload was ")
			}

			if tt.wantSuccess {
				claims, err := jwt.ExtractAuthClaims(response.Payload.(string), testJWTSecret)
				assert.NoError(t, err, "Claims could not be extracted")

				foundUser, err := mockUserRepo.FindByID(claims.UserID)
				assert.NoError(t, err, "Could not get claimed user from mock repo")

				assert.Equal(t, tt.fields.user.Username, foundUser.Username, "Found username did not match the expected username")
			}

			if tt.wantSuccess != response.Success {
				t.Fatalf("response.Success was the wrong value. Want = %v Got = %v", tt.wantSuccess, response.Success)
			}
		})
	}
}

func Test_authHandler_RefreshUser(t *testing.T) {
	logger, _ := log.NewLogger(true, false)
	echoApp := echo.New()
	const testJWTSecret = "djswg2%T2tst^!@%!bte72tSB^"

	type fields struct {
		user   *mock.MockUser
		secure bool
	}
	tests := []struct {
		name        string
		fields      fields
		wantSuccess bool
	}{
		{
			name: "http.api.auth.refreshuser.1",
			fields: fields{
				user: &mock.MockUser{
					UnhashedPassword: "password",
					User: &refractor.User{
						UserID:              1,
						Email:               "test@test.com",
						Username:            "testuser.1",
						Password:            mock.HashPassword("password"),
						AccessLevel:         config.AL_USER,
						Activated:           true,
						NeedsPasswordChange: false,
					},
				},
				secure: false,
			},
			wantSuccess: true,
		},
		{
			name: "http.api.auth.refreshuser.2",
			fields: fields{
				user: &mock.MockUser{
					UnhashedPassword: "password",
					User: &refractor.User{
						UserID:              1,
						Email:               "test@test.com",
						Username:            "testuser.1",
						Password:            mock.HashPassword("password"),
						AccessLevel:         config.AL_USER,
						Activated:           false,
						NeedsPasswordChange: false,
					},
				},
				secure: false,
			},
			wantSuccess: false,
		},
	}
	for _, tt := range tests {
		mockUserRepo := mock.NewMockUserRepository(map[int64]*mock.MockUser{
			1: tt.fields.user,
		})
		authService := auth.NewAuthService(mockUserRepo, logger, testJWTSecret)
		authHandler := NewAuthHandler(authService, tt.fields.secure)

		req := httptest.NewRequest(http.MethodPost, "/api/v1/auth/refresh", strings.NewReader(""))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
		rec := httptest.NewRecorder()

		// Generate baseline refresh token and set it as a cookie
		_, refreshToken, err := jwt.GenerateAuthTokens(tt.fields.user.User, []byte(testJWTSecret))
		if err != nil {
			t.Fatalf("Could not generate baseline refreshToken for test %s. Error: %v", tt.name, err)
		}

		// Set refresh_token cookie
		req.AddCookie(&http.Cookie{
			Name:     "refresh_token",
			Value:    refreshToken,
			Expires:  time.Now().Add(14 * (24 * time.Hour)),
			HttpOnly: true,
		})

		c := echoApp.NewContext(req, rec)

		t.Run(tt.name, func(t *testing.T) {
			assert.NoError(t, authHandler.RefreshUser(c), "RefreshUser returned an error")

			var response *Response
			assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &response), "Could not unmarshal response")

			// Verify claims of generated token if we're expect a successful login
			responsePayload := response.Payload
			if responsePayload == nil && tt.wantSuccess {
				t.Fatalf("Could not extract claims from JWT. response.Payload was nil")
			}

			if tt.wantSuccess != response.Success {
				t.Fatalf("response.Success was the wrong value. Want = %v Got = %v", tt.wantSuccess, response.Success)
			}

			if tt.wantSuccess {
				// Extract cookie set by the RefreshUser handler
				cookies := rec.Result().Cookies()
				var newRefreshToken string

				for _, cookie := range cookies {
					if cookie.Name != "refresh_token" {
						continue
					}

					newRefreshToken = cookie.Value
				}

				if newRefreshToken == "" {
					assert.Failf(t, "No refresh_token cookie was set for test %s", tt.name)
				}

				// Extract claims from the new refresh token cookie
				newRefreshClaims, err := jwt.ExtractRefreshClaims(newRefreshToken, testJWTSecret)
				if err != nil || newRefreshClaims == nil {
					assert.Failf(t, "Could not extract claims from new refresh_token cookie for test %s", tt.name)
				}

				// Compare new refresh claims to user info
				if newRefreshClaims != nil && newRefreshClaims.UserID != tt.fields.user.UserID {
					assert.Failf(t, "New refresh token claims UserID do not match UserID provided for test %s. Want = %v Got = %v", tt.name, tt.fields.user.UserID, newRefreshClaims.UserID)
				}
			}
		})
	}
}
