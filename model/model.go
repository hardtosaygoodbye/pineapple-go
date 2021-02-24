package model

// type User struct {
// 	BaseModel
// 	Username string    `gorm:"size:50;default:'';comment:'登录名';not null;index" form:"username" binding:"required"`
// 	Password string    `gorm:"type:varchar(100);default:'';comment:'密码';not null" form:"pass"`
// 	Status   int       `gorm:"type:tinyint;default:1"`
// 	CreateAt time.Time `gorm:"column:ctime;comment:'创建时间';not null" json:'-'`
// }

// type Post struct {
// 	PostId  int    `gorm:"primary_key"`
// 	Title   string `gorm:"type:varchar(255);index;not null;default:'';comment:'标题'"`
// 	Content string `gorm:"type:"text";not null;default('')`
// 	UserId  int    `gorm:"column:user_id;type:int;size:11;default(0)"`
// }

type Course struct {
	Index   int    `json:"index"`
	Name    string `json:"name"`
	Teacher string `json:"teacher"`
	Address string `json:"address"`
	Week    string `json:"week"`
	Time    string `json:"time"`
}

/*
avatarUrl: "https://thirdwx.qlogo.cn/mmopen/vi_32/0czGNpAAUcypaChmXwc1yIEUG8LSFoRHpcxk1aPkmtd3atibngiaT6XQibwwKstP9OWYVVphJjAlZbyquhaMhaBzw/132"
city: "Yangpu"
country: "China"
gender: 1
language: "zh_CN"
nickName: "m"
province: "Shanghai"
*/

type User struct {
	BaseModel
	WxOpenID string `json:"-"`
	Phone    string `json:"-"`
	Avatar   string `json:"avatar"`
	Gender   string `json:"gender"`
	Language string `json:"-"`
	NickName string `json:"nick_name"`
	Country  string `json:"-"`
	Province string `json:"-"`
	City     string `json:"-"`
}

type AuthCode struct {
	BaseModel
	Phone string
	Code  string
}

type Weico struct {
	BaseModel
	Content    string     `json:"content"`
	UserID     int        `gorm:"index" json:"-"`
	Pics       []WeicoPic `json:"pics"`
	User       User       `gorm:"-" json:"user"`
	LikeNum    int        `json:"like_num"`
	CommentNum int        `json:"comment_num"`
	TS         int        `json:"ts"`
	CateID     int        `gorm:"index" json:"cate_id"`
	IsLike     int        `gorm:"-" json:"is_like"`
}

type WeicoPic struct {
	BaseModel
	WeicoID int    `gorm:"index" json:"-"`
	Url     string `json:"url" gorm:"type:varchar(511)"`
}

type WeicoLike struct {
	BaseModel
	WeicoID int
	UserID  int
}

type WeicoComment struct {
	BaseModel
	WeicoID    int    `json:"-"`
	Content    string `json:"content"`
	FromUserID int    `json:"-"`
	ToUserID   int    `json:"-"`
	FromUser   User   `gorm:"-" json:"from_user"`
	ToUser     User   `gorm:"-" json:"to_user"`
	TS         int    `json:"ts"`
}

type WeicoCate struct {
	BaseModel
	Content string `json:"content"`
}
