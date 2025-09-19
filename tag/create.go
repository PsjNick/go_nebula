package tag

import (
	"fmt"

	"github.com/PsjNick/go_nebula/nebula"
)

// CsreateTagIfNotExists 检查并创建 Nebula 标签
func CreateTagIfNotExists(tagName, tagSchema, comment string) error {
	if nebula.NebulaSessionPool == nil {
		return fmt.Errorf("the Nebula session pool is not initialized")
	}

	// 直接尝试创建标签，如果已存在会自动跳过
	createTagStmt := fmt.Sprintf("CREATE TAG IF NOT EXISTS `%s` (%s) COMMENT = \"%s\"", tagName, tagSchema, comment)
	resp, err := nebula.NebulaSessionPool.Execute(createTagStmt)
	if err != nil {
		return fmt.Errorf("failed to create tag: %w", err)
	}
	if !resp.IsSucceed() {
		return fmt.Errorf("CREATE TAG failed: %s", resp.GetErrorMsg())
	}

	return nil
}
