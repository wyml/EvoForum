package user

import (
	"forum/app/models"
	"forum/pkg/database"
	"forum/pkg/hash"
)

// User 用户模型
type User struct {
	models.BaseModel

	Name     string `json:"name,omitempty"`
	Email    string `json:"-"`
	Phone    string `json:"-"`
	Password string `json:"-"`

	models.CommonTimestampsField
}

// Create 创建用户，通过 User.ID 来判断是否创建成功
func (u *User) Create() {
	database.DB.Create(&u)
}

// ComparePassword 密码是否正确
func (u *User) ComparePassword(_password string) bool {
	return hash.BcryptCheck(_password, u.Password)
}

func (u *User) Save() (rowsAffected int64) {
	result := database.DB.Save(&u)
	return result.RowsAffected
}
