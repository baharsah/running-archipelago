package jawatoken

import "github.com/golang-jwt/jwt/v4"

var KodeNuklir = "KodeNuklirBerbahaya"

func GenerateToken(claims *jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	webtoken, err := token.SignedString([]byte(KodeNuklir))
	if err != nil {
		return "", err
	}
	return webtoken, nil
}

// func DecodeToken(tokenString string) (*jwt.Token, error) {

// }
