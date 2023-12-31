package utils

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/akers1023/models"
	jwt "github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

var SECRET_KEY string = os.Getenv("SECRET_KEY")

type SignedDetails struct {
	Email string
	Uid   string
	Role  string
	jwt.StandardClaims
}

func GenerateAllTokens(email string, uid string, role string) (signedToken string, signedRefreshToken string, err error) {
	claims := &SignedDetails{
		Email: email,
		Uid:   uid,
		Role:  role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	refreshClaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))

	if err != nil {
		log.Panic(err)
		return
	}

	return token, refreshToken, err
}

func ValidateToken(signedToken string) (claims *SignedDetails, msg string) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&SignedDetails{},

		func(token *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		},
	)

	if err != nil {
		msg = err.Error()
		return
	}

	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		msg = fmt.Sprintf("the token is invalid")
		msg = err.Error()
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		msg = fmt.Sprintf("token is expired")
		msg = err.Error()
		return
	}

	return claims, msg
}

// func UpdateAllTokens(signedToken string, signedRefreshToken string, userId string) {
// 	return
// }

func UpdateAllTokens(db *gorm.DB, userId string, signedToken, signedRefreshToken string) error {
	var user models.User

	// Find the user by ID
	if err := db.Where("id = ?", userId).First(&user).Error; err != nil {
		return err
	}

	// Update the user's token and refresh token
	user.Token = &signedToken
	user.RefreshToken = &signedRefreshToken
	user.UpdatedAt = time.Now()

	// Save the updated user back to the database
	if err := db.Save(&user).Error; err != nil {
		return err
	}

	return nil
}
