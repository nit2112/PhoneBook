package common

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
	jwt "github.com/dgrijalva/jwt-go"
)

const(
		privKeyPath = "keys/app.rsa"
		pubKeyPath = "keys/app.rsa.pub"
)

var(
	verifyKey ,signKey []byte
)

func initKeys (){
	var err error

	signKey , err = ioutil.ReadFile(privKeyPath)
	if err != nil {
			log.Fatalf("[initKeys]:%s\n",err)
	}

	verifyKey, err = ioutil.ReadFile(pubKeyPath)
	if err != nil {
		  log.Fatalf("[initKeys]:%s\n",err)
		  panic(err)
	}
}

func GenerateJWT(name , role string) (string, error) {
	t := jwt.New(jwt.GetSigningMethod("RS256"))
	t.Claims["iss"] = "admin"
	t.Claims["UserInfo"] = struct {
		Name string
		role string
	}{name, role}

t.Claims["exp"] = time.Now().Add(time.Minute * 20).Unix()
tokenString, err := t.SignedString(signKey)
if err != nil {
	return "", err
}
return tokenString, nil
}

func Authorize(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	token , err := jwt.ParseFromRequest(r, func(token *jwt.Token)(interface{},error){
		return verifyKey, nil
		})
	if err != nil {
		switch err.(type){
		case *jwt.ValidationError:
			vErr := err.(*jwt.ValidationError)
			switch vErr.Errors {
			case jwt.ValidationErrorExpired:
				DisplayAppError(
					w,
					err,
					"Access Token is expired, get a new Toekn",
					401,
				)
				return
			default:
				DisplayAppError(w,
					err,
					"Error while parsing Access Token!",
					500,
					)
				return
			}
		default:
			DisplayAppError(w,
				err,
				"Error while parsing Access Token!",
				500,
				)
			return
		}
	}
	if token. Valid {
		next(w,r)
	} else {
		DisplayAppError(w,
			err,
			"Invalid Access Token",
			401,
			)
	}
}