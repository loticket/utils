package treex

//用户菜单管理
type Trees struct {
	Id          int64          `form:"id,default=0" json:"id"`
	ParentId    int64          `form:"parent_id,default=0" json:"parent_id"`    // 父菜单ID
	Name        string         `form:"name" json:"name"`         // 路由名称
	Content     interface{}    `form:"content" json:"content"`         // 路由名称
}

//无限级分类
//最多支持4级分类
type TreesList []Trees

type TreesItem struct {
	Trees
	ChildrenNodes   []TreesItem `json:"children"`
}

func (m *TreesList) ProcessToTree(pid int64, level int64) []TreesItem {
	var menuTree []TreesItem
	if level == 4 {
		return menuTree
	}

	list := m.findChildren(pid)
	if len(list) == 0 {
		return menuTree
	}

	for _, v := range list {
		child := m.ProcessToTree(v.Id, level+1)
		menuTree = append(menuTree, TreesItem{v, child})
	}

	return menuTree
}

func (m *TreesList) findChildren(pid int64) []Trees {
	child := []Trees{}

	for _, v := range *m {
		if v.ParentId == pid {
			child = append(child, v)
		}
	}
	return child
}