package heap2

import "fmt"

// 本例为最小堆
// 最大堆只需要修改less函数即可
type Heap []int

func (h Heap) swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h Heap) less(i, j int) bool {
	return h[i] < h[j]
}

// 自下而上进行堆化
func (h Heap) up(i int) {
	for {
		f := (i - 1) / 2 // 父亲结点
		if i == f || h.less(f, i) {
			break
		}
		h.swap(f, i)
		i = f
	}
}

// 插入
// 注意go中所有参数转递都是值传递
// 所以要让h的变化在函数外也起作用，此处得传指针
func (h *Heap) Push(x int) {
	*h = append(*h, x)
	h.up(len(*h) - 1)
}

//自上而下进行对话
func (h Heap) down(i int) {
	for {
		l := 2*i + 1 // 左孩子
		if l >= len(h) {
			break // i已经是叶子结点了
		}
		j := l
		if r := l + 1; r < len(h) && h.less(r, l) { //比较左右子节点那个小，父亲节点就和那个进行交换
			j = r // 右孩子
		}
		if h.less(i, j) {
			break // 如果父结点比孩子结点小，则不交换
		}
		h.swap(i, j) // 交换父结点和子结点
		i = j        //继续向下比较
	}
}

// 删除堆中位置为i的元素
// 返回被删元素的值
func (h *Heap) Remove(i int) (int, bool) {
	if i < 0 || i > len(*h)-1 {
		return 0, false
	}
	n := len(*h) - 1
	h.swap(i, n) // 用最后的元素值替换被删除元素
	// 删除最后的元素
	x := (*h)[n]
	*h = (*h)[0:n]
	// 如果当前元素大于父结点，向下筛选
	if (*h)[i] > (*h)[(i-1)/2] {
		h.down(i)
	} else { // 当前元素小于父结点，向上筛选
		h.up(i)
	}
	return x, true
}

//弹出堆顶 当i=0 时，remove就是 pop
// 弹出堆顶的元素，并返回其值
func (h *Heap) Pop() int {
	n := len(*h) - 1
	h.swap(0, n)
	x := (*h)[n]
	*h = (*h)[0:n]
	h.down(0)
	return x
}

//初始化
//用数组进行构建一个堆
// 将所有数据放到一个数组中，此时形成的完全二叉树不是最小堆，但是仅包含叶子的子树已经是堆
func (h Heap) Init() {
	n := len(h)
	// i > n/2-1 的结点为叶子结点本身已经是堆了
	for i := n/2 - 1; i >= 0; i-- {
		h.down(i)
	}
}

//建堆
//从下到上调整堆的过程，从 i = len(arr) /2 开始依次向上调整，i = len(arr) /2是堆中末尾节点的父节点， i=0是根节点。
func BuildHeap(arr []int) Heap {
	h := Heap(arr)
	n := len(h)
	// 从第一个非叶子节点，到根节点
	for i := n/2 - 1; i >= 0; i-- {
		h.down(i)
	}

	return h
}

//堆排序
//采用最小堆的方式 每从堆顶弹出一个元素（这个元素就是堆的最小值，即可实现升序排序）
func HeapSort(arr []int) {
	// 创建堆
	heap := BuildHeap(arr)
	var sortedArr []int
	for len(heap) > 0 {
		sortedArr = append(sortedArr, heap.Pop())
	}

	fmt.Println(sortedArr)
}
