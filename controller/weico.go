package controller

import (
	"pineapple-go/model"
	"pineapple-go/service"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// PublishWeico 发布动态
func PublishWeico(ctx *gin.Context) {
	uid := ctx.GetInt("uid")
	content := ctx.PostForm("content")
	pics := ctx.PostForm("pics")
	picArr := strings.Split(pics, "|")
	if len(pics) == 0 {
		picArr = make([]string, 0)
	}
	err := service.WeicoService.Publish(ctx, uid, content, picArr)
	if err != nil {
		ErrorWithMsg(ctx, err.Error())
		return
	}
	Success(ctx, gin.H{})
}

// WeicoList 动态列表
func WeicoList(ctx *gin.Context) {
	var weicos []model.Weico
	weicos, err := service.WeicoService.FindList(ctx, &weicos)
	if err != nil {
		ErrorWithMsg(ctx, err.Error())
		return
	}
	Success(ctx, gin.H{
		"weicos": weicos,
	})
}

// DeleteWeico 删除动态
func DeleteWeico(ctx *gin.Context) {
	uid := ctx.GetInt("uid")
	weicoIDStr := ctx.PostForm("weico_id")
	weicoID, err := strconv.Atoi(weicoIDStr)
	weico := model.Weico{
		BaseModel: model.BaseModel{
			ID: uint(weicoID),
		},
	}
	err = service.WeicoService.Find(ctx, &weico)
	if err != nil {
		ErrorWithMsg(ctx, err.Error())
		return
	}
	if weico.UserID == 0 {
		ErrorWithMsg(ctx, "动态不存在")
		return
	}
	if weico.UserID != uid {
		ErrorWithMsg(ctx, "无权限删除该动态")
		return
	}
	err = service.WeicoService.Delete(ctx, &weico)
	if err != nil {
		ErrorWithMsg(ctx, err.Error())
		return
	}
	Success(ctx, gin.H{})
}

// LikeWeico 点赞动态
func LikeWeico(ctx *gin.Context) {
	likeStr := ctx.PostForm("like")
	weicoIDStr := ctx.PostForm("weico_id")
	weicoID, err := strconv.Atoi(weicoIDStr)
	uid := ctx.GetInt("uid")
	like, err := strconv.Atoi(likeStr)
	if like == 1 {
		err = service.WeicoService.Like(ctx, uid, weicoID)
	} else if like == 0 {
		err = service.WeicoService.CancelLike(ctx, uid, weicoID)
	}
	if err != nil {
		ErrorWithMsg(ctx, err.Error())
		return
	}
	Success(ctx, gin.H{})
}

// CommentWeico 评论动态
func CommentWeico(ctx *gin.Context) {
	content := ctx.PostForm("content")
	uid := ctx.GetInt("uid")
	toUserIDStr := ctx.PostForm("to_user_id")
	weicoIDStr := ctx.PostForm("weico_id")
	if len(toUserIDStr) == 0 {
		toUserIDStr = "0"
	}
	toUserID, err := strconv.Atoi(toUserIDStr)
	if err != nil {
		ErrorWithMsg(ctx, err.Error())
		return
	}
	weicoID, err := strconv.Atoi(weicoIDStr)
	if err != nil {
		ErrorWithMsg(ctx, err.Error())
		return
	}
	comment := model.WeicoComment{
		Content:    content,
		FromUserID: uid,
		ToUserID:   toUserID,
		WeicoID:    weicoID,
	}
	err = service.WeicoService.AddComment(ctx, &comment)
	if err != nil {
		ErrorWithMsg(ctx, err.Error())
		return
	}
	Success(ctx, gin.H{})
}

// DeleteComment 删除评论
func DeleteComment(ctx *gin.Context) {
	commentIDStr := ctx.PostForm("comment_id")
	commentID, err := strconv.Atoi(commentIDStr)
	if err != nil {
		ErrorWithMsg(ctx, err.Error())
		return
	}
	err = service.WeicoService.DeleteComment(ctx, commentID)
	if err != nil {
		ErrorWithMsg(ctx, err.Error())
		return
	}
	Success(ctx, gin.H{})
}

// CommentList 评论列表
func CommentList(ctx *gin.Context) {
	weicoIDStr := ctx.PostForm("weico_id")
	weicoID, err := strconv.Atoi(weicoIDStr)
	if err != nil {
		ErrorWithMsg(ctx, err.Error())
		return
	}
	comments, err := service.WeicoService.CommentList(ctx, weicoID)
	if err != nil {
		ErrorWithMsg(ctx, err.Error())
		return
	}
	Success(ctx, gin.H{
		"comments": comments,
	})
}
