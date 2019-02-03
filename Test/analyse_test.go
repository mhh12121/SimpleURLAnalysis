package test

import (
	"fmt"
	analyse "logAnalysis/Service"
	"testing"
)

func Test_Analyse(t *testing.T) {
	res := analyse.Analyze_requests("../access.log", 10)
	fmt.Println(res)
	res2 := analyse.Analyze_requests("../access.log", 5)
	fmt.Println(res2)
}
