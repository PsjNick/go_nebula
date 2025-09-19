package main

import (
	"github.com/PsjNick/go_nebula/model"
)

type EdgeTx struct {
	model.BaseEdgeModel
}

func (e EdgeTx) Name() string {
	return "MyEdgsdasdasdasde"
}

func main() {

	edge := EdgeTx{}

	name := model.GenName(&edge)

	println(name)

	//err := nebula.InitNebula(
	//	config.NebulaConfig{
	//		Username:  "root",
	//		Password:  "nebula",
	//		SpaceName: "TEST",
	//		Pool: config.PoolConfig{
	//			IdleTime: 2880,
	//			MaxSize:  3,
	//			MinSize:  1,
	//			Timeout:  600,
	//		},
	//		Hosts: []config.NebulaHost{
	//			config.NebulaHost{
	//				Host: "192.168.8.115",
	//				Port: 9669,
	//			},
	//		},
	//	},
	//)
	//
	//if err != nil {
	//	panic(err)
	//}
	//
	//type MyTag struct {
	//	model.BaseTagModel
	//	TxId string `n_name:"tx_id" n_type:"string" n_allow_null:"false" n_default:"" n_comment:"交易哈希"`
	//}
	//
	//err = tag.CreateTag(&MyTag{})
	//
	//if err != nil {
	//	panic(err)
	//}
	//
	//type EdgeBaseAmount struct {
	//	AmountInteger       string `json:"amountInteger" n_name:"amount_integer" n_type:"string" n_allow_null:"true" n_default:"\"0\"" n_comment:"amount 整数部分"`
	//	AmountIntegerPart_1 int64  `json:"amountIntegerPart_1" n_name:"amount_integer_part_1" n_type:"int64" n_allow_null:"true" n_default:"0" n_comment:"amount 整数部分 0~64位"`
	//	AmountIntegerPart_2 int64  `json:"amountIntegerPart_2" n_name:"amount_integer_part_2" n_type:"int64" n_allow_null:"true" n_default:"0" n_comment:"amount 整数部分 64~128位"`
	//	AmountIntegerPart_3 int64  `json:"amountIntegerPart_3" n_name:"amount_integer_part_3" n_type:"int64" n_allow_null:"true" n_default:"0" n_comment:"amount 整数部分 128~192位"`
	//	AmountIntegerPart_4 int64  `json:"amountIntegerPart_4" n_name:"amount_integer_part_4" n_type:"int64" n_allow_null:"true" n_default:"0" n_comment:"amount 整数部分 192~256位"`
	//
	//	AmountDecimal       string `json:"amountDecimal" n_name:"amount_decimal" n_type:"string" n_allow_null:"true" n_default:"\"0\"" n_comment:"amount 小数部分"`
	//	AmountDecimalPart_1 int64  `json:"amountDecimalPart_1" n_name:"amount_decimal_part_1" n_type:"int64" n_allow_null:"true" n_default:"0" n_comment:"amount 小数部分 0~64位"`
	//	AmountDecimalPart_2 int64  `json:"amountDecimalPart_2" n_name:"amount_decimal_part_2" n_type:"int64" n_allow_null:"true" n_default:"0" n_comment:"amount 小数部分 64~128位"`
	//	AmountDecimalPart_3 int64  `json:"amountDecimalPart_3" n_name:"amount_decimal_part_3" n_type:"int64" n_allow_null:"true" n_default:"0" n_comment:"amount 小数部分 128~192位"`
	//	AmountDecimalPart_4 int64  `json:"amountDecimalPart_4" n_name:"amount_decimal_part_4" n_type:"int64" n_allow_null:"true" n_default:"0" n_comment:"amount 小数部分 192~256位"`
	//
	//	AmtUsd float64 `json:"amtUsd" n_name:"amt_usd" n_type:"double" n_allow_null:"true" n_default:"0.0" n_comment:"美元价值"`
	//	AmtCny float64 `json:"amtCny" n_name:"amt_cny" n_type:"double" n_allow_null:"true" n_default:"0.0" n_comment:"人民币价值"`
	//}
	//
	//type MyEdge struct {
	//	model.BaseEdgeModel
	//	EdgeBaseAmount
	//	TxId string `n_name:"tx_id" n_type:"string" n_allow_null:"false" n_default:"" n_comment:"交易哈希"`
	//}
	//
	//err = edge.CreateEdge(&MyEdge{})
	//if err != nil {
	//	panic(err)
	//}
}
