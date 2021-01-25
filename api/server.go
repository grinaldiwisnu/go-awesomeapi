package api

import (
	"awesomeProject/api/controllers"
	"awesomeProject/api/seed"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var s = controllers.Server{}

func Run() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	s.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
	seed.Load(s.DB)
	s.Run(":8080")
}
