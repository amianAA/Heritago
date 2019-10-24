package handlers

import (
	"github.com/99designs/gqlgen/handler"
	"heritago/backend/gql"
	"heritago/backend/gql/resolvers"
	"heritago/backend/orm"
	//"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// User Auth is the endpoint where any user must be authenticated
func UserAuth(orm *orm.ORM) gin.HandlerFunc {
	c := gql.Config{
		Resolvers: &resolvers.Resolver{
			ORM: orm, // pass in the ORM instance in the resolvers to be used
		},
	}

	h := handler.GraphQL(gql.NewExecutableSchema(c))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
