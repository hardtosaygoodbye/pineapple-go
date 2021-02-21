package repository

import (
	"pineapple-go/model"

	"github.com/gin-gonic/gin"
)

type weicoRepository struct {
	baseRepository
}

func NewWeicoRepository(ctx *gin.Context) *weicoRepository {
	weicoRepository := weicoRepository{}
	weicoRepository.New(ctx)
	return &weicoRepository
}

func (wr *weicoRepository) Create(weico *model.Weico) error {
	return wr.db.Create(weico).Error
}

func (wr *weicoRepository) FindList(weicos *[]model.Weico, cateID int) error {
	if cateID != 0 {
		return wr.db.Where("cate_id = ?", cateID).Order("publish_ts desc").Find(weicos).Error
	} else {
		return wr.db.Order("publish_ts desc").Find(weicos).Error
	}
}

func (wr *weicoRepository) Find(weico *model.Weico) error {
	return wr.db.Find(weico).Error
}

func (wr *weicoRepository) Delete(weico *model.Weico) error {
	return wr.db.Delete(weico).Error
}

func (wr *weicoRepository) Update(weico *model.Weico) error {
	return wr.db.Save(weico).Error
}
