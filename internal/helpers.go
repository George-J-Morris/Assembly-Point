package internal

import (
	"encoding/base64"
	"encoding/json"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func Render(ctx echo.Context, statusCode int, t templ.Component) error {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)

	if err := t.Render(ctx.Request().Context(), buf); err != nil {
		return err
	}

	return ctx.HTML(statusCode, buf.String())
}

func GetJSONRawBody(c echo.Context) map[string]interface{} {

	jsonBody := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&jsonBody)
	if err != nil {

		log.Error("empty json body")
		return nil
	}

	return jsonBody
}

func Base64StringToBytes(base64Str string) ([]byte, error) {

	base64Bytes := []byte(base64Str)

	dst := make([]byte, base64.StdEncoding.DecodedLen(len(base64Bytes)))

	n, err := base64.StdEncoding.Decode(dst, []byte(base64Bytes))

	if err != nil {
		return nil, err
	}
	dst = dst[:n]

	return dst, nil
}
