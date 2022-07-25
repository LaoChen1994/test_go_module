package main

import (
	"fmt"
	"go_node_api/city_output"
)

func main() {
	body := city_output.CurrentCity("Hang Zhou")
	fmt.Println(body)
}
