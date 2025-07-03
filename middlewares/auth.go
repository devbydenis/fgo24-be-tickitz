package middlewares

import (
	"backend-cinemax/utils"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func VerifyToken(token string) (*jwt.Token, error) {
	godotenv.Load()

	return jwt.Parse(token, func(t *jwt.Token) (any, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC);
		if !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

}
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		headerAuth := ctx.GetHeader("Authorization")	// ambil header Authorization buat dicek tokennya
		fmt.Println("headerAuth:", headerAuth)
		if headerAuth == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.Response{	// abort buat menghentikan eksekusi middleware selanjutnya
				Success: false,
				Message: "Unauthorized - No Authorization header provided",
			})
			return
		}
		
		splitToken := strings.Split(headerAuth, " ")
		if len(splitToken) != 2 || splitToken[0] != "Bearer" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.Response{
				Success: false,
				Message: "Unauthorized - Invalid Authorization header format",
			})
			return
		}
		
		verifiedToken, err := VerifyToken(splitToken[1])	// verifikasi token yang didapat dari header Authorization
		fmt.Println("verifiedToken:", verifiedToken)
		if err != nil || !verifiedToken.Valid {  
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.Response{
				Success: false,
				Message: "Unauthorized - Invalid or expired token",
				Errors:  err.Error(),
			})
			return
		}

		if claims, ok := verifiedToken.Claims.(jwt.MapClaims); ok {
			if idFloat, ok := claims["id"].(float64); ok {
				userId := int(idFloat)
				ctx.Set("userId", userId)
			}
		}
		ctx.Next()
	}
}