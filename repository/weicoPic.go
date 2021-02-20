package repository

import (
	"pineapple-go/model"

	"github.com/gin-gonic/gin"
)

type weicoPicRepository struct {
	baseRepository
}

func NewWeicoPicRepository(ctx *gin.Context) *weicoPicRepository {
	weicoPicRepository := weicoPicRepository{}
	weicoPicRepository.New(ctx)
	return &weicoPicRepository
}

func (wr *weicoPicRepository) CreateWeicoPics(weicoPics *[]model.WeicoPic) error {
	return wr.db.Create(weicoPics).Error
}

func (wr *weicoPicRepository) DeleteWithWeicoID(weicoID int) error {
	return wr.db.Where("weico_id = ?", weicoID).Delete(&model.WeicoPic{}).Error
}

func (wr *weicoPicRepository) FindWithWeicoID(weicoPics *[]model.WeicoPic, weicoID int) error {
	return wr.db.Where("weico_id = ?", weicoID).Find(weicoPics).Error
}
