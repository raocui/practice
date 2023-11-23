package cake

import (
	"fmt"
	"time"
)

//串行受限
//在下面的例子中，cakes 是串行受限的，首先受限于 baker goroutine，然后受限于 icer goroutine。

type Cake struct{ state string }

var cooked = make(chan *Cake)
var iced = make(chan *Cake)

func baker(cooked chan<- *Cake) {
	for {
		cake := new(Cake)
		cake.state = "cooked"
		time.Sleep(10 * time.Second)
		cooked <- cake // baker 不再访问 cake 变量
	}
}

func icer(iced chan<- *Cake, cooked <-chan *Cake) {
	for cake := range cooked {
		cake.state = "iced"
		iced <- cake //icer 不再访问 cake 变量
	}
}

func MakeCake() {

	go baker(cooked)
	go icer(iced, cooked)
	for v := range iced {
		fmt.Printf("蛋糕：%#v", v)
	}
}
