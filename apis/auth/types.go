package auth

import "mime/multipart"

type RequestPayload struct {
	Login     string                `json:"login"`
	Email     string                `json:"email"`
	Password  string                `json:"password"`
	FirstName string                `json:"first_name"`
	LastName  string                `json:"last_name"`
	Avatar    *multipart.FileHeader `json:"avatar" form:"avatar"`
}

type SignUpPayload struct {
	Login     string                `json:"login" form:"login"`
	Email     string                `json:"email" form:"email"`
	Password  string                `json:"password" form:"password"`
	FirstName string                `json:"first_name" form:"first_name"`
	LastName  string                `json:"last_name" form:"last_name"`
	Avatar    *multipart.FileHeader `json:"avatar" form:"avatar"`
}

type Response struct {
	ID          uint   `json:"id"`
	Login       string `json:"login"`
	Avatar      string `json:"avatar"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	AccessToken string `json:"access_token"`
}

type ResponseSignIn struct {
	ID          uint   `json:"id"`
	Login       string `json:"login"`
	AccessToken string `json:"access_token"`
}

type GooglePayload struct {
	ID         string
	Email      string
	Picture    string
	GivenName  string `json:"given_name"`
	FamilyName string `json:"family_name"`
}
