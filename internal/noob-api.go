package internal

import (
	"github.com/busy-cloud/boat/api"
	"github.com/busy-cloud/boat/curd"
	"github.com/busy-cloud/boat/db"
	"github.com/gin-gonic/gin"
)

func init() {
	api.Register("GET", "noob/noob/list", curd.ApiList[Noob]())
	api.Register("POST", "noob/noob/search", curd.ApiSearch[Noob]())
	api.Register("POST", "noob/noob/create", curd.ApiCreate[Noob]())
	api.Register("GET", "noob/noob/:id", curd.ApiGet[Noob]())
	api.Register("POST", "noob/noob/:id", curd.ApiUpdate[Noob]("id", "name", "disabled"))
	api.Register("GET", "noob/noob/:id/delete", curd.ApiDelete[Noob]())
	api.Register("GET", "noob/noob/:id/enable", curd.ApiDisable[Noob](true))
	api.Register("GET", "noob/noob/:id/disable", curd.ApiDisable[Noob](false))
	api.Register("GET", "noob/noob/:id/sync", func(ctx *gin.Context) {
		id := ctx.Param("id")

		var devices []*Device = make([]*Device, 0)
		err := db.Engine().Where("noob_id=?", id).Find(&devices)
		if err != nil {
			api.Error(ctx, err)
			return
		}

		//下发指令
		cmd, err := Request(id, Command{
			"cmd":  "config_write",
			"name": "devices",
			"data": devices,
		})
		if err != nil {
			api.Error(ctx, err)
			return
		}
		if cmd["ret"] != 1 {
			api.Fail(ctx, cmd["msg"].(string))
			return
		}

		//检查连接
		var links []*NoobLink = make([]*NoobLink, 0)
		err = db.Engine().Where("noob_id=?", id).Find(&links)
		if err != nil {
			api.Error(ctx, err)
			return
		}

		//下发指令
		cmd, err = Request(id, Command{
			"cmd":  "config_write",
			"name": "links",
			"data": links,
		})
		if err != nil {
			api.Error(ctx, err)
			return
		}
		if cmd["ret"] != 1 {
			api.Fail(ctx, cmd["msg"].(string))
			return
		}

		api.OK(ctx, nil)
	})

}
