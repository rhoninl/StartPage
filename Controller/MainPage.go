package Controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type favourites struct {
	url  string
	name string
}
type userinfo struct {
	username string `json:"username"`
}

func GetMainPage(c *gin.Context) {
	html := `
 		<table style="margin-left: 20px">
		`
	var user userinfo
	c.Bind(&user)
	user.username = `root`
	templates := `Select LinkName,Url From Favourite,User Where Favourite.UserId = User.UserId And UserName = ?`
	result, err := DB().Query(templates, user.username)
	if err != nil {
		c.HTML(http.StatusOK, "index.html", nil)
		return
	}
	defer result.Close()
	var favourite favourites
	for result.Next() {
		result.Scan(&favourite)
		html += `
		 <tr>
           <td style="padding-top: 5px"><img  height="20px" src="` + favourite.url + `/favicon.ico" /></td>
           <td align="center"><a  href="` + favourite.url + `"  style="font-size: 18px;" >` + favourite.name + `+</a></td>
         </tr>
		`
	}
	html += `
	</table>
	`
	c.HTML()
	//c.HTML(http.StatusOK,"index.html",gin.H{"favourite":html})
}
