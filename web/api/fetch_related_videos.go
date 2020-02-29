package api

import (
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"google.golang.org/api/youtube/v3"
)

func FetchRelatedVideos() echo.HandlerFunc {
	return func(c echo.Context) error {
		yts := c.Get("yts").(*youtube.Service)

		videoId := c.Param("id")

		call := yts.Search.  // 関連動画の取得
			List("id,snippet").
			RelatedToVideoId(videoId).  // RelatedToVideoIdに動画のIDを渡す
			Type("video").  // ↑を使う際はType("video")とする
			MaxResults(3)

		res, err := call.Do()
		if err != nil {
			logrus.Fatalf("Error calling Youtube API: %v", err)
		}

		return c.JSON(fasthttp.StatusOK, res)
	}
}
