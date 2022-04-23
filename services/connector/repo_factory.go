package connector

import (
	"Deskmate/services"
	"Deskmate/services/flag_handle"
	// "Deskmate/services/gitee"
	"Deskmate/services/github"
)

//定义serve的映射关系
var serveMap = map[string]services.RepoInterface{
	//"gitee": &gitee.GiteeServe{},
	"github": &github.GithubServe{},
}

func RepoCreate() services.RepoInterface {
	return serveMap[flag_handle.PLATFORM]
}
