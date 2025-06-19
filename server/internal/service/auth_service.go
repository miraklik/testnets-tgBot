package service

import (
	"context"
	"errors"
	"log"
	"os"
	"tg-bot-server/internal/models"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	db *mongo.Database
}

func NewAuthService(db *mongo.Database) *AuthService {
	return &AuthService{db: db}
}

func (s *AuthService) AuthenticateUser(ctx context.Context, name, password string) (string, error) {
	user := &models.User{}
	if err := s.db.Collection("admin").FindOne(ctx, bson.M{"name": name}).Decode(user); err != nil {
		if err == mongo.ErrNoDocuments {
			return "", errors.New("user not found")
		}
		return "", err
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid password")
	}

	token, err := generateJWT(name)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *AuthService) ChangePassword(ctx context.Context, name, password, newPassword string) (string, error) {
	user := &models.User{}
	err := s.db.Collection("admin").FindOne(ctx, bson.M{"name": name}).Decode(user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "", errors.New("invalid username")
		}
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid current password")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	_, err = s.db.Collection("admin").UpdateOne(ctx, bson.M{"name": name}, bson.M{"$set": bson.M{"password": string(hashedPassword)}})
	if err != nil {
		return "", err
	}

	return "password updated successfully", nil
}

func generateJWT(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	secretKey := os.Getenv("SECRET_KEY")
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
