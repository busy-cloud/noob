package internal

import (
	"github.com/busy-cloud/boat/db"
	"github.com/god-jason/iot-master/device"
	"time"
)

func init() {
	db.Register(new(Noob), new(NoobLink), new(Device))
}

type Noob struct {
	Id       string    `json:"id,omitempty" xorm:"pk"`
	Name     string    `json:"name,omitempty"`
	Bsp      string    `json:"bsp,omitempty"`
	Firmware string    `json:"firmware,omitempty"`
	Imei     string    `json:"imei,omitempty"`
	Iccid    string    `json:"iccid,omitempty"`
	Imsi     string    `json:"imsi,omitempty"`
	Disabled bool      `json:"disabled,omitempty"`                        //禁用
	Created  time.Time `json:"created,omitempty,omitzero" xorm:"created"` //创建时间
}

type NoobLink struct {
	Id              string         `json:"id,omitempty" xorm:"pk"`
	NoobId          string         `json:"noob_id,omitempty" xorm:"index"`
	Name            string         `json:"name,omitempty"`
	Type            string         `json:"type,omitempty"`                            //serial tcp-client tcp-server tcp-server-multiple tcp-incoming
	Address         string         `json:"address,omitempty"`                         //地址，域名或IP
	Port            uint16         `json:"port,omitempty"`                            //端口号
	BaudRate        int            `json:"baud_rate,omitempty"`                       //9600 115200
	DataBits        int            `json:"data_bits,omitempty"`                       //5 6 7 8
	StopBits        int            `json:"stop_bits,omitempty"`                       //1 2
	Parity          string         `json:"parity,omitempty"`                          //N O E
	Protocol        string         `json:"protocol,omitempty"`                        //通讯协议
	ProtocolOptions map[string]any `json:"protocol_options,omitempty" xorm:"json"`    //通讯协议参数
	Disabled        bool           `json:"disabled,omitempty"`                        //禁用
	Created         time.Time      `json:"created,omitempty,omitzero" xorm:"created"` //创建时间
}

type Device struct {
	device.Device `xorm:"extends"`
	NoobId        string `json:"noob_id,omitempty" xorm:"index"`
	//LinkId        string `json:"link_id,omitempty" xorm:"index"`
}
