package treex

//用户菜单管理
type UserMenus struct {
	Id          int64          `form:"id,default=0" json:"id"`
	Name        string         `form:"name" json:"name"`         // 路由名称
	Path        string         `form:"path" json:"path"`         // 路由path
	Icon        string         `form:"icon" json:"icon"`         // 路由图标
	Sort        int64          `form:"sort,default=0" json:"sort"`         // 排序标记
	ParentId    int64          `form:"parent_id,default=0" json:"parent_id"`    // 父菜单ID
	Hidden      int64          `form:"hidden,default=1" json:"hidden"`     // 是否在列表隐藏
	DefaultMenu int64          `form:"default_menu,default=0" json:"default_menu"` // 附加属性
	Component   string         `form:"component" json:"component"`    // 对应前端文件路径
}

//无限级分类
//最多支持4级分类
type MenuList []UserMenus

type MenuItem struct {
	UserMenus
	ChildrenNodes   []MenuItem `json:"children"`
}

func (m *MenuList) ProcessToTree(pid int64, level int64) []MenuItem {
	var menuTree []MenuItem
	if level == 4 {
		return menuTree
	}

	list := m.findChildren(pid)
	if len(list) == 0 {
		return menuTree
	}

	for _, v := range list {
		child := m.ProcessToTree(v.Id, level+1)
		menuTree = append(menuTree, MenuItem{v, child})
	}

	return menuTree
}

func (m *MenuList) findChildren(pid int64) []UserMenus {
	child := []UserMenus{}

	for _, v := range *m {
		if v.ParentId == pid {
			child = append(child, v)
		}
	}
	return child
}