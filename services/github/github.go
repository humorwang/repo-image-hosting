package github

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

type GithubServe struct {
	services.RepoInterface
}

func (serve *GithubServe) Push(filename, content string) (string, string, string) {
	return Push(filename, config.Conf.App.Path, content)
}

func (serve *GithubServe) GetFiles() []map[string]interface{} {
	return GetFiles()
}

func (serve *GithubServe) Del(filepath, sha string) string {
	return DelFile(filepath, sha)
}

func (serve *GithubServe) PushHasPath(filename, filepath, content string) (string, string, string) {
	return Push(filename, filepath, content)
}

func (serve *GithubServe) PushByBase64(imagesDto request.ImageDto) response.ImageDto {
	imageData := response.ImageDto{}
	imageType := strings.TrimLeft(strings.Split(imagesDto.Image, ";")[0], "data:")
	imageExt := global.ImageExt[imageType]
	timeStr := time.Now().Format("20060102150405")
	filename := timeStr + "_" + services.GetRandomString(16) + imageExt
	imageContent := strings.Split(imagesDto.Image, ",")[1]
	url, path, sha := Push(filename, imagesDto.Path, imageContent)
	imageData.Url = url
	imageData.Path = path
	imageData.Sha = sha
	return imageData
}

func Push(filename, filepath, content string) (string, string, string) {

	url := "https://api.github.com/repos/" + config.Conf.App.Owner + "/" + config.Conf.App.Repo + "/contents/" + filepath + "/" + filename

	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer func() {
		// 用完需要释放资源
		fasthttp.ReleaseResponse(resp)
		fasthttp.ReleaseRequest(req)
	}()

	// 设置请求方法
	req.Header.SetMethod("PUT")
	req.Header.SetBytesKV([]byte("Content-Type"), []byte("application/json"))
	req.Header.SetBytesKV([]byte("Accept"), []byte("application/vnd.github.v3+json"))
	req.Header.SetBytesKV([]byte("Authorization"), []byte("token "+config.Conf.App.Token))

	// 设置请求的目标网址
	req.SetRequestURI(url)

	args := make(map[string]string)
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
			path := mapResult["content"].(map[string]interface{})["path"].(string)
			d = "https://cdn.jsdelivr.net/gh/" + config.Conf.App.Owner + "/" + config.Conf.App.Repo + "@" + config.Conf.App.Branch + "/" + path
			p = path
			s = mapResult["content"].(map[string]interface{})["sha"].(string)
		}
	}

	return d, p, s
}

func GetFiles() []map[string]interface{} {
	// 初始化请求与响应
	url := "https://api.github.com/repos/" +
		config.Conf.App.Owner + "/" +
		config.Conf.App.Repo +
		"/contents/" +
		config.Conf.App.Path +
		"?ref=" + config.Conf.App.Branch

	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer func() {
		// 用完需要释放资源
		fasthttp.ReleaseResponse(resp)
		fasthttp.ReleaseRequest(req)
	}()

	// 设置请求方法
	req.Header.SetMethod("GET")
	req.Header.SetBytesKV([]byte("Accept"), []byte("application/vnd.github.v3+json"))
	req.Header.SetBytesKV([]byte("Authorization"), []byte("token "+config.Conf.App.Token))
	// 设置请求的目标网址
	req.SetRequestURI(url)

	// 发起请求
	if err := fasthttp.Do(req, resp); err != nil {
		log.Println(" 请求失败:", url, err.Error())
	}

	// 获取响应的数据实体
	body := resp.Body()

	//log.Println(string(body),url)

	var mapResult []map[string]interface{}

	err := json.Unmarshal(body, &mapResult)
	if err != nil {
		fmt.Println("JsonToMapDemo err: ", err)
	}

	for _, v := range mapResult {
		v["download_url"] = "https://cdn.jsdelivr.net/gh/" + config.Conf.App.Owner + "/" + config.Conf.App.Repo + "@" + config.Conf.App.Branch + "/" + v["path"].(string)
	}
	return mapResult
}

func DelFile(filepath, sha string) string {
	// 初始化请求与响应
	url := "https://api.github.com/repos/" +
		config.Conf.App.Owner + "/" +
		config.Conf.App.Repo +
		"/contents/" + filepath

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
	req.Header.SetBytesKV([]byte("Accept"), []byte("application/vnd.github.v3+json"))
	req.Header.SetBytesKV([]byte("Authorization"), []byte("token "+config.Conf.App.Token))

	// 设置请求的目标网址
	req.SetRequestURI(url)

	args := make(map[string]string)
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
