package model

import (
	"fmt"
	"reflect"

	"github.com/PsjNick/go_nebula/interface_n"
)

func GenName(t interface_n.BaseModeN) string {

	if t.Name() != "" {
		return t.Name()
	}

	fmt.Printf("%T", t)

	return reflect.TypeOf(t).Name()

}

type BaseEdgeModel struct {
}

func (BaseEdgeModel) Name() string {
	return ""
}

func (m BaseEdgeModel) Comment() string {
	return ""
}

type BaseTagModel struct {
}

func (BaseTagModel) Name() string {
	return ""
}

func (m BaseTagModel) Comment() string {
	return ""
}
