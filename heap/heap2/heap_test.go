package heap2

import (
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("开始测试")
	m.Run()

}

func TestHeap(t *testing.T) {
	fmt.Println("开始测试heap2中相关方法")
	t.Run("测试push方法", testPush)
}

func testPush(t *testing.T) {
	var h = Heap{20, 7, 3, 10, 15, 25, 30, 17, 19}
	h.Init()
	h.Push(6)
	fmt.Println(h)

	x, ok := h.Remove(5)
	fmt.Println(x, ok, h)

	y, ok := h.Remove(1)
	fmt.Println(y, ok, h) // 6 true [3 7 15 10 19 20 30 17]

	z := h.Pop()
	fmt.Println(z, h) // 3 [7 10 15 17 19 20 30]
}
