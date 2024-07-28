package model

type RegisterMemberRequestModel struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
	NoTlp    string `json:"no_tlp" binding:"required"`
}

type RegisterMemberResponseModel struct {
}
