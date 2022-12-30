package users

import "mime/multipart"

type RequestPayload struct {
	Login     string                `json:"login"`
	Email     string                `json:"email"`
	Password  string                `json:"password"`
	FirstName string                `json:"first_name"`
	LastName  string                `json:"last_name"`
	Avatar    *multipart.FileHeader `json:"avatar"`
}
