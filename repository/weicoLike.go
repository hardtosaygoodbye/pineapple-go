package repository

import (
	"pineapple-go/model"

	"github.com/gin-gonic/gin"
)

type weicoLikeRepository struct {
	baseRepository
}

func NewWeicoLikeRepository(ctx *gin.Context) *weicoLikeRepository {
	weicoLikeRepository := weicoLikeRepository{}
	weicoLikeRepository.New(ctx)
	return &weicoLikeRepository
}

func (wr *weicoLikeRepository) Find(weicoLike *model.WeicoLike, weicoID int, userID int) error {
	return wr.db.Where("weico_id = ? and user_id = ?", weicoID, userID).Find(weicoLike).Error
}

func (wr *weicoLikeRepository) Create(weicoLike *model.WeicoLike) error {
	return wr.db.Create(weicoLike).Error
}

func (wr *weicoLikeRepository) Delete(weicoID int, userID int) error {
	return wr.db.Where("weico_id = ? and user_id = ?", weicoID, userID).Delete(&model.WeicoLike{}).Error
}
