package main

import (
	"flag"
	"log"
	"user-service/config"
	"user-service/db/seed"
	"user-service/internal/delivery/http"
	"user-service/internal/delivery/http/handler"
	"user-service/internal/domain/repository"
	"user-service/internal/usecase"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Flag untuk menentukan mode
	seedFlag := flag.Bool("seed", false, "Run database seeder")
	flag.Parse()

	cfg := config.Load()

	// Koneksi ke database
	db, err := gorm.Open(postgres.Open(cfg.DatabaseURL), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ failed to connect database: ", err)
	}

	// Jalankan seeder bila ada flag --seed
	if *seedFlag {
		log.Println("🌱 Running database seeder...")
		if err := seed.RunUserSeeder(db); err != nil {
			log.Fatal("❌ failed to seed database: ", err)
		}
		log.Println("✅ Seeding completed.")
		return
	}

	// Normal mode (run server)
	userRepo := repository.NewUserRepository(db)
	userUC := usecase.NewUserUsecase(userRepo)
	userHandler := handler.NewUserHandler(userUC)

	r := http.SetupRouter(userHandler)
	log.Println("🚀 Server running at port", cfg.AppPort)
	r.Run(":" + cfg.AppPort)
}
