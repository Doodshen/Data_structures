package main

import (
	"DataStructures/tool"

	"fmt"
)

func InsertSort(list []int) {
	if len(list) <= 1 {
		return
	}

	for i := 0; i < len(list); i++ {
		value := list[i] //未排序列找到需要插入的数据
		j := i - 1
		//查找插入的位置
		for ; j >= 0; j-- { // 从已排序列找到需要插入的位置
			if list[j] > value {
				list[j+1] = list[j] //数据移动
			} else {
				break
			}
		}
		list[j+1] = value //插入数据 第一次j通过--以后变成了-1
	}
}

//下边的做法就太复杂了

func standInsertSort(list []int) []int {
	resultList := []int{}
	length := len(list)
	i := 1
	resultList = append(resultList, list[0]) // 首先定义 已排序列第一个数据
	for i < length {
		tool.PrintSlice(resultList)
		for j := 0; j < len(resultList); j++ { //从已排序列中找位置
			if list[i] < resultList[j] {
				resultList = insertList(resultList, j, list[i])
				break
			}
			if j == len(resultList)-1 && list[i] > resultList[j] { //最后一个未排序列的数据
				resultList = insertList(resultList, j+1, list[i])
			}
		}
		i++
	}
	return resultList
}

func insertList(list []int, i int, x int) []int {
	returnlist := []int{}
	n := 0
	if i == len(list) {
		returnlist = append(list, x)
		return returnlist
	}
	for n < len(list) {
		if n < i { //先将需要插入的位置前面的数据插入
			returnlist = append(returnlist, list[n])
		} else if n == i { //将插入位置的元素先插入 然后再插入原本位置上的数据
			returnlist = append(returnlist, x)
			returnlist = append(returnlist, list[n])
		} else { // 将剩余元素插入
			returnlist = append(returnlist, list[n])
		}
		n++
	}
	return returnlist
}

func main() {
	arrList := []int{8, 1, 2, 11, 3, 6, 8, 4, 9, 343, 3}
	InsertSort(arrList)
	fmt.Println("InsertSort:",arrList)
	
}
