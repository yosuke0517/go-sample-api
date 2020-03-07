package routes

import (
	"app/middlewares"
	"app/web/api"
	"github.com/labstack/echo"
)

func Init(e *echo.Echo) {
	g := e.Group("/api")
	{
		g.GET("/popular", api.FetchMostPopularVideos())
		g.GET("/video/:id", api.GetVideo())
		g.GET("/related/:id", api.FetchRelatedVideos())
		g.GET("/search", api.SearchVideos())
	}
	//第 2 引数に middlewares.FirebaseGuard() を設定することでAuthミドルウェアが適用される
	fg := g.Group("/favorite", middlewares.FirebaseGuard())
	{
		fg.POST("/:id/toggle", api.ToggleFavoriteVideo())
		fg.GET("", api.FetchFavoriteVideos())
	}

}
