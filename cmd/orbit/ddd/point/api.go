package point

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/go-log"
)

type Req struct {
	Lat float64 `json:"latitude,omitempty"`
	Lon float64 `json:"longitude,omitempty"`
}

func point(ctx iris.Context) {
	log.Debug("receive point")
	var req Req
	var resp app.Response
	var err error
	err = ctx.ReadJSON(&req)
	if err != nil {
		log.Debug(err)
		resp.Code = tool.RespCodeError
		resp.Message = err.Error()
		tool.ResponseJSON(ctx, resp)
		return
	}

	log.Info("%f:%f", req.Lon, req.Lat)

	resp.Code = tool.RespCodeSuccess
	tool.ResponseJSON(ctx, resp)
}
