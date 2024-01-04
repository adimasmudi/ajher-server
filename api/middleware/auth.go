package middleware

import (
	"ajher-server/internal/user"
	"ajher-server/utils"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type authMiddleware struct {
	userService user.Service
}

func NewAuthMiddleware(userService user.Service) *authMiddleware {
	return &authMiddleware{userService}
}

func Auth(ctx *gin.Context) (jwt.MapClaims, error) {
	authHeader := ctx.GetHeader("Authorization")

	if !strings.Contains(authHeader, "Bearer") {
		return nil, errors.New("unauthorized")
	}

	tokenString := ""
	arrayToken := strings.Split(authHeader, " ")
	if len(arrayToken) == 2 {
		tokenString = arrayToken[1]
	}

	token, err := utils.ValidateToken(tokenString)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		return nil, err
	}

	return claims, err
}

func (a *authMiddleware) AuthMiddleware(ctx *gin.Context) {

	claims, err := Auth(ctx)

	if err != nil {
		response := utils.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
	}

	userId := int(claims["user_id"].(float64))

	user, err := a.userService.GetUserById(userId)

	if err != nil {
		response := utils.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	ctx.Set("currentUser", user)
}

func (a *authMiddleware) RefreshTokenMiddleware(ctx *gin.Context) {
	claims, err := Auth(ctx)

	if err != nil {
		response := utils.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
	}

	expiration := claims["exp"].(float64)

	if !(int64(expiration) > time.Now().Unix()) {
		response := utils.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
	}
}
