package api

import (
	"app/middlewares"
	"app/models"
	"firebase.google.com/go/auth"
	"github.com/labstack/echo"
	"github.com/valyala/fasthttp"
)

type ToggleFavoriteVideoResponse struct {
	VideoId    string `json:"video_id"`
	IsFavorite bool   `json:"is_favorite"`
}

func ToggleFavoriteVideo() echo.HandlerFunc {
	return func(c echo.Context) error {
		dbs := c.Get("dbs").(*middlewares.DatabaseClient)
		videoId := c.Param("id")
		token := c.Get("auth").(*auth.Token)
		user := models.User{}
		//users テーブルにトークンから取得した UID を保持するレコードが存在 するかを確認
		if dbs.DB.Table("users").
			Where(models.User{UID: token.UID}).First(&user).RecordNotFound() {
			//レコードが見つからなかった場合は、新規ユーザーとして users テーブ ルにレコードを作成
			user = models.User{UID: token.UID}
			dbs.DB.Create(&user)
		}

		favorite := models.Favorite{}
		isFavorite := false
		//お気に入りの追加
		//favorites テーブルに条件に合致するレコードが存 在しない場合は追加、逆は削除
		if dbs.DB.Table("favorites").
			Where(models.Favorite{UserId: user.ID, VideoId: videoId}).First(&favorite).RecordNotFound() {

			favorite = models.Favorite{UserId: user.ID, VideoId: videoId}
			dbs.DB.Create(&favorite)
			isFavorite = true
		} else {
			//削除
			dbs.DB.Delete(&favorite)
		}

		res := ToggleFavoriteVideoResponse{
			VideoId:    videoId,
			IsFavorite: isFavorite,
		}

		return c.JSON(fasthttp.StatusOK, res)
	}
}
