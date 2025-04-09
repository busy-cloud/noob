package internal

import (
	"encoding/json"
	"github.com/busy-cloud/boat/db"
	"github.com/busy-cloud/boat/log"
	"github.com/busy-cloud/boat/mqtt"
	"github.com/god-jason/iot-master/product"
	"strings"
)

type Register struct {
	Bsp           string   `json:"bsp"`
	Firmware      string   `json:"firmware"`
	Imei          string   `json:"imei"`
	Iccid         string   `json:"iccid"`
	Imsi          string   `json:"imsi"`
	Devices       int      `json:"devices"`
	Links         int      `json:"links"`
	WantedConfigs []string `json:"wanted_configs"`
}

func subscribe() {

	//订阅数据变化
	mqtt.SubscribeStruct[Register]("noob/+/register", func(topic string, register *Register) {
		ss := strings.Split(topic, "/")
		id := ss[1]

		var noob Noob
		has, err := db.Engine().ID(id).Get(&noob)
		if err != nil {
			log.Error(err)
			return
		}
		if !has {
			noob.Id = id
			noob.Bsp = register.Bsp
			noob.Firmware = register.Firmware
			noob.Imei = register.Imei
			noob.Imsi = register.Imsi
			noob.Iccid = register.Iccid
			_, _ = db.Engine().InsertOne(&noob)
			return
		} else {
			var noob Noob
			noob.Bsp = register.Bsp
			noob.Firmware = register.Firmware
			noob.Imei = register.Imei
			noob.Imsi = register.Imsi
			noob.Iccid = register.Iccid
			_, _ = db.Engine().ID(id).Cols("bsp", "firmware", "imei", "iccid", "imsi").Update(&noob)
		}

		//检查设备列表
		cnt, err := db.Engine().Where("noob_id=?", noob.Id).Count(new(Device))
		if err != nil {
			log.Error(err)
			return
		}
		if int(cnt) != register.Devices {
			var devices []*Device = make([]*Device, 0)
			err = db.Engine().Where("noob_id=?", noob.Id).Find(&devices)
			if err != nil {
				log.Error(err)
				return
			}

			//下发指令
			//mqtt.Publish("noob/"+id+"/command", map[string]any{"cmd": "config_write", "name": "devices", "data": devices})
			Execute(id, Command{
				"cmd":  "config_write",
				"name": "devices",
				"data": devices,
			})
		}

		//检查连接
		cnt, err = db.Engine().Where("noob_id=?", noob.Id).Count(new(NoobLink))
		if err != nil {
			log.Error(err)
			return
		}
		if int(cnt) != register.Links {
			var links []*NoobLink = make([]*NoobLink, 0)
			err = db.Engine().Where("noob_id=?", noob.Id).Find(&links)
			if err != nil {
				log.Error(err)
				return
			}

			//下发指令
			//mqtt.Publish("noob/"+id+"/command", map[string]any{"cmd": "config_write", "name": "links", "data": links})
			Execute(id, Command{
				"cmd":  "config_write",
				"name": "links",
				"data": links,
			})
		}

		//检查项目配置
		if register.WantedConfigs != nil {
			for _, config := range register.WantedConfigs {
				ss := strings.Split(config, "/")
				if len(ss) != 2 {
					continue
				}

				var pc product.ProductConfig
				has, err := db.Engine().Where("id=?", ss[0]).And("name=?", ss[1]).Get(&pc)
				if err != nil {
					log.Error(err)
					continue
				}
				if has {
					//下发指令
					//mqtt.Publish("noob/"+id+"/command", map[string]any{"cmd": "config_write", "name": config, "data": pc.Content})
					Execute(id, Command{
						"cmd":  "config_write",
						"name": config,
						"data": pc.Content,
					})
				}
			}
		}
	})

	//订阅数据变化，服务类型
	mqtt.Subscribe("noob/+/status", func(topic string, payload []byte) {
		//ss := strings.Split(topic, "/")
		log.Info(topic, string(payload))

		//TODO 存储状态
	})

	//订阅数据变化，服务类型
	mqtt.Subscribe("noob/+/command/response", func(topic string, payload []byte) {
		//ss := strings.Split(topic, "/")
		//id := ss[1]
		var cmd Command
		err := json.Unmarshal(payload, &cmd)
		if err != nil {
			log.Error(err)
			return
		}

		if v, ok := cmd["id"]; ok {
			mid := v.(string)
			HandleResponse(mid, cmd)
		}
	})
}
