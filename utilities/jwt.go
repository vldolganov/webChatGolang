package utilities

import (
	"time"

	"github.com/pascaldekloe/jwt"

	"encoding/json"
)

func CreateToken(userId uint, lifeTime time.Duration, secret string) string {

	var claims jwt.Claims
	claims.Issued = jwt.NewNumericTime(time.Now().Round(time.Second))
	claims.Set = map[string]interface{}{"id": userId}
	claims.Expires = jwt.NewNumericTime(time.Now().Add(lifeTime))
	var extraString = ExtraString{
		"HS256",
		"JWT",
	}

	jsonExtra, _ := json.Marshal(extraString)
	token, _ := claims.HMACSign(jwt.HS256, []byte(secret), jsonExtra)
	return string(token)
}

func CheckToken(token []byte, secret []byte) interface{} {
	verify, _ := jwt.HMACCheck(token, secret)
	verifiedId := verify.Set["id"]
	return verifiedId
}
