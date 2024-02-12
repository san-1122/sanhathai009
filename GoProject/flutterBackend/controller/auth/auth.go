package auth

import (
	// "crypto/hmac"
	"flutterBackend/orm"
	"net/http"
	"time"

	// "time"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var  hmacSampleSecret []byte
// Binding from JSON
type RegisterBody struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Fullname string `json:"fullname" binding:"required"`
	Avatar   string `json:"avatar" binding:"required"`
}

func Register(c *gin.Context){
	var json RegisterBody
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//
	var userExists orm.Tbl_User
	orm.Db.Where("username = ?", json.Username).First(&userExists)
	if(userExists.ID > 0){
		c.JSON(http.StatusOK, gin.H{"Status": "error","message": "User Exists",})  
		return
	}
	//
	encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte(json.Password), 10)
	user := orm.Tbl_User{Username: json.Username, Password:string(encryptedPassword), Fullname:json.Fullname, Avatar:json.Avatar}
	orm.Db.Create(&user)
	if(user.ID > 0){
		c.JSON(http.StatusOK, gin.H{"Status": "Ok","message": "User Create Success","User_ID": user.ID,})  
	}else{
		c.JSON(http.StatusOK, gin.H{"Status": "error","message": "User Create Failed",})  
	}
}

// Binding from JSON
type LoginBody struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context){
	var json LoginBody
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userExists orm.Tbl_User
	orm.Db.Where("username = ?", json.Username).First(&userExists)
	if(userExists.ID == 0){
		c.JSON(http.StatusOK, gin.H{"Status": "error","message": "User Does Not Exists",})  
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(userExists.Password),[]byte(json.Password))

	if err == nil {
		hmacSampleSecret = []byte(os.Getenv("JWT_SECRET_KET"))
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"userId": userExists.ID,
			"exp": time.Now().Add(time.Minute * 1).Unix(),
		})
		// Sign and get the complete encoded token as a string using the secret
		tokenString, err := token.SignedString(hmacSampleSecret)
		fmt.Println(tokenString, err)

		c.JSON(http.StatusOK, gin.H{"Status": "Success","message": "Login Success", "token": tokenString})
	}else{
		c.JSON(http.StatusOK, gin.H{"Status": "error","message": "Login failed",})
	}
}