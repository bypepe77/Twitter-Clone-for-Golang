package jwt

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type Manager interface {
	GenerateToken(userID uint, username string) (string, error)
	ValidateToken(encodedToken string) (*jwt.Token, error)
	FromMapClaims(userClaims jwt.MapClaims) (*Claims, error)
	GetClaims(c *gin.Context) (*Claims, error)
}

type Claims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

type tokenManager struct {
	secretKey string
	issure    string
}

// New creates a new token manager with the given secret key and issure.
func New(secretKey, issure string) Manager {
	return &tokenManager{
		secretKey: secretKey,
		issure:    issure,
	}
}

func (m *tokenManager) GenerateToken(userID uint, username string) (string, error) {
	claims := &Claims{
		userID,
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			Issuer:    m.issure,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(m.secretKey))
	if err != nil {
		return "", err
	}
	return t, nil
}

func (m *tokenManager) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, errors.New("invalid token")
		}
		return []byte(m.secretKey), nil
	})
}

func (m *tokenManager) FromMapClaims(userClaims jwt.MapClaims) (*Claims, error) {
	b, err := json.Marshal(userClaims)
	if err != nil {
		return nil, err
	}
	finalClaims := &Claims{}
	err = json.Unmarshal(b, &finalClaims)
	return finalClaims, err
}

func (m *tokenManager) GetClaims(c *gin.Context) (*Claims, error) {
	anyClaims, ok := c.Get("claims")
	if !ok {
		return nil, errors.New("error with claims in context")
	}
	claims, ok := anyClaims.(*Claims)
	if !ok || claims == nil {
		return nil, errors.New("error type of claims")
	}
	return claims, nil
}
