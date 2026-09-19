package sevices

import (
	"errors"
	"os"
	"time"

	"log-management/backend/config"
	"log-management/backend/entity"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Login(
	username string,
	password string,
) (string, entity.User, error) {

	var user entity.User

	if err := config.DB.
		Where("username = ?", username).
		First(&user).Error; err != nil {

		return "", user, errors.New(
			"invalid username or password",
		)
	}

	err := bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(password),
	)

	if err != nil {
		return "", user, errors.New(
			"invalid username or password",
		)
	}

	claims := jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"role":     user.Role,
		"tenant":   user.Tenant,
		"exp":      time.Now().Add(8 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	tokenString, err := token.SignedString(
		[]byte(os.Getenv("JWT_SECRET")),
	)

	if err != nil {
		return "", user, err
	}

	return tokenString, user, nil
}
