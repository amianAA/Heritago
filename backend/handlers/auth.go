package handlers

import (
	"heritago/backend/orm"
	"net/http"

	//"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// User Auth is the endpoint where any user must be authenticated
func UserAuth(orm *orm.ORM) gin.HandlerFunc {

	return func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	}
}
