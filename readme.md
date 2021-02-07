# gitee-image-hosting 🐽

> Gitee 图床工具 (基于 Golang(Gin) 实现) [github地址](https://github.com/hezhizheng/gitee-image-hosting)

```
还是用Gitee当图床算了，不然哪里有国内访问又快又免费又稳的图床服务提供......
```

## 页面
![](https://gitee.com/hezhizheng/pictest/raw/master/newdir/20210207160931_ZPUDMEHPYYLWOQFJ.png)
![](https://img.vim-cn.com/b8/ba7180c764628025bb4153a9c3995bb5ace0c9.gif)

## 功能
- 一键启动，跨平台支持，运行只依赖编译后的二进制文件
- 可视化web操作界面(PS: 页面有点丑，但基本能用......)
- 多图上传，支持 'jpeg', 'jpg', 'gif', 'png' 格式
- 复制图片url 、删除图片

## 使用
windows用户可直接下载 [release](https://github.com/hezhizheng/gitee-image-hosting/release) 文件启动即可，参数说明：
![](https://gitee.com/hezhizheng/pictest/raw/master/image-hosting/20210207154953_ZHKKGZZVAYDEZHAO.png)
```
完整启动命令： ./gitee-image-hosting.exe -owner hezhizheng -repo pictest -path image-hosting -token xxxtoken -port 2047
实际参数替换成自己的就行
```
token获取：https://gitee.com/profile/personal_access_tokens/new

其他品台自行编译 `go build`


## 关于Gitee限制图片大于1M访问的处理方案
- 使用第三方图片压缩工具进行压缩，之后再进行上传。推荐 [iloveimg](https://www.iloveimg.com/zh-cn/compress-image/compress-jpg)
- 启用Gitee的pages功能(非付费用户上传图片之后需要手动进行pages服务的部署)，程序会自动替换pages域名进行图片的展示。


## License
[MIT](./LICENSE.txt)