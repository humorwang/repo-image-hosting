package services

import (
	"repo-image-hosting/dto/request"
	"repo-image-hosting/dto/response"
)

type RepoInterface interface {
	GetFiles() []map[string]interface{}
	Push(filename, content string) (string, string, string)
	PushHasPath(filename, filepath, content string) (string, string, string)
	PushByBase64(imageDto request.ImageDto) response.ImageDto
	Del(filepath, sha string) string
}
