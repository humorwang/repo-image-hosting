package connector

import (
	"repo-image-hosting/config"
	"repo-image-hosting/services"
	"repo-image-hosting/services/gitee"
	"repo-image-hosting/services/github"
)

// 定义 serve 的映射关系
var serveMap = map[string]services.RepoInterface{
	"gitee":  &gitee.GiteeServe{},
	"github": &github.GithubServe{},
}

func RepoCreate() services.RepoInterface {
	return serveMap[config.Conf.App.Platform]
}
