package main

import (
	"fmt"
	analyse "logAnalysis/Service"
)

func main() {
	quantity := 10
	res := analyse.Analyze_requests("./access.log", quantity)
	fmt.Println("result:", res)
}
