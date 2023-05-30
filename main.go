package main

import (
	"log"
	"mygram/auth"
	"mygram/campaign"
	"mygram/comment"
	"mygram/handler"
	"mygram/helper"
	"mygram/sosialMedia"
	"mygram/user"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := "root:@tcp(127.0.0.1:3306)/tugas2?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Db Connestion Error")
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	authService := auth.NewService()
	userHandler := handler.NewUserHandler(userService, authService)
	photoRepository := campaign.NewRepository(db)
	photoService := campaign.NewService(photoRepository)
	photoHandler := handler.NewPhotoHandler(photoService)
	commentRepository := comment.NewRepository(db)
	commentService := comment.NewService(commentRepository)
	commentHandler := handler.NewCommentHandler(commentService)
	sosialMediaRepository := sosialMedia.NewRepository(db)
	sosialMediaService := sosialMedia.NewService(sosialMediaRepository)
	sosialMediaHandler := handler.NewSosmedHandler(sosialMediaService)

	db.AutoMigrate(&user.User{})

	router := gin.Default()
	api := router.Group("/users")
	apiPhotos := router.Group("/photos")
	apiComments := router.Group("/comments")
	apiSosmed := router.Group("/socialmedias")
	api.POST("/register", userHandler.RegisterUser)
	api.POST("/login", userHandler.Login)
	api.PUT("/:id", authMiddleware(authService, userService), userHandler.UpdatedUser)
	api.DELETE("/", authMiddleware(authService, userService), userHandler.DeletedUser)
	apiPhotos.GET("/", authMiddleware(authService, userService), photoHandler.GetCampaigns)
	apiPhotos.POST("/", authMiddleware(authService, userService), photoHandler.CreateImage)
	apiPhotos.PUT("/:id", authMiddleware(authService, userService), photoHandler.UpdatedCampaign)
	apiPhotos.DELETE("/:id", authMiddleware(authService, userService), photoHandler.DeletePhoto)
	// api.GET("/comments", commentHandler.GetComments)
	apiComments.GET("/", authMiddleware(authService, userService), commentHandler.GetComments)
	apiComments.POST("/", authMiddleware(authService, userService), commentHandler.CreateComment)
	apiComments.PUT("/:id", authMiddleware(authService, userService), commentHandler.UpdateComment)
	apiComments.DELETE("/:id", authMiddleware(authService, userService), commentHandler.DeletedComment)
	apiSosmed.POST("/", authMiddleware(authService, userService), sosialMediaHandler.CreateSosmed)
	apiSosmed.GET("/", authMiddleware(authService, userService), sosialMediaHandler.GetSosmed)
	apiSosmed.PUT("/:id", authMiddleware(authService, userService), sosialMediaHandler.UpdateSosmed)
	apiSosmed.DELETE("/:id", authMiddleware(authService, userService), sosialMediaHandler.DeletedSosmed)

	router.Run(":8080")
}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		// fmt.Println(authHeader)
		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIresponse(http.StatusUnauthorized, nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		tokenString := ""
		arrToken := strings.Split(authHeader, " ")
		if len(arrToken) == 2 {
			//nah ini kalau emang ada dua key nya dan sesuai, maka tokenString tadi masuk ke arrtoken index ke1
			tokenString = arrToken[1]
		}
		token, err := authService.ValidasiToken(tokenString)
		// fmt.Println(token, err)
		if err != nil {
			response := helper.APIresponse(http.StatusUnauthorized, nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		claim, ok := token.Claims.(jwt.MapClaims)
		// fmt.Println(claim, ok)
		if !ok || !token.Valid {
			response := helper.APIresponse(http.StatusUnauthorized, nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		userID := int(claim["user_id"].(float64))

		user, err := userService.GetUserByid(userID)
		// fmt.Println(user, err)
		if err != nil {
			response := helper.APIresponse(http.StatusUnauthorized, nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		c.Set("currentUser", user)
	}
}
