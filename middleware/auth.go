package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"

)

// Auth middleware
func Auth() gin.HandlerFunc  {
	return func (c *gin.Context)  {
		var token string
		if tokenString, _ := c.Request.Header["Token"]; len(tokenString)>0{
			token = tokenString[0]
			t,_ := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
				// Don't forget to validate the alg is what you expect:
				if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
					c.Abort()
				}
			
				// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
				return []byte("luowanshun"), nil
			})
			
			if _, ok := t.Claims.(jwt.MapClaims); ok && t.Valid {
				c.Next()
			} else {
				c.Abort()
			}
		}else {

			c.JSON(200, gin.H{
				"token": tokenString[0],
			})
			c.Abort()
		}

		
	}
}