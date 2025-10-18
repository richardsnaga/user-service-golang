package seed

import (
	"log"
	"time"
	"user-service/internal/domain/user"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// RunUserSeeder menjalankan seeding untuk tabel users
func RunUserSeeder(db *gorm.DB) error {
	log.Println("🚀 Starting user seeder...")

	users := []user.User{
		{
			Name:      "Admin User",
			Email:     "admin@example.com",
			Password:  hashPassword("admin123"),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Name:      "John Doe",
			Email:     "john@example.com",
			Password:  hashPassword("john123"),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Name:      "Jane Smith",
			Email:     "jane@example.com",
			Password:  hashPassword("jane123"),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	for _, u := range users {
		var existing user.User
		err := db.Where("email = ?", u.Email).First(&existing).Error
		if err == gorm.ErrRecordNotFound {
			if err := db.Create(&u).Error; err != nil {
				return err
			}
			log.Printf("✅ User %s created.\n", u.Email)
		} else {
			log.Printf("⚠️ User %s already exists, skipping.\n", u.Email)
		}
	}

	log.Println("🎉 User seeder completed successfully.")
	return nil
}

func hashPassword(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash)
}
