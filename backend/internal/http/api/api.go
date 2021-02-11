package api

import (
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/sniddunc/refractor/pkg/config"
	"github.com/sniddunc/refractor/pkg/jwt"
	"github.com/sniddunc/refractor/pkg/log"
	"github.com/sniddunc/refractor/refractor"
	"net/http"
	"net/url"
	"os"
)

type API struct {
	echo *echo.Echo
	port string
	log  log.Logger
	*Handlers
}

// Handlers holds the handlers for the various application domains
type Handlers struct {
	AuthHandler   refractor.AuthHandler
	UserHandler   refractor.UserHandler
	GameHandler   refractor.GameHandler
	ServerHandler refractor.ServerHandler
}

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
	Payload interface{} `json:"payload,omitempty"`
}

func NewAPI(handlers *Handlers, port string, logger log.Logger) *API {
	echoApp := echo.New()

	api := &API{
		echo:     echoApp,
		port:     port,
		log:      logger,
		Handlers: handlers,
	}

	api.setupRoutes()

	return api
}

func (api *API) setupRoutes() {
	// API routing group
	apiGroup := api.echo.Group("/api/v1")

	jwtSecret := os.Getenv("JWT_SECRET")

	// Create JWT Middleware
	jwtMiddleware := echoMiddleware.JWTWithConfig(echoMiddleware.JWTConfig{
		Skipper:       echoMiddleware.DefaultSkipper,
		SigningKey:    []byte(jwtSecret),
		SigningMethod: echoMiddleware.AlgorithmHS256,
		TokenLookup:   "header:" + echo.HeaderAuthorization,
		AuthScheme:    "Bearer",
		Claims:        &jwt.Claims{},
	})

	// Auth endpoints
	authGroup := apiGroup.Group("/auth")
	authGroup.POST("/login", api.AuthHandler.LogInUser)
	authGroup.POST("/refresh", api.AuthHandler.RefreshUser)
	authGroup.GET("/check", api.AuthHandler.CheckAuth)

	// User endpoints
	userGroup := apiGroup.Group("/users", jwtMiddleware, AttachClaims())
	userGroup.GET("/me", api.UserHandler.GetOwnUserInfo)
	userGroup.POST("/changepassword", api.UserHandler.ChangeUserPassword)

	// Game endpoints
	gameGroup := apiGroup.Group("/games", jwtMiddleware, AttachClaims())
	gameGroup.GET("/", api.GameHandler.GetAllGames)

	// Server endpoints
	serverGroup := apiGroup.Group("/servers", jwtMiddleware, AttachClaims())
	serverGroup.POST("/", api.ServerHandler.CreateServer, api.RequireAccessLevel(config.AL_ADMIN))
}

func (api *API) ListenAndServe() error {
	return api.echo.Start(api.port)
}

type Validable interface {
	Validate() (bool, url.Values)
}

// ValidateRequest is a helper method to map out and validate request bodies quickly.
// If errors are found, it automatically sends them back to the user.
func ValidateRequest(body Validable, c echo.Context) bool {
	if err := c.Bind(body); err != nil {
		_ = c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: "Something went wrong. Please try again later.",
		})
	}

	// Validate the contents of the body.
	if ok, validationErrors := body.Validate(); !ok || len(validationErrors) > 0 {
		_ = c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Errors:  validationErrors,
		})

		return false
	}

	return true
}
