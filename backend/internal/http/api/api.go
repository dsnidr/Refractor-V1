package api

import (
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/sniddunc/refractor/pkg/jwt"
	"github.com/sniddunc/refractor/pkg/log"
	"github.com/sniddunc/refractor/pkg/perms"
	"github.com/sniddunc/refractor/refractor"
	"net/http"
	"net/url"
	"os"
)

type API struct {
	echo             *echo.Echo
	port             string
	log              log.Logger
	websocketService refractor.WebsocketService
	*Handlers
}

// Handlers holds the handlers for the various application domains
type Handlers struct {
	AuthHandler       refractor.AuthHandler
	UserHandler       refractor.UserHandler
	ServerHandler     refractor.ServerHandler
	PlayerHandler     refractor.PlayerHandler
	GameServerHandler refractor.GameServerHandler
	InfractionHandler refractor.InfractionHandler
	SummaryHandler    refractor.SummaryHandler
	SearchHandler     refractor.SearchHandler
}

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
	Payload interface{} `json:"payload,omitempty"`
}

func NewAPI(handlers *Handlers, port string, logger log.Logger, websocketService refractor.WebsocketService) *API {
	echoApp := echo.New()

	api := &API{
		echo:             echoApp,
		port:             port,
		log:              logger,
		websocketService: websocketService,
		Handlers:         handlers,
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
	userGroup.GET("/all", api.UserHandler.GetAllUsers, api.RequirePerms(perms.FULL_ACCESS))
	userGroup.POST("/", api.UserHandler.CreateUser, api.RequirePerms(perms.FULL_ACCESS))
	userGroup.PATCH("/activate/:id", api.UserHandler.ActivateUser, api.RequirePerms(perms.FULL_ACCESS))
	userGroup.PATCH("/deactivate/:id", api.UserHandler.DeactivateUser, api.RequirePerms(perms.FULL_ACCESS))
	userGroup.PATCH("/setpassword", api.UserHandler.SetUserPassword, api.RequirePerms(perms.FULL_ACCESS))
	userGroup.PATCH("/forcepasswordchange/:id", api.UserHandler.ForcePasswordChange, api.RequirePerms(perms.FULL_ACCESS))
	userGroup.PATCH("/setpermissions", api.UserHandler.SetUserPermissions, api.RequirePerms(perms.SUPER_ADMIN))

	// Game endpoints
	gameGroup := apiGroup.Group("/gameservers", jwtMiddleware, AttachClaims())
	gameGroup.GET("/", api.GameServerHandler.GetAllGameServers)

	// Server endpoints
	serverGroup := apiGroup.Group("/servers", jwtMiddleware, AttachClaims())
	serverGroup.POST("/", api.ServerHandler.CreateServer, api.RequirePerms(perms.FULL_ACCESS))
	serverGroup.GET("/", api.ServerHandler.GetAllServers)
	serverGroup.GET("/data", api.ServerHandler.GetAllServerData)
	serverGroup.PATCH("/:id", api.ServerHandler.UpdateServer, api.RequirePerms(perms.FULL_ACCESS))
	serverGroup.DELETE("/:id", api.ServerHandler.DeleteServer, api.RequirePerms(perms.FULL_ACCESS))

	// Infraction endpoints
	infractionGroup := apiGroup.Group("/infractions", jwtMiddleware, AttachClaims())
	infractionGroup.POST("/warning", api.InfractionHandler.CreateWarning, api.RequirePerms(perms.LOG_WARNING))
	infractionGroup.POST("/mute", api.InfractionHandler.CreateMute, api.RequirePerms(perms.LOG_MUTE))
	infractionGroup.POST("/kick", api.InfractionHandler.CreateKick, api.RequirePerms(perms.LOG_KICK))
	infractionGroup.POST("/ban", api.InfractionHandler.CreateBan, api.RequirePerms(perms.LOG_BAN))
	infractionGroup.DELETE("/:id", api.InfractionHandler.DeleteInfraction, api.RequireOneOfPerms(perms.DELETE_OWN_INFRACTIONS, perms.DELETE_ANY_INFRACTION))
	infractionGroup.PATCH("/:id", api.InfractionHandler.UpdateInfraction, api.RequireOneOfPerms(perms.EDIT_OWN_INFRACTIONS, perms.EDIT_ANY_INFRACTION))
	infractionGroup.GET("/:id/warnings", api.InfractionHandler.GetPlayerInfractions(refractor.INFRACTION_TYPE_WARNING))
	infractionGroup.GET("/:id/mutes", api.InfractionHandler.GetPlayerInfractions(refractor.INFRACTION_TYPE_MUTE))
	infractionGroup.GET("/:id/kicks", api.InfractionHandler.GetPlayerInfractions(refractor.INFRACTION_TYPE_KICK))
	infractionGroup.GET("/:id/bans", api.InfractionHandler.GetPlayerInfractions(refractor.INFRACTION_TYPE_BAN))

	// Player endpoints
	playerGroup := apiGroup.Group("/players", jwtMiddleware, AttachClaims())
	playerGroup.GET("/recent", api.PlayerHandler.GetRecentPlayers)
	playerGroup.GET("/summary/:id", api.SummaryHandler.GetPlayerSummary)

	// Search endpoints
	searchGroup := apiGroup.Group("/search", jwtMiddleware, AttachClaims())
	searchGroup.POST("/players", api.SearchHandler.SearchPlayers)

	// Websocket endpoint
	api.echo.Any("/ws", api.websocketHandler)
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
