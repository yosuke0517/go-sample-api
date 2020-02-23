package api

import (
	"context"
    "github.com/labstack/echo"
    "github.com/valyala/fasthttp"
	"github.com/sirupsen/logrus"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
	"os"
)

func FetchMostPopularVideos() echo.HandlerFunc { return func(c echo.Context) error {
	key := os.Getenv("API_KEY") //APIKey

	ctx := context.Background() // コンテキスト
	yts, err := youtube.NewService(ctx, option.WithAPIKey(key)) // サービスを生成
	if err != nil {
		logrus.Fatalf("Error creating new Youtube service: %v", err)
	}

	call := yts.Videos.List("id,snippet").Chart("mostPopular").MaxResults(3) // API を実行する際の条件を設定

	res, err := call.Do() // YouTube API を実行
	if err != nil {
		logrus.Fatalf("Error calling Youtube API: %v", err)
	}
	return c.JSON(fasthttp.StatusOK, res)
  }
}
