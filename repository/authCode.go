package repository

import (
	"pineapple-go/model"

	"github.com/gin-gonic/gin"
)

type authCodeRepository struct {
	baseRepository
}

func NewAuthCodeRepository(ctx *gin.Context) *authCodeRepository {
	authCodeRepository := authCodeRepository{}
	authCodeRepository.New(ctx)
	return &authCodeRepository
}

func (ar *authCodeRepository) Create(authCode *model.AuthCode) error {
	return ar.db.Create(authCode).Error
}

func (ar *authCodeRepository) FindWithPhone(authCode *model.AuthCode, phone string) error {
	return ar.db.Where("phone = ?", phone).Find(authCode).Error
}

func (ar *authCodeRepository) UpdateCode(authCode *model.AuthCode, code string) error {
	return ar.db.Model(authCode).Update("code", code).Error
}

func (ar *authCodeRepository) FindWithPhoneAndCode(authCode *model.AuthCode, phone string, code string) error {
	return ar.db.Where("phone = ? and code = ?", phone, code).Find(authCode).Error
}
