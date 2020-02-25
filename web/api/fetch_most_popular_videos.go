package api

import (
	"context"
	"os"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

func FetchMostPopularVideos() echo.HandlerFunc {
	return func(c echo.Context) error {

		key := os.Getenv("API_KEY") //APIKey

		ctx := context.Background()                                 // コンテキスト
		yts, err := youtube.NewService(ctx, option.WithAPIKey(key)) // サービスを生成

		// API を実行する際の条件を設定
		call := yts.Videos.List("id,snippet").
			Chart("mostPopular").MaxResults(3)

		// NuxtからpageTokenを受け取る
		pageToken := c.QueryParam("pageToken")

		// pageTokenがあればYouTube API 呼び出し時に PageToken() を使う
		if len(pageToken) > 0 {
			call = call.PageToken(pageToken)
		}
		if err != nil {
			logrus.Fatalf("Error creating new Youtube service: %v", err)
		}

		res, err := call.Do() // YouTube API を実行
		if err != nil {
			logrus.Fatalf("Error calling Youtube API: %v", err)
		}
		return c.JSON(fasthttp.StatusOK, res)
	}
}
