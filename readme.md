# 创建一个自己的包

具体步骤如下：

1. `git init`创建一个新的git地址仓库，这里试过如果是已有的仓库对于小白好像有问题，怎么解决后续再看
2. 使用`go mod init go_node_api`创建一个名为`go_node_api`的包
3. 编写如下代码：
```go
package city_output

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func CurrentCity(cityCode string) string {
	fmt.Println("开始请求")
	resp, err := http.Get("http://127.0.0.1:3000/test/dataOutput?id=" + cityCode)

	if err != nil {
		fmt.Println("请求失败", err)
		panic(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("网络请求失败：", err)
		panic(err)
	}

	return string(body)
}
```
4. 在`main.go`中进行导入并使用如下：

```go
package main

import (
	"fmt"
	"go_node_api/city_output"
)

func main() {
	body := city_output.CurrentCity("Hang Zhou")
	fmt.Println(body)
}
```

**注意点**：
1. 子模块导出的时候，直接申明包名`city_output`, 后续通过这个`city_output`来引用对应包内部的方法
2. 导入时，这里导入的名称是`当前包名` + `子模块名`
