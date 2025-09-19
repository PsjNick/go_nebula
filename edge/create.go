package edge

import (
	"fmt"

	"github.com/PsjNick/go_nebula/interface_n"
	"github.com/PsjNick/go_nebula/nebula"
	"github.com/PsjNick/go_nebula/schema"
)

// CreateEdgeIfNotExists 检查并创建 Nebula 边类型
func CreateEdgeIfNotExists[T interface_n.BaseModeN](edge T) error {
	if nebula.NebulaSessionPool == nil {
		return fmt.Errorf("the Nebula session pool is not initialized")
	}

	edgeName := edge.GetName()
	edgeSchema := schema.GenNebulaSchema(edge)
	comment := edge.GetComment()

	// 直接尝试创建边类型，如果已存在会返回特定错误
	createEdgeStmt := fmt.Sprintf("CREATE EDGE IF NOT EXISTS `%s` (%s) COMMENT = \"%s\"", edgeName, edgeSchema, comment)
	resp, err := nebula.NebulaSessionPool.Execute(createEdgeStmt)
	if err != nil {
		return fmt.Errorf("failed to create edge: %w", err)
	}
	if !resp.IsSucceed() {
		return fmt.Errorf("CREATE EDGE failed: %s", resp.GetErrorMsg())
	}

	return nil
}
