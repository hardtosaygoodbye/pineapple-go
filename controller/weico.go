package controller

import (
	"pineapple-go/model"
	"pineapple-go/service"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
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
	cateID := cast.ToInt(ctx.PostForm("cate_id"))
	err := service.WeicoService.Publish(ctx, uid, content, picArr, cateID)
	if err != nil {
		ErrorWithMsg(ctx, err.Error())
		return
	}
	Success(ctx, gin.H{})
}

// WeicoList 动态列表
func WeicoList(ctx *gin.Context) {
	uid := ctx.GetInt("uid")
	cateID := cast.ToInt(ctx.PostForm("cate_id"))
	ts := cast.ToInt(ctx.PostForm("ts"))
	var weicos []model.Weico
	weicos, err := service.WeicoService.FindList(ctx, &weicos, cateID, ts, uid)
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
	weicoID := cast.ToInt(ctx.PostForm("weico_id"))
	weico := model.Weico{
		BaseModel: model.BaseModel{
			ID: uint(weicoID),
		},
	}
	err := service.WeicoService.Find(ctx, &weico)
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
	like := cast.ToInt(ctx.PostForm("like"))
	weicoID := cast.ToInt(ctx.PostForm("weico_id"))
	uid := ctx.GetInt("uid")
	var err error
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
	toUserID := cast.ToInt(ctx.PostForm("to_user_id"))
	weicoID := cast.ToInt(ctx.PostForm("weico_id"))
	comment := model.WeicoComment{
		Content:    content,
		FromUserID: uid,
		ToUserID:   toUserID,
		WeicoID:    weicoID,
	}
	err := service.WeicoService.AddComment(ctx, &comment)
	if err != nil {
		ErrorWithMsg(ctx, err.Error())
		return
	}
	Success(ctx, gin.H{})
}

// DeleteComment 删除评论
func DeleteComment(ctx *gin.Context) {
	commentID := cast.ToInt(ctx.PostForm("comment_id"))
	weicoID := cast.ToInt(ctx.PostForm("weico_id"))
	err := service.WeicoService.DeleteComment(ctx, commentID, weicoID)
	if err != nil {
		ErrorWithMsg(ctx, err.Error())
		return
	}
	Success(ctx, gin.H{})
}

// CommentList 评论列表
func CommentList(ctx *gin.Context) {
	weicoID := cast.ToInt(ctx.PostForm("weico_id"))
	comments, err := service.WeicoService.CommentList(ctx, weicoID)
	if err != nil {
		ErrorWithMsg(ctx, err.Error())
		return
	}
	Success(ctx, gin.H{
		"comments": comments,
	})
}

// WeicoCateList 动态分类
func WeicoCateList(ctx *gin.Context) {
	weicoCates, err := service.WeicoService.WeicoCateList(ctx)
	if err != nil {
		ErrorWithMsg(ctx, err.Error())
		return
	}
	var tmp []model.WeicoCate
	tmp = append(tmp, model.WeicoCate{
		Content: "全部",
	})
	tmp = append(tmp, weicoCates...)
	Success(ctx, gin.H{
		"cates": tmp,
	})
}
