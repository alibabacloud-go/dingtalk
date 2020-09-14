[English](README.md) | 简体中文

![](https://aliyunsdk-pages.alicdn.com/icons/AlibabaCloud.svg)

## Alibaba Cloud DingTalk SDK

## 要求
- 您需要确保本地安装的 go 环境版本大于 1.12.0.

## 安装

你可以使用 `go mod` 来管理你的依赖，
```sh
$  go get github.com/alibabacloud-go/dingtalk-sdk
```

## 使用示例
```go
package main

import (
	"fmt"
	"os"

	"github.com/alibabacloud-go/tea/tea"
	ding "github.com/alibabacloud-go/dingtalk-sdk/client"
	fileform "github.com/alibabacloud-go/tea-fileform/service"
)

func main() {
	config := new(ding.Config).
		SetServerUrl("https://oapi.dingtalk.com").
		SetSession("access_token")

	client, err := ding.NewClient(config)
	if err != nil {
		panic(err)
	}

	file, err := os.Open("./demo.jpg")
	if err != nil {
		panic(err)
	}

	fileObj := new(fileform.FileField).
		SetFilename("demo.jpg").
		SetContent(file).
		SetContentType("")

	params := new(ding.UploadOapiMediaParams).
		SetType("image").
		SetMedia(fileObj)

	req := new(ding.UploadOapiMediaRequest).SetParams(params)
	resp, err := client.UploadOapiMedia(req)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
}
```

## 问题
[Opening an Issue](https://github.com/aliyun/dingtalk-sdk/issues/new)，不符合指南的问题可能会立即关闭。

## 相关
* [Latest Release](https://github.com/aliyun/dingtalk-sdk)

## 许可证
[Apache-2.0](http://www.apache.org/licenses/LICENSE-2.0)

Copyright (c) 2009-present, Alibaba Cloud All rights reserved.
