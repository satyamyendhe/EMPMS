package utils

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	SECRET_KEY = "vsys_jwt_token"
)

func Nop(strs ...string) {
}

func JoinStr(elements ...string) string {
	return strings.Join(elements, "")
}

// Create Jwt token using mobile number and hard coded secret key
// Token expiry timing is kept as 1 hr
func CreateJwtToken(email string) (string, time.Time, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = email
	expire := time.Now().UTC().Add(time.Hour * 1)
	claims["exp"] = expire.Unix()
	claims["orig_iat"] = time.Now().UTC().Unix()

	tokenString, err := token.SignedString([]byte(SECRET_KEY))
	// log.Printf("\n\nToken string is  : %s", tokenString)
	if err != nil {
		log.Printf("Error signing token: %v", err)
		return "", time.Now(), err
	}

	// log.Printf("Created token: %s", tokenString)
	return tokenString, expire, nil
}

// ValidateJwtToken checks the validity of the JWT token
func ValidateJwtToken(tokenString string) bool {
	secretKey := []byte(SECRET_KEY)
	token, err := jwt.Parse(tokenString, func(token1 *jwt.Token) (interface{}, error) {
		if _, ok := token1.Method.(*jwt.SigningMethodHMAC); !ok {
			log.Printf("unexpected signing method: %v", token1.Header["alg"])
			return nil, nil
		}
		return secretKey, nil
	})
	// log.Printf("%v", token) --> token main object hai
	if err != nil {
		// log.Printf("Token parsing error: %v", err)
		return false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if exp, ok := claims["exp"].(float64); ok {
			if int64(exp) > time.Now().UTC().Unix() {
				//todo validate the email here ->  it should be same in login and in cookies
				return true
			}
		}
	}
	log.Printf("Token is invalid or expired")
	return false
}

func GetCookieValueEmail(r *http.Request, cookieName string) (string, error) {
	cookie, err := r.Cookie(cookieName)
	if err != nil {
		// Cookie not found
		return "", fmt.Errorf("cookie not found: %w", err)
	}

	// Decode the JWT token
	token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRET_KEY), nil
	})
	if err != nil {
		return "", fmt.Errorf("error parsing token: %w", err)
	}

	// Check if the token is valid
	if !token.Valid {
		return "", errors.New("invalid token")
	}

	// Extract claims from the token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("invalid claims")
	}

	// Get the email from the claims
	email, ok := claims["email"].(string)
	if !ok {
		return "", errors.New("email claim not found or invalid")
	}

	return email, nil
}
