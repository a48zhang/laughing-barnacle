package main

import (
	"log"
	"main/model/db"
	"main/router"
	"os"
)

func main() {
	db.Connect()
	//qiniuStorage.Connect()
	r := router.NewRouter()
	port := os.Getenv("SH_PORT")
	if port == "" {
		port = "8080"
	}
	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
