package internal

import (
	"github.com/busy-cloud/boat/api"
	"github.com/busy-cloud/boat/curd"
)

func init() {
	api.Register("GET", "noob/device/list", curd.ApiList[Device]())
	api.Register("POST", "noob/device/search", curd.ApiSearch[Device]())
	api.Register("POST", "noob/device/create", curd.ApiCreate[Device]())
	api.Register("GET", "noob/device/:id", curd.ApiGet[Device]())
	api.Register("POST", "noob/device/:id", curd.ApiUpdate[Device]("id", "noob_id", "link_id", "name", "disabled", "product_id", "description", "station"))
	api.Register("GET", "noob/device/:id/delete", curd.ApiDelete[Device]())
	api.Register("GET", "noob/device/:id/enable", curd.ApiDisable[Device](true))
	api.Register("GET", "noob/device/:id/disable", curd.ApiDisable[Device](false))
}
