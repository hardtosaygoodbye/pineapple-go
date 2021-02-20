package controller

import (
	"pineapple-go/config"
	"pineapple-go/constant"
	"pineapple-go/core/redis"
	"pineapple-go/model"
	"pineapple-go/service"
	"pineapple-go/util"
	"strconv"

	"github.com/gin-gonic/gin"
)

// WxLogin 微信登录
func WxLogin(ctx *gin.Context) {
	code := ctx.PostForm("code")
	if len(code) == 0 {
		ErrorWithMsg(ctx, "参数code缺失")
		return
	}
	openid, err := service.UserService.WxCode2openid(code)
	if err != nil {
		ErrorWithMsg(ctx, err.Error())
		return
	}
	user, err := service.UserService.FindWithWXOpenID(ctx, openid)
	if err != nil {
		ErrorWithMsg(ctx, err.Error())
		return
	}
	isNewUser := 0
	if user.BaseModel.ID == 0 {
		isNewUser = 1
		user.WxOpenID = openid
		err = service.UserService.Create(ctx, &user)
		if err != nil {
			ErrorWithMsg(ctx, err.Error())
			return
		}
	}
	accessToken, err := util.CreateToken(user.ID, config.App.JWT_TOKEN)
	if err != nil {
		Error(ctx, constant.USER_JWT_ERROR, "登录生成jwt token失败")
		return
	}
	err = redis.Client.HMSet(ctx, "jwt:user:"+strconv.Itoa(int(user.ID)),
		map[string]interface{}{
			"userID": user.ID,
		},
	).Err()
	if err != nil {
		ErrorWithMsg(ctx, err.Error())
		return
	}
	Success(ctx, gin.H{"token": accessToken, "is_new": isNewUser, "uid": user.BaseModel.ID})
}

// AuthCode 验证码
func AuthCode(ctx *gin.Context) {
	phone := ctx.PostForm("phone")
	err := service.UserService.RefreshAuthCode(ctx, phone)
	if err != nil {
		ErrorWithMsg(ctx, err.Error())
		return
	}
	Success(ctx, gin.H{})
}

// PhoneLogin 手机号登录
func PhoneLogin(ctx *gin.Context) {
	phone := ctx.PostForm("phone")
	code := ctx.PostForm("code")
	correct, err := service.UserService.CheckAuthCode(ctx, phone, code)
	if err != nil {
		ErrorWithMsg(ctx, err.Error())
		return
	}
	if !correct {
		ErrorWithMsg(ctx, "验证码错误")
		return
	}
	var user model.User
	err = service.UserService.FindWithPhone(ctx, &user, phone)
	if err != nil {
		ErrorWithMsg(ctx, err.Error())
		return
	}
	if user.BaseModel.ID == 0 {
		user = model.User{
			Phone: phone,
		}
		err = service.UserService.Create(ctx, &user)
		if err != nil {
			ErrorWithMsg(ctx, err.Error())
			return
		}
	}
	token, err := util.CreateToken(user.BaseModel.ID, config.App.JWT_TOKEN)
	if err != nil {
		ErrorWithMsg(ctx, err.Error())
		return
	}
	err = redis.Client.HMSet(ctx, "jwt:user:"+strconv.Itoa(int(user.ID)),
		map[string]interface{}{
			"userID": user.ID,
		},
	).Err()
	if err != nil {
		ErrorWithMsg(ctx, err.Error())
		return
	}
	Success(ctx, gin.H{"token": token})
}

func UserInfo(ctx *gin.Context) {
	var user model.User
	user.BaseModel.ID = uint(ctx.GetInt("uid"))
	err := service.UserService.FindWithUserID(ctx, &user)
	if err != nil {
		ErrorWithMsg(ctx, err.Error())
		return
	}
	Success(ctx, gin.H{
		"user_info": user,
	})
}

func UpdateUserInfo(ctx *gin.Context) {
	var user model.User
	user.BaseModel.ID = uint(ctx.GetInt("uid"))
	err := service.UserService.FindWithUserID(ctx, &user)
	if err != nil {
		ErrorWithMsg(ctx, err.Error())
		return
	}
	phone := ctx.PostForm("phone")
	if len(phone) != 0 {
		user.Phone = phone
	}
	avatar := ctx.PostForm("avatarUrl")
	if len(avatar) != 0 {
		user.Avatar = avatar
	}
	gender := ctx.PostForm("gender")
	if len(gender) != 0 {
		user.Gender = gender
	}
	language := ctx.PostForm("language")
	if len(language) != 0 {
		user.Language = language
	}
	nickName := ctx.PostForm("nickName")
	if len(nickName) != 0 {
		user.NickName = nickName
	}
	country := ctx.PostForm("country")
	if len(country) != 0 {
		user.Country = country
	}
	province := ctx.PostForm("province")
	if len(province) != 0 {
		user.Province = province
	}
	city := ctx.PostForm("city")
	if len(city) != 0 {
		user.City = city
	}
	err = service.UserService.Save(ctx, &user)
	if err != nil {
		ErrorWithMsg(ctx, err.Error())
		return
	}
	Success(ctx, gin.H{})
}

// func Login(ctx *gin.Context) {

// 	var user model.User
// 	if err := ctx.ShouldBind(&user); err != nil {
// 		ErrorWithMessage(ctx, constant.USER_LOGIN_FAILED, "登录失败")
// 		return
// 	}
// 	err := service.UserService.Find(ctx, &user)
// 	log.Info(ctx, "user", user)
// 	if err != nil {
// 		ErrorWithMessage(ctx, constant.USER_NOT_EXISTS, "用户不存在")
// 		return
// 	}
// 	access_token, err := util.CreateToken(user.ID, config.App.JWT_TOKEN)
// 	if err != nil {
// 		Error(ctx, constant.USER_JWT_ERROR)
// 		return
// 	}
// 	err = redis.Client.HMSet(ctx, "jwt:user:"+strconv.Itoa(int(user.ID)),
// 		map[string]interface{}{
// 			"userId":   user.ID,
// 			"username": user.Username,
// 		},
// 	).Err()

// 	if err != nil {
// 		log.Info(ctx, "login jwt", err)
// 		ErrorWithMessage(ctx, constant.REDIS_ERROR, err.Error())
// 		return
// 	}
// 	ctx.Writer.Header().Set("Authentication", access_token)
// 	ctx.Writer.Header().Set("TraceId", uuid.New().String())
// 	Success(ctx, "登录成功", gin.H{"info": user})

// }

// func UserInfo(ctx *gin.Context) {
// 	userId, _ := ctx.Get("userId")
// 	Success(ctx, "用户信息如下", gin.H{"userId": userId})
// }
