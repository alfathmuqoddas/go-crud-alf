package main

import (
	"github.com/alfathmuqoddas/go-crud-alf/initializers"
	"github.com/alfathmuqoddas/go-crud-alf/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
  initializers.DB.AutoMigrate(&models.Post{})
}