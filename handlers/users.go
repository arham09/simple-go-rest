package handlers

import (
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/arham09/book-rest/middlewares"
	"github.com/arham09/book-rest/models"

	jwt "github.com/dgrijalva/jwt-go"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var response models.Response

	err := r.ParseForm()
	if err != nil {
		log.Print(err)
	}

	name := r.Form.Get("name")
	email := r.Form.Get("email")
	password := middlewares.EncryptPassword(r.Form.Get("password"))

	if middlewares.VerifyEmail(email) == false {
		middlewares.Error(w, 400, "Email Invalid")
		return
	}

	status, err := models.CheckUser(&email)
	if err != nil {
		log.Print(err)
		return
	}

	if *status == 1 {
		middlewares.Error(w, 400, "Email exists")
		return
	}

	log.Print(status)

	err = models.CreateUser(&name, &email, &password)

	if err != nil {
		log.Print("ada error")
		log.Print(err)
		return
	}

	response.Status = http.StatusOK
	response.Message = "Successfully inserted"

	middlewares.Response(w, http.StatusOK, response)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var response models.ResponseLogin

	err := r.ParseForm()
	if err != nil {
		log.Print(err)
		return
	}

	email := r.Form.Get("email")
	password := middlewares.EncryptPassword(r.Form.Get("password"))

	user, err := models.GetUser(&email)
	if err != nil {
		log.Print(err)
	}

	if user.Password == password {

		sign := jwt.New(jwt.SigningMethodHS256)

		claims := sign.Claims.(jwt.MapClaims)

		claims["authorized"] = true
		claims["userid"] = user.ID
		claims["user"] = user.Name
		claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

		token, err := sign.SignedString([]byte("aqOeh4ck3R"))
		if err != nil {
			log.Print(err)
			return
		}

		response.Status = http.StatusOK
		response.Message = "Login Success"
		response.Data.ID = user.ID
		response.Data.Email = user.Email
		response.Data.Name = user.Name
		response.Data.AccessToken = token

		middlewares.Response(w, http.StatusOK, response)
	} else {

		middlewares.Error(w, 400, "Wrong Password")
	}

}
