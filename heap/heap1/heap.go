package heap
import ("fmt")



type Node struct {
	Value int
	Key   string
}

type Heap struct {
	list   []*Node
	length int
}



//node slice 交换
func (h *Heap) SliceNodeSwap(i int, j int) {
	x := h.list[i]
	h.list[i] = h.list[j]
	h.list[j] = x
}



//自上而下堆化
func (h *Heap) down(length int,pos int){
	for {
		maxPos := pos
		//因为是大顶堆 所以要先找到左右子节点 最大的那个数据 
		if pos*2 < length && h.list[pos].Value<h.list[pos*2].Value {
			maxPos = pos * 2 
		}

		if pos*2+1<length && h.list[maxPos].Value < h.list[pos*2+1].Value {
			maxPos = pos *2 +1 
		}

		// 如果maxPos == pos 就代表不能向下交换了 就结束向下比较 
		if maxPos == pos  {  
			break
		}

		h.SliceNodeSwap(pos,maxPos)
		//继续向下比较 
		pos = maxPos    
	}
}

//自下而上进行比较 
func (h *Heap) up(length int){
	if length <1  {
		return
	}

	if  length == 2  {
		if h.list[length].Value > h.list[length-1].Value {
			h.SliceNodeSwap(length,length-1)
		}
		return 
	}

	//与父节点进行比较 
	i := length
	for i/2 > 0 && h.list[i].Value >h.list[i/2].Value {
		h.SliceNodeSwap(i,i/2)
		i = i/2
	}
}


//插入操作 ，基于向上堆化 
func (h *Heap) InsertHeap(one *Node){
	h.list = append(h.list, one)
	length := len(h.list)
	h.length = length -1 
	h.up(h.length)
}
//输出heap
func heapShow(heaps []*Node){
	for one,value := range heaps{
		fmt.Println(one,value)
	}
}










//获取堆顶
func (h *Heap) GetTop() *Node{
	if h.length == 0{
		panic("heap is empty")
	}

	top := h.list[1]
	//堆顶与堆底进行交换 
	h.SliceNodeSwap(1,len(h.list)-1)
	length := len(h.list) -2 
	fmt.Println(length)
	h.down(length,1)
	heapShow(h.list)
	h.list = append(h.list[:length+1], h.list[length+2:]...)
	h.length--
	return top
}


//创建堆 将一个数组 变成堆
func CreatgeHeap(){
	arrList := []int{1, 2, 11, 3, 7, 8, 4, 5}
	var myHeap Heap
	myHeap.list = append(myHeap.list, &Node{})
	for _,value := range arrList{
		tmp := Node{}
		tmp.Value = value
		myHeap.InsertHeap(&tmp)
	}
	fmt.Println(myHeap.list)
	
}