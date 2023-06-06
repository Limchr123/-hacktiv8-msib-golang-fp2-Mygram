package auth

import (
	"errors"
	"mygram/campaign"
	"mygram/comment"
	"mygram/helper"
	"mygram/sosialMedia"
	"mygram/user"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Service interface {
	GenerateToken(userID int) (string, error)
	ValidasiToken(token string) (*jwt.Token, error)
	PhotoAuthorization() gin.HandlerFunc
	CommentAuthorization() gin.HandlerFunc
	SosmedAuthorization() gin.HandlerFunc
}

type jwtService struct {
	photoRepository   campaign.Repository
	commentRepository comment.Repository
	sosmedRepository  sosialMedia.Repository
}

var SECRET_KEY = []byte("mygram")

func NewService(photoRepository campaign.Repository, commentRepository comment.Repository, sosmedRepository sosialMedia.Repository) *jwtService {
	return &jwtService{photoRepository, commentRepository, sosmedRepository}
}

func (s *jwtService) PhotoAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var getParam campaign.GetPhotoDetailInput
		err := ctx.ShouldBindUri(&getParam)
		if err != nil {
			response := helper.APIresponse(http.StatusUnprocessableEntity, err)
			ctx.JSON(http.StatusUnprocessableEntity, response)
			return
		}

		currentUser := ctx.MustGet("currentUser").(user.User)
		// input.User.ID = currentUser.ID
		// user := ctx.MustGet("userData").(entity.User)

		// movieId, err := helper.GetParamId(getParam, "movieId")

		getOnePhoto, err := s.photoRepository.FindById(getParam.ID)

		if err != nil {
			ctx.AbortWithStatusJSON(404, err)
			return
		}

		if getOnePhoto.UserId != currentUser.ID {
			unauthorizedErr := errors.New("you are not authorized to modify the movie data")
			ctx.AbortWithStatusJSON(403, unauthorizedErr)
			return
		}

		ctx.Next()
	}
}

func (s *jwtService) CommentAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var getParam comment.GetCommentInput
		err := ctx.ShouldBindUri(&getParam)
		if err != nil {
			response := helper.APIresponse(http.StatusUnprocessableEntity, err)
			ctx.JSON(http.StatusUnprocessableEntity, response)
			return
		}

		currentUser := ctx.MustGet("currentUser").(user.User)
		// input.User.ID = currentUser.ID
		// user := ctx.MustGet("userData").(entity.User)

		// movieId, err := helper.GetParamId(getParam, "movieId")

		getOnePhoto, err := s.commentRepository.FindById(getParam.ID)

		if err != nil {
			ctx.AbortWithStatusJSON(404, err)
			return
		}

		if getOnePhoto.UserId != currentUser.ID {
			unauthorizedErr := errors.New("you are not authorized to modify the movie data")
			ctx.AbortWithStatusJSON(403, unauthorizedErr)
			return
		}

		ctx.Next()
	}
}

func (s *jwtService) SosmedAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var getParam sosialMedia.GetSosmedInput
		err := ctx.ShouldBindUri(&getParam)
		if err != nil {
			response := helper.APIresponse(http.StatusUnprocessableEntity, err)
			ctx.JSON(http.StatusUnprocessableEntity, response)
			return
		}

		currentUser := ctx.MustGet("currentUser").(user.User)
		// input.User.ID = currentUser.ID
		// user := ctx.MustGet("userData").(entity.User)

		// movieId, err := helper.GetParamId(getParam, "movieId")

		getOnePhoto, err := s.commentRepository.FindById(getParam.ID)

		if err != nil {
			ctx.AbortWithStatusJSON(404, err)
			return
		}

		if getOnePhoto.UserId != currentUser.ID {
			unauthorizedErr := errors.New("you are not authorized to modify the movie data")
			ctx.AbortWithStatusJSON(403, unauthorizedErr)
			return
		}

		ctx.Next()
	}
}

func (s *jwtService) GenerateToken(userID int) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signedToken, err
	}
	return signedToken, nil
}

func (s *jwtService) ValidasiToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Invalid Token")
		}
		return []byte(SECRET_KEY), nil
	})
	if err != nil {
		return token, err
	}

	return token, nil
}
