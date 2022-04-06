package utils

type UserInfo struct {
	Id       int64
	UserName string `json:UserName`
	PassWord string `json:Password`
}

type UserSetting struct {
	Id     int64
	ImgUrl string `json:"ImgUrl""`
}

type UserFavourite struct {
	Id       int64  `json:"FavouriteId"`
	UserId   int64  `json:"UserId"`
	Alias    string `json:"Alias"`
	Url      string `json:"Url"`
	IconUrl  string `json:"IconUrl"`
	Priority int    `json:"Priority"`
}
