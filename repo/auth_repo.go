package repo

import (
	"authentication/model"

	"gorm.io/gorm"
)

type AuthRepo interface {
	CheckUsername(username string) bool
	Register(user *model.User) error
	Login(username string) (string, error)
	CheckID(id int) bool
	Delete(id int) error
}

type authRepo struct {
	db *gorm.DB
}

func NewAuthRepo(db *gorm.DB) AuthRepo {
	return &authRepo{db: db}
}

func (r *authRepo) CheckUsername(username string) bool {
	var count int64
	if err := r.db.Table("users").Where("username = ?", username).Count(&count).Error; err != nil {
		return false
	}

	if count > 0 {
		return false
	}

	return true
}

func (r *authRepo) Register(user *model.User) error {
	if err := r.db.Table("users").Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func (r *authRepo) Login(username string) (string, error) {
	var user model.User
	if err := r.db.Table("users").Where("username = ?", username).First(&user).Error; err != nil {
		return "", err
	}

	return user.Password, nil
}

func (r *authRepo) CheckID(id int) bool {
	var count int64
	if err := r.db.Table("users").Where("id = ?", id).Count(&count).Error; err != nil {
		return false
	}

	if count < 1 {
		return false
	}

	return true
}

func (r *authRepo) Delete(id int) error {
	if err := r.db.Delete(&model.User{}, id).Error; err != nil {
		return err
	}

	return nil
}
