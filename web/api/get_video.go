package api

import (
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"google.golang.org/api/youtube/v3"
)

type VideoResponse struct {
	VideoList *youtube.VideoListResponse `json:"video_list"`
}

func GetVideo() echo.HandlerFunc {
	return func(c echo.Context) error {
		// youtube.Service はコンテキストから取得
		yts := c.Get("yts").(*youtube.Service)

		// 動画の ID を取得
		videoID := c.Param("id")

		// YouTube API リクエスト時に Id() の引数として使用
		call := yts.Videos.List("id,snippet").Id(videoID)
		res, err := call.Do()

		if err != nil {
			logrus.Fatalf("Error calling Youtube API: %v", err)
		}

		// 動画をお気に入り追加済みかどうか判定するためのフラグをレスポンスに追加することになるため構造体で返却する
		v := VideoResponse{
			VideoList: res,
		}
		return c.JSON(fasthttp.StatusOK, v)
	}
}
