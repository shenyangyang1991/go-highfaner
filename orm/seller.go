package orm

import (
	"errors"

	"github.com/jinzhu/gorm"
)

// Seller ...
type Seller struct {
	gorm.Model
	Username string `description:"用户名"`
	Password string `description:"密码"`
	State    bool   `description:"状态"`
	Remark   string `description:"备注"`
}

// IsExist ...
func (s *Seller) IsExist(db *gorm.DB) bool {
	return !(db.Where("username = ?", s.Username).First(&Seller{}).RecordNotFound())
}

// BeforeCreate ...
func (s *Seller) BeforeCreate(scope *gorm.Scope) error {
	if s.IsExist(scope.DB()) {
		return errors.New("您的账号已注册")
	}

	return nil
}

// Create ...
func (s *Seller) Create(db *gorm.DB) error {
	result := db.Create(s)
	return result.Error
}

// Find ...
func (s *Seller) Find(db *gorm.DB) error {
	result := db.Where("username = ? and password = ?", s.Username, s.Password).Find(s)
	return result.Error
}

// AfterFind ...
func (s *Seller) AfterFind() error {
	if !s.State {
		return errors.New(s.Remark)
	}

	return nil
}
