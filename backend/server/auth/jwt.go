package jwtauth

import (
   "github.com/dgrijalva/jwt-go"
   "heritago/backend/orm/models"
)
var mySigningKey = []byte("dexp.io")

func JwtDecode(token string) (*jwt.Token, error) {
   return jwt.ParseWithClaims(token, &models.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
      return mySigningKey, nil
   })
}
func JwtCreate(userID string, expiredAt int64) string {
   claims := models.UserClaims{
      StandardClaims: jwt.StandardClaims{
         ExpiresAt: expiredAt,
         Issuer:    "jwtauth",
      },
      UserId: userID,
   }
   token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
   ss, _ := token.SignedString(mySigningKey)
   return ss
}