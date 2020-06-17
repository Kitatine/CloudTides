package handler

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-openapi/runtime/middleware"

	"tides-server/pkg/restapi/operations/user"

	"tides-server/pkg/config"
	"tides-server/pkg/models"
)

func RegisterUserHandler(params user.RegisterUserParams) middleware.Responder {
	body := params.ReqBody
	db := config.GetDB()
	var queryUser models.User
	db.Where("username = ?", body.Username).First(&queryUser)
	if queryUser.Username != "" {
		return user.NewRegisterUserBadRequest().WithPayload(&user.RegisterUserBadRequestBody{Message: "Username already used!"})
	}

	res := &user.RegisterUserOKBodyUserInfo{
		CompanyName: body.CompanyName,
		Password:    body.Password,
		Priority:    body.Priority,
		Username:    body.Username,
	}
	newUser := models.User{Username: body.Username, Password: body.Password, CompanyName: body.CompanyName, Priority: body.Priority}

	err := db.Create(&newUser).Error
	if err != nil {
		return user.NewRegisterUserBadRequest()
	}

	return user.NewRegisterUserOK().WithPayload(&user.RegisterUserOKBody{UserInfo: res})
}

func UserLoginHandler(params user.UserLoginParams) middleware.Responder {
	body := params.ReqBody

	db := config.GetDB()
	var queryUser models.User
	db.Where("Username = ?", body.Username).First(&queryUser)
	if queryUser.Username == "" || queryUser.Password != body.Password {
		return user.NewUserLoginUnauthorized()
	}
	expirationTime := time.Now().Add(expireTime)
	claims := Claims{
		Id: queryUser.Model.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Issuer:    issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := token.SignedString([]byte(secretKey))

	res := user.UserLoginOKBodyUserInfo{Priority: queryUser.Priority, Username: queryUser.Username}
	return user.NewUserLoginOK().WithPayload(&user.UserLoginOKBody{Token: signedToken, UserInfo: &res})
}

func UserDetailsHandler(params user.UserDetailsParams) middleware.Responder {
	id, err := ParseUserIdFromToken(params.HTTPRequest)
	if err != nil {
		return user.NewUserDetailsUnauthorized()
	}
	db := config.GetDB()
	var queryUser models.User
	db.Where("id = ?", id).First(&queryUser)
	if queryUser.Username == "" {
		return user.NewUserDetailsUnauthorized()
	}

	res := user.UserDetailsOKBodyResults{
		City:        queryUser.City,
		CompanyName: queryUser.CompanyName,
		Country:     queryUser.Country,
		Email:       queryUser.Email,
		FirstName:   queryUser.FirstName,
		LastName:    queryUser.LastName,
		Position:    queryUser.Position,
	}

	return user.NewUserDetailsOK().WithPayload(&user.UserDetailsOKBody{Message: "success", Results: &res})
}
