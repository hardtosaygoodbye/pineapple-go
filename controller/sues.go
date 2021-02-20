package controller

import (
	"pineapple-go/service"

	"github.com/gin-gonic/gin"
)

// ClientConnect 客户端连接
func ClientConnect(ctx *gin.Context) {
	clientID := ctx.Query("client_id")
	_, err := service.SuesService.UpgradeContext(clientID, ctx)
	if err != nil {
		ErrorWithMsg(ctx, err.Error())
		return
	}
}

// SuesCourse 课表
func SuesCourse(ctx *gin.Context) {
	stdno := ctx.PostForm("stdno")
	password := ctx.PostForm("password")
	var (
		err     error
		cookie  string
		captcha string
	)
	i := 5
	for i > 0 {
		i--
		err = nil
		captcha, cookie, err = service.SuesService.GetCaptchaAndCookie()
		if err != nil {
			continue
		}
		err = service.SuesService.LoginJxxt(stdno, password, captcha, cookie)
		if err == nil {
			break
		}
	}
	if err != nil {
		ErrorWithMsg(ctx, err.Error())
		return
	}
	stdID, err := service.SuesService.GetStdID(cookie)
	if err != nil {
		ErrorWithMsg(ctx, err.Error())
		return
	}
	courses, err := service.SuesService.GetCourseTable(cookie, stdID)
	if err != nil {
		ErrorWithMsg(ctx, err.Error())
		return
	}
	data := make(map[string]interface{})
	data["courses"] = courses
	Success(ctx, data)
}
