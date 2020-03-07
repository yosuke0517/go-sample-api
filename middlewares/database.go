package middlewares

import (
	"app/databases"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type DatabaseClient struct {
	DB *gorm.DB
}

func DatabaseService() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			session, _ := databases.Connect()
			d := DatabaseClient{DB: session}

			defer d.DB.Close()

			// 実行したSQLをログに出力するようにします
			d.DB.LogMode(true)

			//アクションから使用できるように接続を保持したインスタンスをコンテキストに登録
			c.Set("dbs", &d)

			if err := next(c); err != nil {
				return err
			}

			return nil
		}
	}
}
