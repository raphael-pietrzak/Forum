package forum

import (
	"encoding/json"
	"errors"
	_ "fmt"
	"io/ioutil"
	"net/http"
	_ "time"
)

type GoogleClaims struct {
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	FirstName     string `json:"given_name"`
	LastName      string `json:"family_name"`
	// jwt.StandardClaims
}

func getGooglePublicKey(keyID string) (string, error) {
	resp, err := http.Get("https://www.googleapis.com/oauth2/v1/certs")
	if err != nil {
		return "", err
	}
	dat, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	myResp := map[string]string{}
	err = json.Unmarshal(dat, &myResp)
	if err != nil {
		return "", err
	}
	key, ok := myResp[keyID]
	if !ok {
		return "", errors.New("key not found")
	}
	return key, nil
}

// func ValidateGoogleJWT(tokenString string) (GoogleClaims, error) {
// 	claimsStruct := GoogleClaims{}

// 	token, err := jwt.ParseWithClaims(
// 		tokenString,
// 		&claimsStruct,
// 		func(token *jwt.Token) (interface{}, error) {
// 			pem, err := getGooglePublicKey(fmt.Sprintf("%s", token.Header["kid"]))
// 			if err != nil {
// 				return nil, err
// 			}
// 			key, err := jwt.ParseRSAPublicKeyFromPEM([]byte(pem))
// 			if err != nil {
// 				return nil, err
// 			}
// 			return key, nil
// 		},
// 	)
// 	if err != nil {
// 		return GoogleClaims{}, err
// 	}

// 	claims, ok := token.Claims.(*GoogleClaims)
// 	if !ok {
// 		return GoogleClaims{}, errors.New("Invalid Google JWT")
// 	}

// 	if claims.Issuer != "accounts.google.com" && claims.Issuer != "https://accounts.google.com" {
// 		return GoogleClaims{}, errors.New("iss is invalid")
// 	}

// 	if claims.Audience != "822476215105-a1qeg4jvqdnjlut874jucd3lh0srpfr8.apps.googleusercontent.com" {
// 		return GoogleClaims{}, errors.New("aud is invalid")
// 	}

// 	if claims.ExpiresAt < time.Now().UTC().Unix() {
// 		return GoogleClaims{}, errors.New("JWT is expired")
// 	}

// 	return *claims, nil
// }

// func (cfg config) loginHandler(w http.ResponseWriter, r *http.Request) {
// 	defer r.Body.Close()

// 	// parse the GoogleJWT that was POSTed from the front-end
// 	type parameters struct {
// 		GoogleJWT *string
// 	}
// 	decoder := json.NewDecoder(r.Body)
// 	params := parameters{}
// 	err := decoder.Decode(&params)
// 	if err != nil {
// 		respondWithError(w, 500, "Couldn't decode parameters")
// 		return
// 	}

// 	// Validate the JWT is valid
// 	claims, err := auth.ValidateGoogleJWT(*params.GoogleJWT) 
// 	if err != nil {
// 		respondWithError(w, 403, "Invalid google auth")
// 		return
// 	}
// 	if claims.Email != user.Email {
// 		respondWithError(w, 403, "Emails don't match")
// 		return
// 	}

// 	// create a JWT for OUR app and give it back to the client for future requests
// 	tokenString, err := auth.MakeJWT(claims.Email, cfg.JWTSecret)
// 	if err != nil {
// 		respondWithError(w, 500, "Couldn't make authentication token")
// 		return
// 	}

// 	respondWithJSON(w, 200, struct {
// 		Token string `json:"token"`
// 	}{
// 		Token: tokenString,
// 	})
// }

// func respondWithError(w http.ResponseWriter, i int, s string) {
// 	panic("unimplemented")
// }

// // func respondWithJSON(w http.ResponseWriter, i int, _ struct{Token string "json:\"token struct{Token string "json:\"token\""}) {
// // 	panic("unimplemented")
// // }
