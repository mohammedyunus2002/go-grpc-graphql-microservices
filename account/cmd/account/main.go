package main

import (
	"log"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
	"github.com/mohammedyunus2002/go-grpc-graphql-microservices/account"
	"github.com/tinrab/retry"
)

type Config struct {
	DatabaseURL string `envconfig:"DATABASE_URL"`
}

func main() {
	// Load environment variables from .env
	err := godotenv.Load("cmd/account/.env") // Adjust path if necessary
	if err != nil {
		log.Println("No .env file found, relying on system environment variables")
	}
	var cfg Config
	err = envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	var r account.Repository
	retry.ForeverSleep(2*time.Second, func(_ int) (err error) {
		r, err = account.NewPostgresRepository(cfg.DatabaseURL)
		if err != nil {
			log.Println(err)
		}
		return
	})
	defer r.Close()
	log.Println("Listening on port 8080...")
	s := account.NewService(r)
	log.Fatal(account.ListenGRPC(s, 8080))
}
