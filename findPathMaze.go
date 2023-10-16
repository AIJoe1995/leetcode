package main

import "fmt"

// 参考数据结构与算法python实现
// 状态空间搜索问题 或 路径搜索问题

func mark(maze [][]int, pos [2]int) {
	// 把maze处标记为一走过
	maze[pos[0]][pos[1]] = 2
}

func passable(maze [][]int, pos [2]int) bool {
	rows := len(maze)
	cols := len(maze[0])
	if pos[0] < 0 || pos[0] >= rows || pos[1] < 0 || pos[1] >= cols {
		return false
	}
	return maze[pos[0]][pos[1]] == 0
}

var directions = [4][2]int{
	{-1, 0}, // 上
	{1, 0},  // 下
	{0, -1}, // 左
	{0, 1},  // 右
}

// maze 0表示可通行 1表示不可通行 2表示已经走过该结点
func find_path(maze [][]int, pos [2]int, end [2]int) bool {
	// 递归调用
	// 递归 借助系统栈 来保存信息
	// 当前位置可以有多个探查方向，但是每次只能探查一种可能，需要记录不能立即考察的其他可能方向
	mark(maze, pos) // 避免重复探查同一个位置 进入死循环
	if pos == end { // 递归调用的终止条件
		fmt.Println(pos)
		return true
	}
	for _, direction := range directions {
		nextpos := [2]int{pos[0] + direction[0], pos[1] + direction[1]}
		if passable(maze, nextpos) { // 能向前走就一直向前走 如果不能继续向前走，就在当前结点走前后左右四个可选方向的下一个方向
			if find_path(maze, nextpos, end) {
				fmt.Println(nextpos)
				return true
			}
		}
	}
	return false
}

type st struct {
	pos [][2]int
	idx []int
}

func NewSt() *st {
	return &st{}
}

func (st *st) pop() ([2]int, int) {
	pos := st.pos[len(st.pos)-1]
	idx := st.idx[len(st.pos)-1]
	st.pos = st.pos[:len(st.pos)-1]
	st.idx = st.idx[:len(st.pos)-1]
	return pos, idx
}

func (st *st) push(pos [2]int, startidx int) {
	st.pos = append(st.pos, pos)
	st.idx = append(st.idx, startidx)
}

// 基于栈的路径搜索 栈的特性 后进先出 会一直向前走
// 深度优先搜索
func find_path_V1(maze [][]int, start [2]int, end [2]int) bool {
	// 将递归转换为递推 使用循环完成 利用自己创建的栈保存信息
	// 栈里面保存的是所有可能 栈不为空就继续找下去
	if start == end {
		fmt.Println(end)
		return true
	}
	st := NewSt()
	mark(maze, start)
	st.push(start, 0) // 栈里面保存当前在探查的结点 以及上下左右四个待选方向 还剩几种可能没有探查
	for len(st.pos) > 0 {
		pos, nxt := st.pop()
		for i := nxt; i < 4; i++ {
			nextpos := [2]int{
				pos[0] + directions[i][0],
				pos[1] + directions[i][1],
			}
			if nextpos == end {
				fmt.Println("path found")
				return true
			}
			if passable(maze, nextpos) {
				st.push(pos, i+1) // 当前结点i位置已经尝试过了，下一次从栈取出来 应该从当前结点的i+1选择尝试
				mark(maze, nextpos)
				st.push(nextpos, 0) // 这个结点可达 把这个结点相连的四个待探查的方向放入栈
				break
			}
		}
	}
	return false
}

type queue struct {
	pos [][2]int
}

func NewQueue() *queue {
	return &queue{}
}

func (qu *queue) enqueue(pos [2]int) {
	qu.pos = append(qu.pos, pos)
}

func (qu *queue) dequeue() [2]int {
	res := qu.pos[0]
	qu.pos = qu.pos[1:]
	return res
}

// 基于队列的路径搜索 队列的特性 先进先出 先考虑是离开始位置近的位置
// 广度优先搜索
func find_path_V2(maze [][]int, start [2]int, end [2]int) bool {
	if start == end {
		fmt.Println("path found")
		return true
	}
	qu := NewQueue()
	mark(maze, start)
	qu.enqueue(start)
	for len(qu.pos) > 0 {
		pos := qu.dequeue()
		for i := 0; i < 4; i++ {
			nextpos := [2]int{
				pos[0] + directions[i][0],
				pos[1] + directions[i][1],
			}
			if passable(maze, nextpos) {
				if nextpos == end {
					fmt.Println("path found")
					return true
				}
				mark(maze, nextpos)
				qu.enqueue(nextpos)

			}
		}
	}
	return false
}

func main() {
	maze := [][]int{
		{0, 0, 0, 1},
		{0, 1, 0, 0},
	}
	start := [2]int{0, 0}
	end := [2]int{1, 3}
	//ok := find_path(maze, start, end)
	//ok := find_path_V1(maze, start, end)
	ok := find_path_V2(maze, start, end)
	fmt.Println("result,", ok)
}
