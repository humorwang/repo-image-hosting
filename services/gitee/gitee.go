package gitee

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"log"
	"repo-image-hosting/config"
	"repo-image-hosting/dto/request"
	"repo-image-hosting/dto/response"
	"repo-image-hosting/global"
	"repo-image-hosting/services"
	"strings"
	"time"
)

type GiteeServe struct {
	services.RepoInterface
}

func (serve *GiteeServe) Push(filename, content string) (string, string, string) {
	return PushGitee(filename, config.Conf.App.Path, content)
}

func (serve *GiteeServe) GetFiles() []map[string]interface{} {
	return GetGiteeFiles()
}

func (serve *GiteeServe) Del(filepath, sha string) string {
	return DelFile(filepath, sha)
}

func (serve *GiteeServe) PushHasPath(filename, filepath, content string) (string, string, string) {
	return PushGitee(filename, filepath, content)
}

func (serve *GiteeServe) PushByBase64(imagesDto request.ImageDto) response.ImageDto {
	imageData := response.ImageDto{}
	imageType := strings.TrimLeft(strings.Split(imagesDto.Image, ";")[0], "data:")
	imageExt := global.ImageExt[imageType]
	timeStr := time.Now().Format("20060102150405")
	filename := timeStr + "_" + services.GetRandomString(16) + imageExt
	imageContent := strings.Split(imagesDto.Image, ",")[1]
	url, path, sha := PushGitee(filename, imagesDto.Path, imageContent)
	imageData.Url = url
	imageData.Path = path
	imageData.Sha = sha
	return imageData
}

func PushGitee(filename, filepath, content string) (string, string, string) {

	url := "https://gitee.com/api/v5/repos/" + config.Conf.App.Owner + "/" + config.Conf.App.Repo + "/contents/" + filepath + "/" + filename

	// 初始化请求与响应
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer func() {
		// 用完需要释放资源
		fasthttp.ReleaseResponse(resp)
		fasthttp.ReleaseRequest(req)
	}()

	// 设置请求方法
	req.Header.SetMethod("POST")
	req.Header.SetBytesKV([]byte("Content-Type"), []byte("application/json"))

	// 设置请求的目标网址
	req.SetRequestURI(url)

	args := make(map[string]string)
	args["access_token"] = config.Conf.App.Token
	args["content"] = content
	args["message"] = "upload pic for repo-image-hosting"
	args["branch"] = config.Conf.App.Branch

	jsonBytes, _ := json.Marshal(args)
	req.SetBodyRaw(jsonBytes)

	// 发起请求
	if err := fasthttp.Do(req, resp); err != nil {
		log.Println(" 请求失败:", url, err.Error())
	}

	// 获取响应的数据实体
	body := resp.Body()

	var mapResult map[string]interface{}

	err := json.Unmarshal(body, &mapResult)
	if err != nil {
		fmt.Println("JsonToMapDemo err: ", err)
	}

	d := ""
	p := ""
	s := ""

	_, ok := mapResult["content"]

	if ok {
		if mapResult["content"] != nil {
			d = mapResult["content"].(map[string]interface{})["download_url"].(string)
			p = mapResult["content"].(map[string]interface{})["path"].(string)
			s = mapResult["content"].(map[string]interface{})["sha"].(string)
		}
	}

	return d, p, s
}

func GetGiteeFiles() []map[string]interface{} {

	url := "https://gitee.com/api/v5/repos/" +
		config.Conf.App.Owner + "/" +
		config.Conf.App.Repo +
		"/contents/" +
		config.Conf.App.Path +
		"?access_token=" +
		config.Conf.App.Token + "&ref=" + config.Conf.App.Branch

	// 初始化请求与响应
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer func() {
		// 用完需要释放资源
		fasthttp.ReleaseResponse(resp)
		fasthttp.ReleaseRequest(req)
	}()

	// 设置请求方法
	req.Header.SetMethod("GET")
	// 设置请求的目标网址
	req.SetRequestURI(url)

	// 发起请求
	if err := fasthttp.Do(req, resp); err != nil {
		log.Println(" 请求失败:", url, err.Error())
	}

	// 获取响应的数据实体
	body := resp.Body()

	var mapResult []map[string]interface{}

	err := json.Unmarshal(body, &mapResult)
	if err != nil {
		fmt.Println("JsonToMapDemo err: ", err)
	}

	return mapResult
}

func DelFile(filepath, sha string) string {

	url := "https://gitee.com/api/v5/repos/" + config.Conf.App.Owner + "/" + config.Conf.App.Repo + "/contents/" + filepath

	// 初始化请求与响应
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()

	defer func() {
		// 用完需要释放资源
		fasthttp.ReleaseResponse(resp)
		fasthttp.ReleaseRequest(req)
	}()

	// 设置请求方法
	req.Header.SetMethod("DELETE")
	req.Header.SetBytesKV([]byte("Content-Type"), []byte("application/json"))

	// 设置请求的目标网址
	req.SetRequestURI(url)

	args := make(map[string]string)
	args["access_token"] = config.Conf.App.Token
	args["sha"] = sha
	args["message"] = "delete pic for repo-image-hosting"
	args["branch"] = config.Conf.App.Branch

	jsonBytes, _ := json.Marshal(args)
	req.SetBodyRaw(jsonBytes)

	// 发起请求
	if err := fasthttp.Do(req, resp); err != nil {
		log.Println(" 请求失败:", url, err.Error())
	}

	// 获取响应的数据实体
	body := resp.Body()

	return string(body)
}
