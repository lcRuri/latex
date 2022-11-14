package define

import "github.com/dgrijalva/jwt-go"

type UserClaim struct {
	Id       int
	Identity string
	Name     string
	jwt.StandardClaims
}

var JwtSecret = "latex"

var TokenExpireTime = 3600 * 7

const BaseUploadPath = "./file"
