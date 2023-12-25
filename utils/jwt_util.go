package utils

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

type SignedDetails struct {
	Email      string
	First_name string
	Last_name  string
	Uid        string
	User_type  string
	jwt.StandardClaims
}

type Token struct {
	UserID       string `gorm:"primaryKey"`
	Token        string
	RefreshToken string
	UpdatedAt    time.Time
}

var SECRET_KEY string = os.Getenv("SECRET_KEY")

var db *gorm.DB

func GenerateAllTokens(email string, firstName string, lastName string, userType string, uid string) (signedToken string, signedRefreshToken string, err error) {
	claims := &SignedDetails{
		Email:      email,
		First_name: firstName,
		Last_name:  lastName,
		Uid:        uid,
		User_type:  userType,
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

func UpdateAllTokens(signedToken string, signedRefreshToken string, userId string) {
	var token Token

	// Check if the record already exists
	result := db.Where("user_id = ?", userId).First(&token)

	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		log.Panic(result.Error)
		return
	}

	// If the record exists, update it; otherwise, create a new record
	if result.RowsAffected > 0 {
		token.Token = signedToken
		token.RefreshToken = signedRefreshToken
		token.UpdatedAt = time.Now()
		db.Save(&token)
	} else {
		newToken := Token{
			UserID:       userId,
			Token:        signedToken,
			RefreshToken: signedRefreshToken,
			UpdatedAt:    time.Now(),
		}
		db.Create(&newToken)
	}

	fmt.Println("Tokens updated successfully")
}
