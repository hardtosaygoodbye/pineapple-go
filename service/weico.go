package service

import (
	"errors"
	"pineapple-go/model"
	"pineapple-go/repository"
	"time"

	"github.com/gin-gonic/gin"
)

type weicoService struct {
}

var WeicoService = weicoService{}

// Publish 发布
func (ws weicoService) Publish(ctx *gin.Context, userID int, content string, urls []string, cateID int) (err error) {
	now := time.Now().Unix()
	weico := model.Weico{
		Content: content,
		UserID:  userID,
		TS:      int(now),
		CateID:  cateID,
	}
	weicoRepository := repository.NewWeicoRepository(ctx)
	err = weicoRepository.Create(&weico)
	if err != nil {
		return
	}
	if len(urls) == 0 {
		return
	}
	var weicoPics []model.WeicoPic
	for _, url := range urls {
		weicoPic := model.WeicoPic{
			WeicoID: int(weico.BaseModel.ID),
			Url:     url,
		}
		weicoPics = append(weicoPics, weicoPic)
	}
	weicoPicRepository := repository.NewWeicoPicRepository(ctx)
	err = weicoPicRepository.CreateWeicoPics(&weicoPics)
	return
}

// Delete 删除
func (ws weicoService) Delete(ctx *gin.Context, weico *model.Weico) (err error) {
	weicoRepository := repository.NewWeicoRepository(ctx)
	err = weicoRepository.Delete(weico)
	if err != nil {
		return
	}
	weicoPicRepository := repository.NewWeicoPicRepository(ctx)
	err = weicoPicRepository.DeleteWithWeicoID(int(weico.BaseModel.ID))
	return
}

// Find 查询
func (ws weicoService) Find(ctx *gin.Context, weico *model.Weico) (err error) {
	weicoRepository := repository.NewWeicoRepository(ctx)
	err = weicoRepository.Find(weico)
	return
}

// FindList 动态列表
func (ws weicoService) FindList(ctx *gin.Context, weicos *[]model.Weico, cateID int, ts int, userID int) (tmp []model.Weico, err error) {
	weicoRepository := repository.NewWeicoRepository(ctx)
	err = weicoRepository.FindList(weicos, cateID, ts)
	if err != nil {
		return
	}
	weicoPicRepository := repository.NewWeicoPicRepository(ctx)
	for _, weico := range *weicos {
		weico.IsLike, err = ws.IsLike(ctx, userID, int(weico.BaseModel.ID))
		err = weicoPicRepository.FindWithWeicoID(&weico.Pics, int(weico.BaseModel.ID))
		if err != nil {
			return
		}
		weico.User.BaseModel.ID = uint(weico.UserID)
		err = UserService.FindWithUserID(ctx, &weico.User)
		if err != nil {
			return
		}
		tmp = append(tmp, weico)
	}
	return
}

func (ws weicoService) IsLike(ctx *gin.Context, userID int, weicoID int) (isLike int, err error) {
	weicoLikeRepository := repository.NewWeicoLikeRepository(ctx)
	var weicoLike model.WeicoLike
	err = weicoLikeRepository.Find(&weicoLike, weicoID, userID)
	if err != nil {
		return
	}
	if weicoLike.BaseModel.ID != 0 {
		isLike = 1
	} else {
		isLike = 0
	}
	return
}

// Like 点赞
func (ws weicoService) Like(ctx *gin.Context, userID int, weicoID int) (err error) {
	weicoLikeRepository := repository.NewWeicoLikeRepository(ctx)
	var weicoLike model.WeicoLike
	err = weicoLikeRepository.Find(&weicoLike, weicoID, userID)
	if err != nil {
		return
	}
	if weicoLike.BaseModel.ID != 0 {
		err = errors.New("已点赞")
		return
	}
	weicoLike.UserID = userID
	weicoLike.WeicoID = weicoID
	err = weicoLikeRepository.Create(&weicoLike)
	if err != nil {
		return
	}
	weicoRepository := repository.NewWeicoRepository(ctx)
	var weico model.Weico
	weico.BaseModel.ID = uint(weicoID)
	err = weicoRepository.Find(&weico)
	if err != nil {
		return
	}
	weico.LikeNum++
	err = weicoRepository.Update(&weico)
	if err != nil {
		return
	}
	return
}

// CancelLike 取消点赞
func (ws weicoService) CancelLike(ctx *gin.Context, userID int, weicoID int) (err error) {
	weicoLikeRepository := repository.NewWeicoLikeRepository(ctx)
	var weicoLike model.WeicoLike
	err = weicoLikeRepository.Find(&weicoLike, weicoID, userID)
	if weicoLike.BaseModel.ID == 0 {
		err = errors.New("未点赞，无法取消")
		return
	}
	err = weicoLikeRepository.Delete(weicoID, userID)
	if err != nil {
		return
	}
	weicoRepository := repository.NewWeicoRepository(ctx)
	var weico model.Weico
	weico.BaseModel.ID = uint(weicoID)
	err = weicoRepository.Find(&weico)
	if err != nil {
		return
	}
	weico.LikeNum--
	err = weicoRepository.Update(&weico)
	if err != nil {
		return
	}
	return
}

// AddComment
func (ws weicoService) AddComment(ctx *gin.Context, comment *model.WeicoComment) (err error) {
	weicoCommentRepository := repository.NewWeicoCommentRepository(ctx)
	now := time.Now().Unix()
	comment.TS = int(now)
	err = weicoCommentRepository.Create(comment)

	weicoRepository := repository.NewWeicoRepository(ctx)
	var weico model.Weico
	weico.BaseModel.ID = uint(comment.WeicoID)
	err = weicoRepository.Find(&weico)
	if err != nil {
		return
	}
	weico.CommentNum++
	err = weicoRepository.Update(&weico)
	return
}

// CommentList 评论列表
func (ws weicoService) CommentList(ctx *gin.Context, weicoID int) (comments []model.WeicoComment, err error) {
	weicoCommentRepository := repository.NewWeicoCommentRepository(ctx)
	var tmp []model.WeicoComment
	err = weicoCommentRepository.FindList(weicoID, &tmp)
	for _, comment := range tmp {
		comment.FromUser.BaseModel.ID = uint(comment.FromUserID)
		err = UserService.FindWithUserID(ctx, &comment.FromUser)
		if err != nil {
			return
		}
		comments = append(comments, comment)
	}
	return
}

// DeleteComment 删除评论
func (ws weicoService) DeleteComment(ctx *gin.Context, commentID int, weicoID int) (err error) {
	weicoCommentRepository := repository.NewWeicoCommentRepository(ctx)
	err = weicoCommentRepository.Delete(commentID)

	weicoRepository := repository.NewWeicoRepository(ctx)
	var weico model.Weico
	weico.BaseModel.ID = uint(weicoID)
	err = weicoRepository.Find(&weico)
	if err != nil {
		return
	}
	weico.CommentNum--
	err = weicoRepository.Update(&weico)
	return
}

// WeicoCateList 动态分类
func (ws weicoService) WeicoCateList(ctx *gin.Context) (weicoCates []model.WeicoCate, err error) {
	weicoCateRepository := repository.NewWeicoCateRepository(ctx)
	return weicoCateRepository.FindList()
}
