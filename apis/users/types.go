package users

import (
	"mime/multipart"
)

type RequestPayload struct {
	Login     string                `json:"login" form:"login"`
	Email     string                `json:"email" form:"email"`
	Password  string                `json:"password"`
	FirstName string                `json:"first_name"`
	LastName  string                `json:"last_name"`
	Avatar    *multipart.FileHeader `json:"avatar" form:"avatar"`
}

type UpdateResponse struct {
	ID          uint   `json:"id"`
	Login       string `json:"login"`
	Avatar      string `json:"avatar"`
	Email       string `json:"email"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	AccessToken string `json:"access_token"`
}
