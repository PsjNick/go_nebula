package schema

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/PsjNick/go_nebula/interface_n"
)

func doGenSchema(t reflect.Type) (schList []string) {

	for i := 0; i < t.NumField(); i++ {

		field := t.Field(i)

		if field.Anonymous && field.Type.Kind() == reflect.Struct {
			embeddedType := field.Type
			schList = append(schList, doGenSchema(embeddedType)...)
			continue
		}

		schema := ""

		// 读取n_name标签
		if nNameTag := field.Tag.Get("n_name"); nNameTag != "" {
			schema += fmt.Sprintf("`%s` ", nNameTag)
		}

		// 读取n_type标签
		if nTypeTag := field.Tag.Get("n_type"); nTypeTag != "" {
			schema += fmt.Sprintf("%s ", nTypeTag)
		}

		// 读取n_allow_null标签
		if nAllowNullTag := field.Tag.Get("n_allow_null"); nAllowNullTag != "" {
			if nAllowNullTag == "false" {
				schema += " NOT NULL "
			} else {
				schema += " NULL "
			}
		}

		// 读取n_default标签
		if nDefaultTag := field.Tag.Get("n_default"); nDefaultTag != "" {
			schema += fmt.Sprintf(" DEFAULT %s ", nDefaultTag)
		}

		// 读取n_comment标签
		if nCommentTag := field.Tag.Get("n_comment"); nCommentTag != "" {
			schema += fmt.Sprintf(" COMMENT \"%s\" ", nCommentTag)
		}

		if schema != "" {
			schList = append(schList, schema)
		}

	}

	return schList
}

func GenNebulaSchema[T interface_n.BaseModeN](model T) string {
	t := reflect.TypeOf(model)

	schList := doGenSchema(t)

	schema := strings.Join(schList, " , ")

	return schema
}
