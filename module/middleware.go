package SO_Module

import (
	SO_Class "SOApp_GO/class"
	"net/http"

	"github.com/gin-gonic/gin"
)

type middleware struct{}

// TokenAuthMiddleware Check token
func (m middleware) TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		// c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		// c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		// c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		// if c.Request.Method == "OPTIONS" {
		// 	c.AbortWithStatus(204)
		// 	return
		// }

		var hasil SO_Class.Hasil
		hasil.Sukses = false
		err := TokenValid(c.Request)
		if err != nil {
			hasil.Pesan = "unauthorized"
			hasil.Kode = "mdl001"
			c.JSON(http.StatusUnauthorized, hasil)
			c.Abort()
			return
		}

		// Begin 2025 11 03
		// Untuk handle koneksi database, baik yang sudah login atau belum login (Tetap harus bisa select)
		token, _ := verifyToken(c.Request)
		t, _ := extract(token)
		c.Set("globalDB", t.Database)
		// End 2025 11 03

		c.Next()
	}
}

func NoTokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}

var Middleware middleware
