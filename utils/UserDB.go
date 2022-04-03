package utils

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
