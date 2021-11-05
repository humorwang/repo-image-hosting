package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"repo-image-hosting/config"
	"repo-image-hosting/services/connector"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"current_dir": config.Conf.App.Owner + "/" + config.Conf.App.Repo + "/" + config.Conf.App.Path,
		"owner":       config.Conf.App.Owner,
		"repo":        config.Conf.App.Repo,
		"platform":    config.Conf.App.Platform,
	})
}

func Images(c *gin.Context) {
	images := connector.RepoCreate().GetFiles()
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"success": true,
		"data": map[string]interface{}{
			"images": images,
		},
	})
}
