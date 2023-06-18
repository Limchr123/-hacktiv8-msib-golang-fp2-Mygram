package main

import (
	"mygram/auth"
	"mygram/campaign"
	"mygram/comment"
	"mygram/database"
	"mygram/handler"
	"mygram/helper"
	"mygram/sosialMedia"
	"mygram/user"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func main() {

	var port = os.Getenv("PORT")

	db := database.GetDataBaseInstance()

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	photoRepository := campaign.NewRepository(db)
	commentRepository := comment.NewRepository(db)
	sosialMediaRepository := sosialMedia.NewRepository(db)
	authService := auth.NewService(photoRepository, commentRepository, sosialMediaRepository)
	userHandler := handler.NewUserHandler(userService, authService)
	photoService := campaign.NewService(photoRepository)
	photoHandler := handler.NewPhotoHandler(photoService)
	commentService := comment.NewService(commentRepository)
	commentHandler := handler.NewCommentHandler(commentService)
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
	apiPhotos.DELETE("/:id", authMiddleware(authService, userService), authService.PhotoAuthorization(), photoHandler.DeletePhoto)
	// api.GET("/comments", commentHandler.GetComments)
	apiComments.GET("/", authMiddleware(authService, userService), commentHandler.GetComments)
	apiComments.POST("/", authMiddleware(authService, userService), commentHandler.CreateComment)
	apiComments.PUT("/:id", authMiddleware(authService, userService), commentHandler.UpdateComment)
	apiComments.DELETE("/:id", authMiddleware(authService, userService), authService.CommentAuthorization(), commentHandler.DeletedComment)
	apiSosmed.POST("/", authMiddleware(authService, userService), sosialMediaHandler.CreateSosmed)
	apiSosmed.GET("/", authMiddleware(authService, userService), sosialMediaHandler.GetSosmed)
	apiSosmed.PUT("/:id", authMiddleware(authService, userService), sosialMediaHandler.UpdateSosmed)
	apiSosmed.DELETE("/:id", authMiddleware(authService, userService), authService.SosmedAuthorization(), sosialMediaHandler.DeletedSosmed)

	router.Run(":" + port)
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
