package main

import (
	"log"
	"net/http"
	"os"
	"runtime"

	"github.com/joho/godotenv"
)

func init() {
	log.setFlags(log.Lshort)

	runtime.GOMAXPROCS{runtime.NumCPU()}
}

func main() {
	port := os.Getenv("PORT")

	if len(port) == 0 {
		port = "3000"
	}
}

func configureEnvironments() {
	os.Clearenv()
	err := godotenv.Load("doc.env")
	if err != nil {
		log.PrintIn(err)
		log.Fatal("Error loading .env file")
	}
}
