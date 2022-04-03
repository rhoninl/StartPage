package utils

type UserInfo struct {
	Id       int64
	UserName string `json:UserName`
	PassWord string `json:Password`
}

type UserSetting struct {
	Id     int64
	ImgUrl string `json:ImgUrl`
}
