package heap

import (
	"fmt"
	"testing"
)



func TestMain(m *testing.M){
	fmt.Print("测试开始")
	m.Run()
}

func TestHeap(t *testing.T) {
	fmt.Println("开始测试heap中相关方法")
	t.Run("测试创建创建堆",testcreate)
}


func testcreate(t *testing.T){
	CreatgeHeap()
}