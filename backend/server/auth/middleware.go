package jwtauth

import (
   "context"
   "heritago/backend/orm/models"
   "net"
   "net/http"
   "strings"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
   name string
}

func JWTMiddleware() func(http.Handler) http.Handler {
   return func(next http.Handler) http.Handler {
      return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
         token := TokenFromHttpRequest(r)

         userId := UserIDFromToken(token)

         ip, _, _ := net.SplitHostPort(r.RemoteAddr)
         userAuth := models.UserAuth{
            UserID:    userId,
            IPAddress: ip,
            Token:     token,
         }

         // put it in context
         ctx := context.WithValue(r.Context(), userCtxKey, &userAuth)
         // and call the next with our new context
         r = r.WithContext(ctx)
         next.ServeHTTP(w, r)
      })
   }
}
func TokenFromHttpRequest(r *http.Request) string {
   reqToken := r.Header.Get("Authorization")
   var tokenString string
   splitToken := strings.Split(reqToken, "Bearer ")
   if len(splitToken) > 1 {
      tokenString = splitToken[1]
   }
   return tokenString
}
func UserIDFromToken(tokenString string) string {

   token, err := JwtDecode(tokenString)
   if err != nil {
      return ""
   }
   if claims, ok := token.Claims.(*models.UserClaims); ok && token.Valid {
      if claims == nil {
         return ""
      }
      return claims.UserId
   } else {
      return ""
   }
}
func ForContext(ctx context.Context) *models.UserAuth {
   raw := ctx.Value(userCtxKey)
   if raw == nil {
      return nil
   }
   return raw.(*models.UserAuth)
}
func GetAuthFromContext(ctx context.Context) *models.UserAuth {
   return ForContext(ctx)
}