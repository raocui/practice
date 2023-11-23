package map1_test

import (
	"_/gopl.io2/ch9/map1"
	"testing"
	"time"
)

// 竞态检测器  go test -race .\map1_test.go
func TestIcon(t *testing.T) {

	done := make(chan struct{})
	go func() {
		done <- struct{}{}

		map1.Icon("aaaa", "11")
	}()
	go func() {
		time.Sleep(1 * time.Second)
		map1.Icon("aaaa", "22")

		done <- struct{}{}

	}()
	<-done
	<-done
	map1.GetIcons()

}
