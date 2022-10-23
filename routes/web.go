package routes

import (
	"false.kr/Monitor-Web/controllers"
	"github.com/gofiber/fiber"
	"github.com/gofiber/template/html"
)

func Router() *fiber.App {
	app := fiber.New(fiber.Config{
		Views: html.New("./_html", ".html"),
	})

	app.Get("/", controllers.Index)
	app.Get("/index", controllers.Index)
	app.Get("/additional", controllers.Additional)
	app.Post("/additionaldata", controllers.AdditionalData)
	app.Get("/modify/:idx", controllers.Modify)
	app.Post("/modifydata", controllers.ModifyData)
	app.Get("/checker/:idx", controllers.Checker)
	app.Get("/delete/:idx", controllers.DeleteCheck)

	app.Get("/groupmanage", controllers.GroupManage)
	app.Get("/groupadd", controllers.GroupAdd)
	app.Post("/groupadddata", controllers.GroupAddData)
	app.Get("/groupcheck/:name", controllers.GroupCheck)
	app.Get("/groupdelete/:name", controllers.GroupDelete)
	app.Get("/groupmodify/:name", controllers.GroupModify)
	app.Post("/groupmoddata", controllers.GroupModData)

	app.Static("/", "./_css/")
	app.Static("/", "./_js/")
	return app
}
