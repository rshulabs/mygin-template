package model

type User struct {
	ID
	Name     string `json:"name" gorm:"not null;comment:用户名称"`
	Mobile   string `json:"mobile" gorm:"not null;index;comment:用户手机号"`
	Password string `json:"password" gorm:"not null;default:'';comment:用户密码"`
	Email    string `json:"email" gorm:"not null;comment:邮箱"`
	Timestamps
	SoftDeletes
}
