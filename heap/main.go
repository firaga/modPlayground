package main

import (
	"container/heap"
	"fmt"
)

// 1. 定义堆类型：基于 int 切片的小顶堆
type IntMinHeap []int

// 2. 实现 heap.Interface 接口的 5 个方法

// Len：返回堆的长度
func (h IntMinHeap) Len() int { return len(h) }

// Less：小顶堆规则：a[i] < a[j] 表示 i 元素优先级更高（在堆顶）
func (h IntMinHeap) Less(i, j int) bool { return h[i] < h[j] }

// Swap：交换 i 和 j 位置的元素
func (h IntMinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

// Push：向堆中添加元素（将元素追加到切片末尾，堆包会自动调整堆结构）
func (h *IntMinHeap) Push(x interface{}) {
	// x 是 interface{} 类型，需断言为 int
	*h = append(*h, x.(int))
}

// Pop：从堆中移除并返回堆顶元素（堆包会先将堆顶元素交换到切片末尾，再调用此方法）
func (h *IntMinHeap) Pop() interface{} {
	// 1. 保存切片末尾元素（即堆包调整后的“待移除元素”）
	old := *h
	n := len(old)
	x := old[n-1]
	// 2. 缩小切片长度（移除末尾元素）
	*h = old[:n-1]
	// 3. 返回移除的元素
	return x
}

func main() {
	// 1. 初始化一个普通切片（未满足堆结构）
	nums := []int{5, 3, 8, 1, 2}
	// 2. 转换为堆类型，并初始化堆结构（O(n) 时间复杂度）
	h := IntMinHeap(nums)
	heap.Init(&h)
	fmt.Println("初始化后的小顶堆（底层切片）：", h) // 输出：[1 2 8 5 3]（堆顶是 1，满足小顶堆）

	// 3. 向堆中添加元素（O(log n) 时间复杂度）
	heap.Push(&h, 0)
	fmt.Println("添加元素 0 后的堆：", h) // 输出：[0 2 1 5 3 8]（堆顶更新为 0）

	// 4. 从堆中取出堆顶元素（O(log n) 时间复杂度）
	for h.Len() > 0 {
		// Pop 会返回堆顶元素，并自动调整堆结构
		top := heap.Pop(&h).(int)
		fmt.Printf("取出堆顶元素：%d，剩余堆：%v\n", top, h)
	}
}
