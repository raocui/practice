package zhan1_test

import (
	"_/gopl.io2/ch9/zhan1"
	"fmt"
	"testing"
)

func TestGetMessage(t *testing.T) {
	a := zhan1.A{
		Message: "kakakakak",
	}
	fmt.Println(a.GetMessage())
	
	fmt.Println(a.GetMessage2())
}
