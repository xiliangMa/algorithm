package t_test

import (
	"testing"
	"time"
)
func Test_Channel(t *testing.T) {
	a := make(chan int)

	go func(){
		t.Log("发送数据开始")
		for i := 0; i < 10; i++ {
			a <- i
			time.Sleep(time.Second)
		}
		t.Log("发送数据结束")

	}()

	for data := range a {
		t.Log("数据接收如下：", data)
		if data == 9 {
			break
		}
	}
}
