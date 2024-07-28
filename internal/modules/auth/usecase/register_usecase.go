package usecase

import (
	"golang-project/internal/modules/auth/entity"
	"golang-project/internal/modules/auth/model"
)

func (u *Usecase) Register(req model.RegisterMemberRequestModel) {

	member := entity.Member{
		Email:    req.Email,
		Password: req.Password,
		NoTlp:    req.NoTlp,
	}

	u.Repository.Register(member)
}
