package vertex

import (
	"fmt"
	"reflect"

	"github.com/PsjNick/go_nebula/interface_n"
	"github.com/PsjNick/go_nebula/nebula"
)

// INSERT VERTEX IF NOT EXISTS t2 (name, age) VALUES "1":("n3", 14);
func AddVertexIfNotExists[T interface_n.BaseModeN](id string, vertex T) (err error) {
	if nebula.NebulaSessionPool == nil {
		return fmt.Errorf("the Nebula session pool is not initialized")
	}

	var tagName string
	var vertexID string

	var fields []string
	var values []interface{}

	tagName = vertex.GetName()
	vertexID = id

	t := reflect.TypeOf(vertex)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if nNameTag := field.Tag.Get("n_name"); nNameTag != "" {
			fields = append(fields, nNameTag)
			values = append(values, reflect.ValueOf(vertex).Field(i).Interface())
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

	// 构建 INSERT 语句
	insertStmt := fmt.Sprintf("INSERT VERTEX IF NOT EXISTS `%s` (%s) VALUES \"%s\":(%s)", tagName, fieldsStr, vertexID, valuesStr)

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
