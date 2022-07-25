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
