package main

import (
	"github.com/PsjNick/go_nebula/config"
	"github.com/PsjNick/go_nebula/model"
	"github.com/PsjNick/go_nebula/nebula"
	"github.com/PsjNick/go_nebula/tag"
)

func main() {
	err := nebula.InitNebula(
		config.NebulaConfig{
			Username:  "root",
			Password:  "nebula",
			SpaceName: "TEST",
			Pool: config.PoolConfig{
				IdleTime: 2880,
				MaxSize:  3,
				MinSize:  1,
				Timeout:  600,
			},
			Hosts: []config.NebulaHost{
				config.NebulaHost{
					Host: "192.168.8.115",
					Port: 9669,
				},
			},
		},
	)

	if err != nil {
		panic(err)
	}

	type MyTag struct {
		model.BaseTagModel
		TxId string `json:"txId" n_name:"tx_id" n_type:"string" n_allow_null:"false" n_default:"" n_comment:"交易哈希"`
	}

	myTag := MyTag{
		BaseTagModel: model.BaseTagModel{
			Name:   "MyTag",
			Common: "MyTag 测试用",
		},
		TxId: "asdasdasdasdasdczxczxzxc",
	}

	err = tag.CreateTagIfNotExists(myTag)

	if err != nil {
		panic(err)
	}

}
