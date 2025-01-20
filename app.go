package main

import (
	"embed"

	"github.com/dimasaryo/fullstack-bootcamp-paas/app/controller"
	"github.com/dimasaryo/fullstack-bootcamp-paas/app/handler"
	"github.com/dimasaryo/fullstack-bootcamp-paas/app/model"
	"github.com/dimasaryo/fullstack-bootcamp-paas/config"
	_ "github.com/jackc/pgx/v4/stdlib"
)

//go:embed app/views/*
var Resources embed.FS

func main() {
	db := config.NewDB()
	conn, err := db.Connect()
	if err != nil {
		panic(err)
	}

	conn.AutoMigrate(&model.Student{})
	studentHandle := handler.NewStudentHandler(conn)

	mainAPI := controller.NewAPI(studentHandle, Resources)
	mainAPI.Start()
}
