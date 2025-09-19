package tag

import (
	"fmt"

	"github.com/PsjNick/go_nebula/interface_n"
	"github.com/PsjNick/go_nebula/model"
	"github.com/PsjNick/go_nebula/nebula"
	"github.com/PsjNick/go_nebula/schema"
)

// todo 暂时简单创建
// todo 后期对比数据对象，正对性修改属性
func CreateTag[T interface_n.BaseModeN](tag T) error {
	if nebula.NebulaSessionPool == nil {
		return fmt.Errorf("the Nebula session pool is not initialized")
	}

	tagName := model.GenName(tag)
	comment := tag.Comment()
	tagSchema := schema.GenNebulaSchema(tag)

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
