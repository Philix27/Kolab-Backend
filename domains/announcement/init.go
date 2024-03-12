package announcement

import (
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Setup(router fiber.Router, db *gorm.DB) {

	repo := InitializeRepository(db)

	svc := InitializeService(repo, validator.New())

	handler :=  InitRoutes(svc)

	router.Route("/announcement", handler.create)
	
}


// announce := router.Group("/")
	// announce.Post("/", handler.create).Name("AnnouncementCreate")
	// announce.Patch("/:id/expired", handler.update).Name("AnnouncementPatch")
	// announce.Get("/", handler.getAll).Name("AnnouncementGetAll")
	// announce.Get("/id", handler.getOne).Name("AnnouncementGetOne")
	// announce.Delete("/", handler.deleteOne).Name("AnnouncementDelete")
	// return GenRoutes(router  ,handler)

// func GenRoutes(router fiber.Router, handler Routes)    {
// 	announce := router.Group("/")
// 	announce.Post("/", handler.create).Name("AnnouncementCreate")
// 	announce.Patch("/:id/expired", handler.update).Name("AnnouncementPatch")
// 	announce.Get("/", handler.getAll).Name("AnnouncementGetAll")
// 	announce.Get("/id", handler.getOne).Name("AnnouncementGetOne")
// 	announce.Delete("/", handler.deleteOne).Name("AnnouncementDelete")
// }