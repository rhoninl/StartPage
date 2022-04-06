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
	template := `Select FavouriteId,Alias,Url,Priority From startpage.Favourite Where UserId = ?`
	result, err := DB().Query(template, id)
	if err != nil {
		return nil, err
	}
	defer result.Close()
	var favourites []UserFavourite
	var favourite UserFavourite
	for result.Next() {
		result.Scan(&favourite.Id, &favourite.Alias, &favourite.Url, &favourite.Priority)
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

func DeleteFavourite(favourite UserFavourite) error {
	template := `Delete From startpage.Favourite Where UserId=? And FavouriteId = ?`
	result, err := DB().Exec(template, favourite.UserId, favourite.Id)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows != 1 {
		return fmt.Errorf("error %d", rows)
	}
	return nil
}

func QueryFavouriteById(id int64) (UserFavourite, error) {
	var userInfo UserFavourite
	template := `Select Url,Alias From startpage.Favourite Where FavouriteId = ?`
	rows, err := DB().Query(template, id)
	if err != nil {
		return userInfo, err
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&userInfo.Url, &userInfo.Alias)
	}
	userInfo.Id = id
	return userInfo, nil
}

func UpdateFavourite(favourite UserFavourite) error {
	template := `Update startpage.Favourite Set Alias=?,Url=? Where UserId = ? And FavouriteId = ?`
	rows, err := DB().Exec(template, favourite.Alias, favourite.Url, favourite.UserId, favourite.Id)
	if err != nil {
		return err
	}
	n, err := rows.RowsAffected()
	if err != nil {
		return err
	}
	if n != 1 {
		return fmt.Errorf("error")
	}
	return nil
}
