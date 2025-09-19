package edge

import (
	"fmt"

	"github.com/PsjNick/go_nebula/interface_n"
	"github.com/PsjNick/go_nebula/nebula"
	"github.com/PsjNick/go_nebula/schema"
)

// todo 暂时简单创建
// todo 后期对比属性，进行增改删除
func CreateEdge[T interface_n.BaseModeN](edge T) error {
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
