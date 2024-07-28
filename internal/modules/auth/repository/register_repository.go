package repository

import "golang-project/internal/modules/auth/entity"

func (r *Repository) Register(member entity.Member) {
	r.DB.Create(&member)
}
