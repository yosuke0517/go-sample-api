package api

import (
	"app/middlewares"
	"app/models"
	"firebase.google.com/go/auth"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"google.golang.org/api/youtube/v3"
)

func FetchFavoriteVideos() echo.HandlerFunc {
	return func(c echo.Context) error {
		dbs := c.Get("dbs").(*middlewares.DatabaseClient)
		token := c.Get("auth").(*auth.Token)

		//User取得
		user := models.User{}
		//SELECT * FROM `users`  WHERE `users`.`deleted_at` IS NULL AND ((`users`.`uid` = 'HOGE')) ORDER BY `users`.`id` ASC LIMIT 1みたいなSQLになる
		dbs.DB.Table("users").Where(models.User{UID: token.UID}).First(&user)

		//お気に入り
		favorites := []models.Favorite{}
		//SELECT * FROM `favorites`  WHERE (`user_id` = 1)みたいなSQLになる
		dbs.DB.Model(&user).Related(&favorites)

		videoIds := ""
		for _, f := range favorites {
			if len(videoIds) == 0 {
				videoIds += f.VideoId
			} else {
				videoIds += "," + f.VideoId
			}
		}
		yts := c.Get("yts").(*youtube.Service)
		call := yts.Videos.
			List("id,snippet").
			Id(videoIds).
			MaxResults(10)

		res, err := call.Do()
		if err != nil {
			logrus.Fatalf("Error calling Youtube API: %v", err)
		}
		return c.JSON(fasthttp.StatusOK, res)
	}
}
