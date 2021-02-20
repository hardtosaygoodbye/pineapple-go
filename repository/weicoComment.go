package repository

import (
	"pineapple-go/model"

	"github.com/gin-gonic/gin"
)

type weicoCommentRepository struct {
	baseRepository
}

func NewWeicoCommentRepository(ctx *gin.Context) *weicoCommentRepository {
	weicoCommentRepository := weicoCommentRepository{}
	weicoCommentRepository.New(ctx)
	return &weicoCommentRepository
}

func (wr *weicoCommentRepository) Create(comment *model.WeicoComment) error {
	return wr.db.Create(comment).Error
}

func (wr *weicoCommentRepository) FindList(weicoID int, comments *[]model.WeicoComment) error {
	return wr.db.Order("ts desc").Where("weico_id = ?", weicoID).Find(comments).Error
}

func (wr *weicoCommentRepository) Delete(commentID int) error {
	return wr.db.Where("id = ?", commentID).Delete(&model.WeicoComment{}).Error
}
