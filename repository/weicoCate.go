package repository

import (
	"pineapple-go/model"

	"github.com/gin-gonic/gin"
)

type weicoCateRepository struct {
	baseRepository
}

func NewWeicoCateRepository(ctx *gin.Context) *weicoCateRepository {
	weicoCateRepository := weicoCateRepository{}
	weicoCateRepository.New(ctx)
	return &weicoCateRepository
}

func (wr *weicoCateRepository) FindList() (weicoCates []model.WeicoCate, err error) {
	err = wr.db.Find(&weicoCates).Error
	return
}
