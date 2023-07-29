package ddd

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/orbit/cmd/orbit/ddd/point"
)

func Route(app *iris.Application) {

	p := app.Party("/api")
	point.Route(p.Party("/point"))
}
