package internal

import (
	"github.com/busy-cloud/boat/api"
	"github.com/busy-cloud/boat/curd"
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

	})

}
