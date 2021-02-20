package util

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

// IdentifyCaptcha 验证码图片识别
func IdentifyCaptcha(fileName string) (captcha string, err error) {
	exec.Command("/bin/bash", "-c", fmt.Sprintf("tesseract %s %s -l eng", fileName, fileName)).Run()
	captchaBytes, err := ioutil.ReadFile(fileName + ".txt")
	if err != nil {
		return
	}
	defer os.Remove(fileName + ".txt")
	// 验证码清洗
	captcha = string(captchaBytes)
	captcha = strings.Split(captcha, "\n")[0]
	captcha = strings.Replace(captcha, " ", "", -1)
	isValid, err := regexp.MatchString("^[a-z]{4,5}$", captcha)
	if err != nil {
		return
	}
	if !isValid {
		err = errors.New("验证码格式错误")
	}
	return
}
