package schema

import (
	"fmt"
	"reflect"
	"strings"
)

func GenNebulaSchema(model any) string {
	t := reflect.TypeOf(model)

	var schList []string

	for i := 0; i < t.NumField(); i++ {

		field := t.Field(i)

		if field.Anonymous && field.Type.Kind() == reflect.Struct {
			embeddedType := field.Type

			for j := 0; j < embeddedType.NumField(); j++ {

				schema := ""

				field := embeddedType.Field(j)
				// 读取n_name标签
				if nNameTag := field.Tag.Get("n_name"); nNameTag != "" {
					schema += fmt.Sprintf(" `%s` ", nNameTag)
				}

				// 读取n_type标签
				if nTypeTag := field.Tag.Get("n_type"); nTypeTag != "" {
					schema += fmt.Sprintf(" %s ", nTypeTag)
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

	schema := strings.Join(schList, " , ")

	return schema
}
