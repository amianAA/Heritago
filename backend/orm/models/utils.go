package models

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

func HashAndSaltPwd(pwd string) string {

    hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
    if err != nil {
        log.Println(err)
    }    // GenerateFromPassword returns a byte slice so we need to
    // convert the bytes to a string and return it
    return string(hash)
}

func ComparePasswords(hashedPwd string, plainPwd []byte) bool {
    byteHash := []byte(hashedPwd)
    err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
    if err != nil {
        log.Println(err)
        return false
    }

    return true
}