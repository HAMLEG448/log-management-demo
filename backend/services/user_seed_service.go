package sevices

import (
	"fmt"
	"os"

	"log-management/backend/config"
	"log-management/backend/entity"

	"golang.org/x/crypto/bcrypt"
)

func SeedUsers() error {
	adminUsername := os.Getenv("ADMIN_USERNAME")
	adminPassword := os.Getenv("ADMIN_PASSWORD")

	viewerUsername := os.Getenv("VIEWER_USERNAME")
	viewerPassword := os.Getenv("VIEWER_PASSWORD")

	if err := createUserIfNotExists(
		adminUsername,
		adminPassword,
		"admin",
		"",
	); err != nil {
		return err
	}

	if err := createUserIfNotExists(
		viewerUsername,
		viewerPassword,
		"viewer",
		"demoA",
	); err != nil {
		return err
	}

	fmt.Println("Demo users ready")

	return nil
}

func createUserIfNotExists(
	username string,
	password string,
	role string,
	tenant string,
) error {

	var count int64

	if err := config.DB.
		Model(&entity.User{}).
		Where("username = ?", username).
		Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		return nil
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return err
	}

	user := entity.User{
		Username:     username,
		PasswordHash: string(hashedPassword),
		Role:         role,
		Tenant:       tenant,
	}

	return config.DB.Create(&user).Error
}
