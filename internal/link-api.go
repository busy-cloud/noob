package internal

import (
	"github.com/busy-cloud/boat/api"
	"github.com/busy-cloud/boat/curd"
)

func init() {
	api.Register("GET", "noob/link/list", curd.ApiList[NoobLink]())
	api.Register("POST", "noob/link/search", curd.ApiSearch[NoobLink]())
	api.Register("POST", "noob/link/create", curd.ApiCreate[NoobLink]())
	api.Register("GET", "noob/link/:id", curd.ApiGet[NoobLink]())
	api.Register("POST", "noob/link/:id", curd.ApiUpdate[NoobLink]("id", "noob_id", "name", "disabled", "type", "address", "port", "baud_rate", "data_bits", "stop_bits", "parity", "protocol", "protocol_options"))
	api.Register("GET", "noob/link/:id/delete", curd.ApiDelete[NoobLink]())
	api.Register("GET", "noob/link/:id/enable", curd.ApiDisable[NoobLink](true))
	api.Register("GET", "noob/link/:id/disable", curd.ApiDisable[NoobLink](false))
}
