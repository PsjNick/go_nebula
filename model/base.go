package model

type BaseEdgeModel struct {
	Name   string `json:"edgeName"`
	Common string `json:"edgeCommon"`
}

func (m BaseEdgeModel) GetName() string {
	if m.Name == "" {
		panic("edgeName is empty")
	}
	return m.Name
}

func (m BaseEdgeModel) GetComment() string {
	return m.Common
}

type BaseTagModel struct {
	Name   string `json:"tagName"`
	Common string `json:"tagCommon"`
}

func (m BaseTagModel) GetName() string {

	if m.Name == "" {
		panic("tag name is empty")
	}

	return m.Name
}
func (m BaseTagModel) GetComment() string {
	return m.Common
}
