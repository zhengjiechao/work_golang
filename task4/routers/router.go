package routers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"
	"work_golang/task4/controllers"
	"work_golang/task4/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func UserRouterInit(r *gin.Engine) {
	userGroup := r.Group("/user")
	{
		userGroup.POST("/register", controllers.UserController{}.Register)
		userGroup.POST("/login", controllers.UserController{}.Login)
	}
}

func PostRouterInit(r *gin.Engine) {
	postGroup := r.Group("/post", authMiddleWare)
	{
		postGroup.POST("/add", controllers.PostController{}.Add)
		postGroup.PUT("/edit", controllers.PostController{}.Edit)
		postGroup.DELETE("/del", controllers.PostController{}.Del)
		postGroup.GET("/findOne", controllers.PostController{}.FindOne)
		postGroup.GET("/findAll", controllers.PostController{}.FindAll)
	}
}

func CommentRouterInit(r *gin.Engine) {
	commentGroup := r.Group("/comment")
	commentGroup.Use(authMiddleWare)
	{
		commentGroup.POST("/add", controllers.CommentController{}.Add)
		commentGroup.GET("/findByPostId/:postId", controllers.CommentController{}.FindByPostId)
	}
}

// 中间件-验证
func authMiddleWare(c *gin.Context) {
	secretKey := []byte("your_secret_key")
	tokenString := c.GetHeader("Authorization")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token")
		}
		return secretKey, nil
	})

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Printf("token的值 %v\n", claims)
		exp := claims["exp"]
		v, ok := exp.(string)
		if ok {
			nowUnix := time.Now().Unix()
			expUnix, _ := strconv.ParseInt(v, 10, 64)
			// token有效期小于当前
			if expUnix < nowUnix {
				c.JSON(http.StatusUnauthorized, gin.H{"error1": "unauthorized"})
				c.Abort()
				return
			}
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error2": "unauthorized"})
			c.Abort()
			return
		}

		key := claims["id"]
		id, ok := key.(float64)
		if ok {
			user := models.User{}
			user.ID = uint(id)
			models.DB.Find(&user)
			if user.Username != "" {
				c.Set("userId", user.ID)
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"error3": "unauthorized"})
				c.Abort()
				return
			}
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error4": "unauthorized"})
			c.Abort()
			return
		}
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error5": "unauthorized"})
		c.Abort()
		return
	}
}
