package api

import (
	"app/middlewares"
	"app/models"
	"context"
	"firebase.google.com/go/auth"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"google.golang.org/api/youtube/v3"
	"strings"
)

type VideoResponse struct {
	VideoList  *youtube.VideoListResponse `json:"video_list"`
	IsFavorite bool                       `json:"is_favorite"`
}

func GetVideo() echo.HandlerFunc {
	return func(c echo.Context) error {
		// youtube.Service はコンテキストから取得
		yts := c.Get("yts").(*youtube.Service)
		dbs := c.Get("dbs").(*middlewares.DatabaseClient)
		//token, errr := c.Get("auth").(*auth.Token)
		// cookieTmp, err := c.Cookie("jwt_token")
		headerAuth := c.Request().Header.Get("Authorization")
		token := strings.Replace(headerAuth, "Bearer ", "", 1)
		authClient0 := c.Get("firebase").(*auth.Client)
		authUser0, err := authClient0.VerifyIDToken(context.Background(), token)
		if err != nil {
			logrus.Debug("tokenError", err)
		}

		// 動画の ID を取得
		videoId := c.Param("id")
		isFavorite := false
		if authUser0 != nil {
			//ログインしていた場合、その動画がお気に入り登録されているかを判定
			favorite := models.Favorite{}
			isFavoriteNotFound := dbs.DB.Table("favorites").
				Joins("INNER JOIN users ON users.id = favorites.user_id").
				Where(models.User{UID: authUser0.UID}).
				Where(models.Favorite{VideoId: videoId}).
				First(&favorite).
				RecordNotFound()

			logrus.Debug("isFavoriteNotFound: ", isFavoriteNotFound)
			if !isFavoriteNotFound {
				isFavorite = true
			}
		}

		// YouTube API リクエスト時に Id() の引数として使用
		call := yts.Videos.List("id,snippet").Id(videoId)
		res, err := call.Do()

		if err != nil {
			logrus.Fatalf("Error calling Youtube API: %v", err)
		}

		// 動画をお気に入り追加済みかどうか判定するためのフラグをレスポンスに追加することになるため構造体で返却する
		v := VideoResponse{
			VideoList:  res,
			IsFavorite: isFavorite,
		}
		return c.JSON(fasthttp.StatusOK, v)
	}
}
