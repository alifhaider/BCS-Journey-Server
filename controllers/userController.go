package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/alifhaider/BCS-Journey-Server/initializers"
	"github.com/alifhaider/BCS-Journey-Server/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	// Get the Email and Password from the request
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read the request body"})
		return
	}


	//Hash the password

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash the password"})
		return
	}


	// Create a new User object
	user := models.User{Email: body.Email, Password: string(hash)}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create the user"})
		return
	}

	//Respond with the user
	c.JSON(http.StatusOK, gin.H{})
}


func Login(c *gin.Context) {
	// Get the Email and Password from the request
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read the request body"})
		return
	}

	// Find the user with the email
	var user models.User
	result := initializers.DB.Where("email = ?", body.Email).First(&user)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	
	// Compare the password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Wrong password"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	"sub": user.ID,
	"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
})

// Sign and get the complete encoded token as a string using the secret
println(os.Getenv("TOKEN_SECRET"))
tokenSecret := os.Getenv("TOKEN_SECRET")
if tokenSecret == "" {
	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get token secret"})
	return
}

tokenString, err := token.SignedString([]byte(tokenSecret))
println(tokenString)

if err != nil {
	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
	return
}

	//Respond with the user
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}