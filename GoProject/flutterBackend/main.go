package main

import (
	// "encoding/json"
	AuthController "flutterBackend/controller/auth"
	UserController "flutterBackend/controller/user"
	"flutterBackend/middleware"
	"flutterBackend/orm"
	"fmt"

	// "net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	// "golang.org/x/crypto/bcrypt"
	// "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Binding from JSON
type Register struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Fullname string `json:"fullname" binding:"required"`
	Avatar   string `json:"avatar" binding:"required"`
}

// Model Tbl_User
type Tbl_User struct {
	gorm.Model
	Username string
	Password string
	Fullname string
	Avatar   string
}

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("Error loading .env file")
	}

	orm.InitDB()
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AddAllowHeaders("authorization")

	r.Use(cors.New(config))
	r.Use(cors.Default())
	//GET
	// r.GET("/register", func(c *gin.Context) {c.JSON(http.StatusOK, gin.H{ "message": "register",})})
	authorized := r.Group("/users", middleware.JWTAuthen())
	authorized.GET("/readall", UserController.ReadAll)
	authorized.GET("/profile", UserController.Profile)
	//POST
	r.POST("/register", AuthController.Register)
	r.POST("/login", AuthController.Login)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
