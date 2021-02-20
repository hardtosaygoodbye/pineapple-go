package service

import (
	"encoding/json"
	"pineapple-go/config"
	"pineapple-go/core/log"
	"pineapple-go/core/req"
	"pineapple-go/model"
	"pineapple-go/repository"
	"pineapple-go/util"

	"github.com/gin-gonic/gin"
)

type userService struct {
}

// UserService create userService instance
var UserService = userService{}

func (us userService) Create(ctx *gin.Context, user *model.User) error {
	userRepository := repository.NewUserRepository(ctx)
	log.Info(ctx, "create user", user)
	return userRepository.Create(user)
}

func (us userService) FindWithPhone(ctx *gin.Context, user *model.User, phone string) error {
	userRepository := repository.NewUserRepository(ctx)
	return userRepository.FindWithPhone(user, phone)
}

func (us userService) FindWithUserID(ctx *gin.Context, user *model.User) error {
	userRepository := repository.NewUserRepository(ctx)
	return userRepository.Find(user)
}

func (us userService) FindWithWXOpenID(ctx *gin.Context, wxOpenID string) (user model.User, err error) {
	userRepository := repository.NewUserRepository(ctx)
	return userRepository.FindWithWXOpenID(wxOpenID)
}

func (us userService) Save(ctx *gin.Context, user *model.User) error {
	userRepository := repository.NewUserRepository(ctx)
	return userRepository.Save(user)
}

func (us userService) RefreshAuthCode(ctx *gin.Context, phone string) error {
	authRepository := repository.NewAuthCodeRepository(ctx)
	var authCode model.AuthCode
	err := authRepository.FindWithPhone(&authCode, phone)
	if err != nil {
		return err
	}
	code := util.CreateRandomNumber(6)
	if authCode.BaseModel.ID == 0 {
		authCode = model.AuthCode{
			Phone: phone,
			Code:  code,
		}
		err = authRepository.Create(&authCode)
		if err != nil {
			return err
		}
	} else {
		err = authRepository.UpdateCode(&authCode, code)
		if err != nil {
			return err
		}
	}
	return nil
}

func (us userService) CheckAuthCode(ctx *gin.Context, phone string, code string) (correct bool, err error) {
	authRespository := repository.NewAuthCodeRepository(ctx)
	var authCode model.AuthCode
	err = authRespository.FindWithPhoneAndCode(&authCode, phone, code)
	if err != nil {
		return
	}
	correct = authCode.BaseModel.ID != 0
	return
}

func (us userService) WxCode2openid(code string) (openid string, err error) {
	resp, err := req.Get(
		"https://api.weixin.qq.com/sns/jscode2session",
		req.QueryParam{
			"appid":      config.Wechat.AppID,
			"secret":     config.Wechat.AppSecret,
			"js_code":    code,
			"grant_type": "authorization_code",
		},
	)
	if err != nil {
		return "", err
	}
	respMap := make(map[string]string)
	err = json.Unmarshal(resp.Data, &respMap)
	if err != nil {
		return "", err
	}
	openid, ok := respMap["openid"]
	if !ok {
		return "", nil
	}
	return openid, nil
}
