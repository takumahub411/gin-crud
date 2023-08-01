package controllers

import (
	"gin-curd/initializers"
	"gin-curd/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"time"
)

func Signup(c *gin.Context) {
	var SignupInfo models.UserInfo
	if c.ShouldBindJSON(&SignupInfo) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "読み込み失敗",
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(SignupInfo.Password), 12)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	user := models.User{Email: SignupInfo.Email, Password: string(hash)}

	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "登録失敗",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "登録完了",
	})
}

func Login(c *gin.Context) {
	var loginUser models.UserInfo
	var user models.User
	if c.ShouldBindJSON(&loginUser) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "読み込み失敗",
		})
		return
	}

	initializers.DB.First(&user, "email = ?", loginUser.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "メールアドレスもしくはパスワードが間違っています。",
		})
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginUser.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "メールアドレスもしくはパスワードが間違っています。",
		})
		return
	}

	// create token

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = user.ID
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	//token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	//	"sub": user.ID,
	//	"exp": time.Now().Add(time.Hour * 72).Unix(),
	//})

	t, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "トークンの作成に失敗しました",
		})
		return
	}

	// set cookie
	//3600 1h
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Auth", t, 3600*24, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"token": t,
	})
}

func Validate(c *gin.Context) {

	user, _ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{
		"massage": user,
	})
}
