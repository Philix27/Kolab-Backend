package app

import (
	"harmony/middleware"
	"harmony/services/announcement"
	"harmony/services/auth"
	"harmony/services/chat"
	"harmony/services/organization"
	"harmony/services/task_manager"
	"harmony/services/team"
	"harmony/services/user"
	"harmony/services/wiki"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"gorm.io/gorm"
)


type AppState struct {
    DB  *gorm.DB
}

func (a *AppState) SetupRoutes(app *fiber.App)  {
    api := app.Group("/api", middleware.ApiHandler) // /api
	v1 := api.Group("/v1", middleware.Version)      // /api/v1
	v1.Route("/announcement", announcement.RoutesHandler)
	v1.Route("/auth", auth.RoutesHandler)
	v1.Route("/chat", chat.RoutesHandler)
	v1.Route("/organization", organization.RoutesHandler)
	v1.Route("/task_manager", task_manager.RoutesHandler)
	v1.Route("/team", team.RoutesHandler)
	v1.Route("/user", user.RoutesHandler)
	v1.Route("/wiki", wiki.RoutesHandler)

	app.Get("/swagger/*", swagger.HandlerDefault) // default
    // app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
	// 	URL: "http://example.com/doc.json",
	// 	DeepLinking: false,
	// 	// Expand ("list") or Collapse ("none") tag groups by default
	// 	DocExpansion: "none",
	// 	// Prefill OAuth ClientId on Authorize popup
	// 	OAuth: &swagger.OAuthConfig{
	// 		AppName:  "OAuth Provider",
	// 		ClientId: "21bb4edc-05a7-4afc-86f1-2e151e4ba6e2",
	// 	},
	// 	// Ability to change OAuth2 redirect uri location
	// 	OAuth2RedirectUrl: "http://localhost:8080/swagger/oauth2-redirect.html",
	// }))
}