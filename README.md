English | [简体中文](README-CN.md)
![](https://aliyunsdk-pages.alicdn.com/icons/AlibabaCloud.svg)

## Alibaba Cloud DingTalk SDK

## Requirements
- It's necessary for you to make sure your system have installed a Go environment which is new than 1.12.0.

## Installation
If you use `go mod` to manage your dependence, you can use the following command:

```sh
$ go get github.com/alibabacloud-go/dingtalk-sdk
```

## Demo
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

## Issues
[Opening an Issue](https://github.com/aliyun/dingtalk-sdk/issues/new), Issues not conforming to the guidelines may be closed immediately.

## References
* [Latest Release](https://github.com/aliyun/dingtalk-sdk)

## License
[Apache-2.0](http://www.apache.org/licenses/LICENSE-2.0)

Copyright (c) 2009-present, Alibaba Cloud All rights reserved.