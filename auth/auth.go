package auth

import (
	"fmt"
	"net/http"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

var (
	jwtSigningMethod = jwt.SigningMethodHS256
	JwtSecret        = "JWT_SECRET"
)

const (
	JWT_KEY = "JWT_SECRET"
)

func init() {
	if s := os.Getenv(JWT_KEY); s != "" {
		JwtSecret = s
	}
}
func NewToken(m map[string]interface{}) (string, error) {
	claims := jwt.MapClaims{
		"iss": "colleague",
		"aud": "colleague",
		"nbf": time.Now().Add(-time.Minute * 5).Unix(),
		"exp": time.Now().Add(time.Minute * 10).Unix(),
	}
	for k, v := range m {
		claims[k] = v
	}
	return jwt.NewWithClaims(jwtSigningMethod, claims).SignedString([]byte(JwtSecret))
}

//生成token的时间比DMZ服务器时间靠前，token才能验证通过
func NewTokenForDMZ(m map[string]interface{}) (string, error) {
	claims := jwt.MapClaims{
		"iss": "colleague",
		"aud": "colleague",
		"nbf": time.Now().Add(-time.Second * 30).Unix(),
		"exp": time.Now().Add(time.Hour * 15).Unix(),
	}
	for k, v := range m {
		claims[k] = v
	}
	return jwt.NewWithClaims(jwtSigningMethod, claims).SignedString([]byte(JwtSecret))
}

func Renew(token string) (string, error) {
	claim, err := Extract(token)
	if err != nil {
		return "", err
	}
	claim["nbf"] = time.Now().Unix()
	claim["exp"] = time.Now().Add(time.Hour * 15).Unix()
	return jwt.NewWithClaims(jwtSigningMethod, claim).SignedString([]byte(JwtSecret))
}

func EditPayload(token string, m map[string]interface{}) (string, error) {
	claimInfo, err := Extract(token)
	if err != nil {
		return "", err
	}

	for k, v := range m {
		claimInfo[k] = v
	}

	return jwt.NewWithClaims(jwtSigningMethod, claimInfo).SignedString([]byte(JwtSecret))
}

func Extract(token string) (jwt.MapClaims, error) {
	return ExtractWithSecret(token, JwtSecret)
}
func ExtractWithSecret(token, jwtSecret string) (jwt.MapClaims, error) {
	if token == "" {
		return nil, &echo.HTTPError{Code: http.StatusUnauthorized, Message: "Required authorization token not found."}
	}

	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) { return []byte(JwtSecret), nil })
	if err != nil {
		return nil, &echo.HTTPError{Code: http.StatusUnauthorized, Message: fmt.Sprintf("Error parsing token: %v", err)}
	}

	if jwtSigningMethod != nil && jwtSigningMethod.Alg() != parsedToken.Header["alg"] {
		return nil, &echo.HTTPError{Code: http.StatusUnauthorized, Message: fmt.Sprintf("Expected %s signing method but token specified %s",
			jwtSigningMethod.Alg(),
			parsedToken.Header["alg"])}
	}

	if !parsedToken.Valid {
		return nil, &echo.HTTPError{Code: http.StatusUnauthorized, Message: "Token is invalid"}
	}

	claimInfo, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, &echo.HTTPError{Code: http.StatusUnauthorized, Message: "invalid token"}
	}
	return claimInfo, nil
}
