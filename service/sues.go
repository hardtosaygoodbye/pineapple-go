package service

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"pineapple-go/core/req"
	"pineapple-go/model"
	"pineapple-go/util"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type suesService struct {
}

var SuesService = suesService{}

var proxy req.Proxy = "http://118.25.210.52:8080"

// GetCaptchaAndCookie 获取cookie
func (ss suesService) GetCaptchaAndCookie() (captcha string, cookie string, err error) {
	resp, err := req.Get("http://jxxt.sues.edu.cn/eams/captcha/image.action", proxy)
	if err != nil {
		return
	}
	// 获取cookie
	cookies := resp.Header["Set-Cookie"]
	JSESSIONID := strings.Split(cookies[0], ";")[0]
	test := strings.Split(cookies[1], ";")[0]
	cookie = JSESSIONID + ";popped='';" + test
	fileName := uuid.New().String()
	util.SaveDataToFile(resp.Data, fileName)
	defer os.Remove(fileName)
	captcha, err = util.IdentifyCaptcha(fileName)
	return
}

// LoginJxxt 登录教学管理系统
func (ss suesService) LoginJxxt(stdno, password, captcha, cookie string) error {
	resp, err := req.Post(
		"http://jxxt.sues.edu.cn/eams/login.action",
		req.FormParam{
			"loginForm.name":     stdno,
			"loginForm.password": password,
			"loginForm.captcha":  captcha,
			"encodedPassword":    "",
		},
		req.Header{
			"Cookie":       cookie,
			"Content-Type": "application/x-www-form-urlencoded",
		},
		proxy,
	)
	if err != nil {
		return err
	}
	fmt.Println(string(resp.Data))
	if strings.Contains(string(resp.Data), "Wrong Catcha String") {
		return errors.New("验证码错误")
	} else if strings.Contains(string(resp.Data), "Error Password") {
		return errors.New("密码错误")
	}
	return nil
}

// GetStdID 获取学生ID
func (ss suesService) GetStdID(cookie string) (stdID string, err error) {
	resp, err := req.Get(
		"http://jxxt.sues.edu.cn/eams/courseTableForStd.action?method=stdHome",
		req.Header{
			"Cookie":       cookie,
			"Referer":      "http://jxxt.sues.edu.cn/eams/defaultHome.action?method=moduleList&parentCode=",
			"Content-Type": "application/x-www-form-urlencoded",
		},
		proxy)
	if err != nil {
		return "", nil
	}
	fmt.Println(string(resp.Data))
	stdID, err = util.SplitContent(string(resp.Data), "javascript:getCourseTable('std','", "',event)")
	return
}

// GetCourseTable 获取课表数据
func (ss suesService) GetCourseTable(cookie, stdID string) (courses []model.Course, err error) {
	courses = make([]model.Course, 0)
	resp, err := req.Get(
		"http://jxxt.sues.edu.cn/eams/courseTableForStd.action",
		req.QueryParam{
			"method":              "courseTable",
			"setting.forSemester": 1,
			"setting.kind":        "std",
			"semester.id":         "462",
			"ids":                 stdID,
			"ignoreHead":          1,
		},
		req.Header{
			"Cookie":       cookie,
			"Content-Type": "application/x-www-form-urlencoded",
		},
		proxy)
	if err != nil {
		return nil, err
	}
	temp, err := util.SplitContent(string(resp.Data), "var activity=null;", "table0.marshalTable")
	courseStrs := strings.Split(temp, "activity = new TaskActivity(")
	for i, class := range courseStrs {
		if i == 0 {
			continue
		}
		lines := strings.Split(class, "\n")
		var course model.Course
		for _, line := range lines {
			if len(line) > 80 {
				// 课程
				courseStrArr := strings.Split(line, "\"")
				course.Teacher = courseStrArr[3]
				course.Name = courseStrArr[7]
				course.Address = courseStrArr[11]
				course.Week = courseStrArr[13]
			} else if len(line) < 30 && len(line) > 10 {
				// 星期和节数
				course.Index, _ = strconv.Atoi(string(line[8]))
				if course.Week != "" {
					course.Time = course.Time + ","
				}
				// index =2*unitCount+7;
				oneTime, _ := util.SplitContent(line, "+", ";\r")
				course.Time = course.Time + oneTime
			}
		}
		course.Time = course.Time[1:]
		courses = append(courses, course)
	}
	return
}

/*  以下为备用方案  */
var onlineConns map[string]*websocket.Conn = make(map[string]*websocket.Conn)

// Configure the upgrader
var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// UpgradeContext 升级客户端连接
func (ss suesService) UpgradeContext(clientID string, ctx *gin.Context) (conn *websocket.Conn, err error) {
	conn, err = upGrader.Upgrade(ctx.Writer, ctx.Request, ctx.Writer.Header())
	if err != nil {
		return nil, err
	}
	onlineConns[clientID] = conn
	return conn, err
}

// ClientRequest 内网服务器代理请求
func (ss suesService) ClientRequest(url string) []byte {
	for _, tmpConn := range onlineConns {
		tmpConn.WriteJSON(gin.H{
			"url": url,
		})
		_, msg, err := tmpConn.ReadMessage()
		if err != nil {
			fmt.Println(err.Error())
			return nil
		}
		return msg
	}
	return nil
}
