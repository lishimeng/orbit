package point

import (
	"github.com/kataras/iris/v12"
)

func Route(p iris.Party) {
	p.Post("/", point)
}
