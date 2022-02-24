package connector

import (
	"Deskmate/services"
	"Deskmate/services/flag_handle"
	"Deskmate/services/gitee"
)

//定义serve的映射关系
var serveMap = map[string]services.RepoInterface{
	"gitee": &gitee.GiteeServe{},
}

func RepoCreate() services.RepoInterface {
	return serveMap[flag_handle.PLATFORM]
}
