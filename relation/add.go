package relation

import (
	"fmt"
	"reflect"

	"github.com/PsjNick/go_nebula/interface_n"
	"github.com/PsjNick/go_nebula/nebula"
)

func AddEdge[T interface_n.BaseModeN](edge T, vertexIdFrom string, vertexIdTo string) (err error) {

	if nebula.NebulaSessionPool == nil {
		return fmt.Errorf("the Nebula session pool is not initialized")
	}

	var edgeName string

	var fields []string
	var values []interface{}

	edgeName = edge.GetName()

	t := reflect.TypeOf(edge)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if nNameTag := field.Tag.Get("n_name"); nNameTag != "" {
			fields = append(fields, nNameTag)
			values = append(values, reflect.ValueOf(edge).Field(i).Interface())
		}
	}

	fieldsStr := ""
	valuesStr := ""

	// 辅助函数：将字段切片转换为逗号分隔的字符串
	joinFields := func(fields []string) string {
		for i, field := range fields {
			if i > 0 {
				fieldsStr += ", "
			}
			fieldsStr += field
		}
		return fieldsStr
	}
	fieldsStr = joinFields(fields)

	// 辅助函数：将值切片转换为逗号分隔的字符串
	joinValues := func(values []interface{}) string {
		for i, value := range values {
			if i > 0 {
				valuesStr += ", "
			}
			switch v := value.(type) {
			case string:
				valuesStr += fmt.Sprintf("\"%s\"", v) // 字符串类型需要加引号
			case int, int8, int16, int32, int64, uint8, uint16, uint32, uint64, float32, float64:
				valuesStr += fmt.Sprintf("%v", v) // 数值类型直接转换
			case bool:
				if v {
					valuesStr += "true"
				} else {
					valuesStr += "false"
				}
			default:
				valuesStr += fmt.Sprintf("\"%v\"", v) // 其他类型统一处理为字符串
			}
		}
		return valuesStr
	}
	valuesStr = joinValues(values)

	// 构建插入语句
	insertStmt := fmt.Sprintf("INSERT EDGE %s (%s) VALUES \"%s\"->\"%s\":(%s);", edgeName, fieldsStr, vertexIdFrom, vertexIdTo, valuesStr)

	// 执行插入操作
	resp, err := nebula.NebulaSessionPool.Execute(insertStmt)
	if err != nil {
		return fmt.Errorf("failed to insert vertex: %w", err)
	}
	if !resp.IsSucceed() {
		return fmt.Errorf("INSERT VERTEX failed: %s", resp.GetErrorMsg())
	}

	return nil
}
