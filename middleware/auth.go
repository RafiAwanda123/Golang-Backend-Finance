package middleware

import (
	"net/http"

	"github.com/RafiAwanda123/Finance-UMKM/utils"

	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			utils.APIError(c, http.StatusUnauthorized, "Token tidak tersedia")
			return
		}

		claims, err := utils.ValidateJWT(tokenString)
		if err != nil {
			utils.APIError(c, http.StatusUnauthorized, "Token tidak valid")
			return
		}

		// Simpan user_id di context
		c.Set("user_id", claims.UserID)
		c.Next()
	}
}
