package repository

import (
	"pineapple-go/model"

	"github.com/gin-gonic/gin"
)

//userRepository the struct of userRepository
type userRepository struct {
	baseRepository
}

// NewUserRepository create an instance of *userRepository
func NewUserRepository(ctx *gin.Context) *userRepository {
	userRepository := userRepository{}
	userRepository.New(ctx)
	return &userRepository
}

// Create user data into database
func (ur *userRepository) Create(user *model.User) error {
	return ur.db.Create(user).Error
}

// Find user record from database
func (ur *userRepository) FindWithPhone(user *model.User, phone string) error {
	return ur.db.Where("phone = ?", phone).Find(user).Error
}

func (ur *userRepository) Find(user *model.User) error {
	return ur.db.Find(user).Error
}

func (ur *userRepository) Save(user *model.User) error {
	return ur.db.Save(user).Error
}

func (ur *userRepository) FindWithWXOpenID(wxOpenID string) (user model.User, err error) {
	err = ur.db.Where("wx_open_id = ?", wxOpenID).Find(&user).Error
	return
}
