package api

import (
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"google.golang.org/api/youtube/v3"
)

func FetchMostPopularVideos() echo.HandlerFunc {
	return func(c echo.Context) error {

		yts := c.Get("yts").(*youtube.Service)

		// API を実行する際の条件を設定
		call := yts.Videos.List("id,snippet").
			Chart("mostPopular").MaxResults(5)

		// NuxtからpageTokenを受け取る
		pageToken := c.QueryParam("pageToken")

		// pageTokenがあればYouTube API 呼び出し時に PageToken() を使う
		if len(pageToken) > 0 {
			call = call.PageToken(pageToken)
		}

		res, err := call.Do() // YouTube API を実行
		if err != nil {
			logrus.Fatalf("Error calling Youtube API: %v", err)
		}
		return c.JSON(fasthttp.StatusOK, res)
	}
}
