package define

import "github.com/golang-jwt/jwt/v4"

type UserClaim struct {
	Id       int
	Identity string
	Name     string 
	jwt.StandardClaims
}

var Jwtkey = "cloud-disk-key"
