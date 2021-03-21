package auth

import (
	"Golang/config"
	"Golang/models"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	KeyUserID       = "id"
	KeyUserRole     = "role"
	KeyTokenExpired = "exp"
)

type IAuth interface {
	Authentication(handler http.Handler) http.Handler
	GenerateToken(user *models.User) (string, error)
}

type Auth struct {
	Secret string
}

func NewAuth(appConfig *config.DBConfig) IAuth {
	return &Auth{
		Secret: appConfig.Secret,
	}
}

func (auth *Auth) GenerateToken(user *models.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims[KeyUserID] = user.ID
	claims[KeyUserRole] = 1
	claims[KeyTokenExpired] = time.Now().Add(time.Minute * 10).Unix()
	signedToken, err := token.SignedString([]byte(auth.Secret))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

//Authentication
func (auth *Auth) Authentication(handler http.Handler) http.Handler {
	//fmt.Println("Entered")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		extractToken := func() string {
			bearToken := r.Header.Get("Authorization")
			fmt.Println(bearToken)
			strArr := strings.Split(bearToken, " ")
			if len(strArr) == 4 {
				fmt.Println("passed")
				//fmt.Println(strArr[3])
				return strArr[3]
			}
			//fmt.Println(len(strArr[1]))
			return ""
		}
		//fmt.Println(r.Header)
		if r.Header["Authorization"] != nil {
			tokenString := extractToken()
			/*
				passing nil toeknstring
				Getnerating error
			*/
			//fmt.Println(tokenString)
			token, err := jwt.Parse(tokenString, func(token2 *jwt.Token) (interface{}, error) {
				fmt.Println("hi")
				if _, ok := token2.Method.(*jwt.SigningMethodHMAC); !ok {

					fmt.Println("error here")
					return nil, errors.New("Invalid authoraization token")
				}
				return []byte(auth.Secret), nil
			})
			if err != nil {
				fmt.Println("errefsdfafsdfadf")
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(err)
				return
			}
			if token.Valid {
				claims := token.Claims.(jwt.MapClaims)
				r.Header.Set(KeyUserID, strconv.FormatFloat((claims[KeyUserID]).(float64), 'f', 0, 64))
				handler.ServeHTTP(w, r)
			} else {
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode("An authorized header is required")
			}
		}
	})
}
