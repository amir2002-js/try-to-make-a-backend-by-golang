package auth

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"myProject/models"
	"os"
	"time"
)

func HashPassword(password string) (string, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPass), nil
}

func Role(username string, email string) string {
	adminU := os.Getenv("ADMIN_USERNAME")
	adminE := os.Getenv("ADMIN_EMAIL")

	if adminE == email && adminU == username {
		return "admin"
	}
	return "user"
}

func GenerateJWTTkn(user *models.User) (string, error) {
	jwtSecret := os.Getenv("JWT_SECRET_KEY")
	if jwtSecret == "" {
		return "", errors.New("JWT_SECRET_KEY environment variable not set  -- 33 auth.go")
	}
	claims := &jwt.MapClaims{
		"role":    user.Role,
		"user_id": user.UserID,
		"exp":     time.Now().Add(time.Hour * 24 * 90).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", errors.New("tkn string err  -- 44 auth.go")
	}
	return tokenString, nil
}
