package user

import (
	// "crypto/hmac"
	"flutterBackend/orm"
	// "go/token"
	"net/http"
	// "strings"

	// "time"

	// "fmt"
	// "os"

	"github.com/gin-gonic/gin"
	// "github.com/golang-jwt/jwt/v5"
	// "golang.org/x/crypto/bcrypt"
)

func ReadAll(c *gin.Context) {
	var users []orm.Tbl_User
	orm.Db.Find(&users)
	c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "User Read Success", "users": users})
}

func Profile(c *gin.Context) {
	userId := c.MustGet("userId").(float64)
	var user orm.Tbl_User
	orm.Db.First(&user, userId)
	c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "User Read Success", "user": user})
}