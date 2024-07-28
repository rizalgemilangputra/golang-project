package entity

type Member struct {
	ID       int64  `gorm:"column:id;primaryKey;default:nextval('member_id_seq'::regclass)"`
	Email    string `gorm:"column:email"`
	Password string `gorm:"column:password"`
	NoTlp    string `gorm:"column:no_tlp"`
}

func (m *Member) TableName() string {
	return "members"
}
