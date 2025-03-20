package handlers

import (
	"go_todolist/db"
	"go_todolist/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

var Jwtkey = []byte("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJ0ZXN0X2tleSIsInNjb3BlIjplbmQwfQ.dhfurhfu4hfurhfu4hfurhf4uhfu4hf4uhfu4hf4uhf4uhfurhf4uhfurhf4uhfurhf4uhf")

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func Register(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if username is empty
	if user.Username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username cannot be empty"})
		return
	}

	// Check if password is empty
	if user.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "password cannot be empty"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user.Password = string(hashedPassword)

	// Execute the query
	var id int
	err = db.DB.QueryRow(`INSERT INTO "user" (username, password) VALUES ($1, $2) RETURNING id`,
		user.Username, user.Password).Scan(&id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User registered successfully",
		"id":      id,
	})
}
func Login(c *gin.Context) {
	var user models.User
	//handle invalid user input
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
	var storedUser models.User
	err := db.DB.QueryRow("select id , username, password from user where username=?", user.Username).Scan(&storedUser.Id, &storedUser.Username, &storedUser.Password)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found with given creds"})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid Credentialse"})
		return
	}
	//create expiration time as 24 hrs
	expirationTime := time.Now().Add(24 * time.Hour)
	//create claims with expiration time and username
	claims := &Claims{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	//create token with claims and sign it with jwtkey
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(Jwtkey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
