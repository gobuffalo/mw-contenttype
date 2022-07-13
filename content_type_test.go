package contenttype_test

import (
	"testing"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/render"
	"github.com/gobuffalo/httptest"
	contenttype "github.com/gobuffalo/mw-contenttype"
	"github.com/stretchr/testify/require"
)

func ctApp() *buffalo.App {
	h := func(c buffalo.Context) error {
		return c.Render(200, render.String(c.Request().Header.Get("Content-Type")))
	}
	a := buffalo.New(buffalo.Options{})
	a.GET("/set", contenttype.Set("application/json")(h))
	a.GET("/add", contenttype.Add("application/json")(h))
	return a
}

func Test_SetContentType(t *testing.T) {
	r := require.New(t)

	w := httptest.New(ctApp())
	res := w.HTML("/set").Get()
	r.Equal("application/json", res.Body.String())

	req := w.HTML("/set")
	req.Headers["Content-Type"] = "text/plain"

	res = req.Get()
	r.Equal("application/json", res.Body.String())
}

func Test_AddContentType(t *testing.T) {
	r := require.New(t)

	w := httptest.New(ctApp())
	res := w.HTML("/add").Get()
	r.Equal("application/json", res.Body.String())

	req := w.HTML("/add")
	req.Headers["Content-Type"] = "text/plain"

	res = req.Get()
	r.Equal("text/plain", res.Body.String())
}
