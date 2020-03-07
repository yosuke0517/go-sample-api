package middlewares

import (
	"context"
	firebase "firebase.google.com/go"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"google.golang.org/api/option"
	"os"
)

func Firebase() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// サービスアカウントjson ファイルのパス
			opt := option.WithCredentialsFile(os.Getenv("KEY_JSON_PATH"))
			//Firebase のプロジェクト ID
			config := &firebase.Config{ProjectID: os.Getenv("PROJECT_ID")}
			//認証情報とプロジェクト ID を使って、*firebase.Appを生成する
			app, err := firebase.NewApp(context.Background(), config, opt)
			if err != nil {
				logrus.Fatal("Error Initializing firebase: %v¥n", err)
			}
			auth, err := app.Auth(context.Background())

			//トークンのバリデーション時に使用するため、コンテキストにセットしておく
			c.Set("firebase", auth)
			if err := next(c); err != nil {
				return err
			}
			return nil
		}
	}
}
