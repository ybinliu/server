package response

import "server/model/system"

type SysMenuResponse struct {
	Menus []system.SysMenu `json:"menus"`
}

type SysBaseMenusResponse struct {
	Menus []system.SysBaseMenu `json:"menus"`
}

type SysBaseMenuResponse struct {
	Menu system.SysBaseMenu `json:"menu"`
}
