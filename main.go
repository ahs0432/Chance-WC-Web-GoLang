package main

import (
	"false.kr/Chance-WC-Web-GoLang/database"
	"false.kr/Chance-WC-Web-GoLang/files"
	"false.kr/Chance-WC-Web-GoLang/routes"
)

func main() {
	app := routes.Router()
	files.Init()
	database.Init(files.Config.DatabaseInfo.Host, files.Config.DatabaseInfo.Port, files.Config.DatabaseInfo.Protocol, files.Config.DatabaseInfo.User, files.Config.DatabaseInfo.Password, files.Config.DatabaseInfo.Name)
	app.Listen(":" + files.Config.Port)
}
