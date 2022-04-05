package utils

import (
	"fmt"
	"strings"
	"time"
)

func GetUserBg(id int64) (string, error) {
	template := `Select BackGroundUrl From startpage.Setting Where UserId =? limit 1`
	result, err := DB().Query(template, id)
	if err != nil {
		return "", err
	}
	defer result.Close()
	result.Next()
	url := ""
	err = result.Scan(&url)
	if err != nil {
		return "", err
	}
	return url, nil
}

func GetFavouriteById(id int64) ([]UserFavourite, error) {
	template := `Select Alias,Url,Priority From startpage.Favourite Where UserId = ?`
	result, err := DB().Query(template, id)
	if err != nil {
		return nil, err
	}
	defer result.Close()
	var favourites []UserFavourite
	var favourite UserFavourite
	for result.Next() {
		result.Scan(&favourite.Alias, &favourite.Url, &favourite.Priority)
		favourite.IconUrl = strings.Split(favourite.Url, "/")[0]
		favourites = append(favourites, favourite)
	}
	return favourites, nil
}
func InsertFavourite(favourite UserFavourite) error {
	template := `Insert Into startpage.Favourite Set UserId=?,Url=?,Alias=?,InsertedDate=?,Priority=?`
	result, err := DB().Exec(template, favourite.Id, favourite.Url, favourite.Alias, time.Now(), favourite.Priority)
	if err != nil {
		fmt.Println(err)
		return err
	}
	_, err = result.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
