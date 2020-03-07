package middlewares

import (
	"context"
	"firebase.google.com/go/auth"
	"github.com/labstack/echo"
	"github.com/valyala/fasthttp"
	"strings"
)

/*
 *auth.Token はUID を持つ構造体。この UID をお気に入り追加・削除で使用。
 */
func verifyFirebaseIDToken(ctx echo.Context, auth *auth.Client) (*auth.Token, error) {

	// リクエストのヘッダーからトークンを取り出す
	headerAuth := ctx.Request().Header.Get("Authorization")
	token := strings.Replace(headerAuth, "Bearer ", "", 1)
	//auth.VerifyIDToken() に取り出したトークンを与えることで検証する
	jwtToken, err := auth.VerifyIDToken(context.Background(), token)

	return jwtToken, err
}

/*
 *ログインした状態でなければ利用できない API に対して使用
 */
func FirebaseGuard() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authClient := c.Get("firebase").(*auth.Client)
			jwtToken, err := verifyFirebaseIDToken(c, authClient)
			if err != nil {
				return c.JSON(fasthttp.StatusUnauthorized, "Not Authenticated")
			}
			//検証後、*auth.Token をコンテキストに保存
			c.Set("auth", jwtToken)
			if err := next(c); err != nil {
				return err
			}
			return nil
		}
	}
}

/*

ログインしていなくとも利用できる API で 使用する*/
func FirebaseAuth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authClient := c.Get("firebase").(*auth.Client)
			jwtToken, _ := verifyFirebaseIDToken(c, authClient)
			c.Set("auth", jwtToken)
			if err := next(c); err != nil {
				return err
			}

			return nil
		}
	}
}
