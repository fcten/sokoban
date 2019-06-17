package main

import (
	"fmt"
	"runtime"
)

// 墙体 '#'
// 空地 ' '
// 箱子 '$'
// 小人 '@'
// 目标 '.'
// 在目标上的箱子 '*'
// 在目标上的小人 '%'

type vector2d struct {
	x int
	y int
}

type node struct {
	puzzle *sokoban
	curMap [][]byte
	// 当前走法
	solution string
	// 剩余箱子数量
	count int
	// 小人位置
	pos vector2d
}

type sokoban struct {
	// 存储已抵达的状态
	maps map[string]bool

	root node
	// 地图尺寸
	size vector2d

	// bfs 队列
	queue []*node

	// 峰值内存
	memory uint64
}

func main() {
	var puzzle sokoban
	/*
		// 1
		// 4
		// uurr
		puzzle.root.curMap = [][]byte{
			[]byte{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
			[]byte{' ', ' ', '#', '#', '#', '#', '#', '#', '#', ' ', ' '},
			[]byte{' ', ' ', '#', ' ', ' ', '$', ' ', '.', '#', ' ', ' '},
			[]byte{' ', ' ', '#', ' ', ' ', '#', '#', '#', '#', ' ', ' '},
			[]byte{' ', ' ', '#', ' ', '@', '#', ' ', ' ', ' ', ' ', ' '},
			[]byte{' ', ' ', '#', '#', '#', '#', ' ', ' ', ' ', ' ', ' '},
			[]byte{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
		}
		// 2
		// 16
		// lurrlddrurudllur
		puzzle.root.curMap = [][]byte{
			[]byte{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
			[]byte{' ', ' ', ' ', ' ', '#', '#', '#', '#', ' ', ' ', ' '},
			[]byte{' ', ' ', '#', '#', '#', ' ', '.', '#', ' ', ' ', ' '},
			[]byte{' ', ' ', '#', ' ', '$', ' ', '.', '#', ' ', ' ', ' '},
			[]byte{' ', ' ', '#', ' ', '@', '$', ' ', '#', ' ', ' ', ' '},
			[]byte{' ', ' ', '#', '#', ' ', ' ', '#', '#', ' ', ' ', ' '},
			[]byte{' ', ' ', ' ', '#', '#', '#', '#', ' ', ' ', ' ', ' '},
		}
		// 3
		// 23
		// ulldldruurrdllrrddlurul
		puzzle.root.curMap = [][]byte{
			[]byte{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
			[]byte{' ', '#', '#', '#', '#', '#', '#', '#', '#', ' ', ' '},
			[]byte{' ', '#', ' ', ' ', ' ', ' ', ' ', ' ', '#', ' ', ' '},
			[]byte{' ', '#', ' ', '.', '*', '*', '$', '@', '#', ' ', ' '},
			[]byte{' ', '#', ' ', ' ', ' ', ' ', ' ', ' ', '#', ' ', ' '},
			[]byte{' ', '#', '#', '#', '#', '#', ' ', ' ', '#', ' ', ' '},
			[]byte{' ', ' ', ' ', ' ', ' ', '#', '#', '#', '#', ' ', ' '},
		}
		// 4
		// 16
		// rddlruulduullddr
		puzzle.root.curMap = [][]byte{
			[]byte{' ', ' ', '#', '#', '#', '#', '#', '#', ' ', ' ', ' '},
			[]byte{' ', ' ', '#', ' ', ' ', ' ', ' ', '#', ' ', ' ', ' '},
			[]byte{' ', ' ', '#', ' ', '#', '@', ' ', '#', ' ', ' ', ' '},
			[]byte{' ', ' ', '#', ' ', '$', '*', ' ', '#', ' ', ' ', ' '},
			[]byte{' ', ' ', '#', ' ', '.', '*', ' ', '#', ' ', ' ', ' '},
			[]byte{' ', ' ', '#', ' ', ' ', ' ', ' ', '#', ' ', ' ', ' '},
			[]byte{' ', ' ', '#', '#', '#', '#', '#', '#', ' ', ' ', ' '},
		}
		// 5
		// 33
		// ulurrrdlullddruluruuldrddrruldluu
		puzzle.root.curMap = [][]byte{
			[]byte{' ', ' ', '#', '#', '#', '#', ' ', ' ', ' ', ' ', ' '},
			[]byte{' ', ' ', '#', ' ', '.', '#', ' ', ' ', ' ', ' ', ' '},
			[]byte{' ', ' ', '#', ' ', ' ', '#', '#', '#', ' ', ' ', ' '},
			[]byte{' ', ' ', '#', '*', ' ', ' ', ' ', '#', ' ', ' ', ' '},
			[]byte{' ', ' ', '#', ' ', ' ', '$', ' ', '#', ' ', ' ', ' '},
			[]byte{' ', ' ', '#', ' ', '@', '#', '#', '#', ' ', ' ', ' '},
			[]byte{' ', ' ', '#', '#', '#', '#', ' ', ' ', ' ', ' ', ' '},
		}
		// 6
		// 25
		// lurlldrdrdrruulldluddllur
		puzzle.root.curMap = [][]byte{
			[]byte{' ', ' ', '#', '#', '#', '#', '#', '#', '#', ' ', ' '},
			[]byte{' ', ' ', '#', ' ', ' ', ' ', ' ', ' ', '#', ' ', ' '},
			[]byte{' ', ' ', '#', ' ', '.', '$', '.', ' ', '#', ' ', ' '},
			[]byte{' ', '#', '#', ' ', '$', '@', '$', ' ', '#', ' ', ' '},
			[]byte{' ', '#', ' ', ' ', '.', '$', '.', ' ', '#', ' ', ' '},
			[]byte{' ', '#', ' ', ' ', ' ', ' ', ' ', ' ', '#', ' ', ' '},
			[]byte{' ', '#', '#', '#', '#', '#', '#', '#', '#', ' ', ' '},
		}
		// 7
		// 21
		// ulldrrurdlllddrruurul
		puzzle.root.curMap = [][]byte{
			[]byte{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
			[]byte{' ', ' ', '#', '#', '#', '#', '#', '#', ' ', ' ', ' '},
			[]byte{' ', ' ', '#', ' ', '.', ' ', ' ', '#', ' ', ' ', ' '},
			[]byte{' ', ' ', '#', ' ', '$', '@', ' ', '#', ' ', ' ', ' '},
			[]byte{' ', ' ', '#', ' ', '#', '$', '.', '#', ' ', ' ', ' '},
			[]byte{' ', ' ', '#', ' ', ' ', ' ', ' ', '#', ' ', ' ', ' '},
			[]byte{' ', ' ', '#', '#', '#', '#', '#', '#', ' ', ' ', ' '},
		}
		// 8
		// 38
		// rluurdrrulldrrddluulldrurruldrddlurull
		puzzle.root.curMap = [][]byte{
			[]byte{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
			[]byte{' ', ' ', '#', '#', '#', '#', '#', '#', ' ', ' ', ' '},
			[]byte{' ', ' ', '#', '.', '*', ' ', ' ', '#', ' ', ' ', ' '},
			[]byte{' ', ' ', '#', '.', '$', ' ', ' ', '#', ' ', ' ', ' '},
			[]byte{' ', ' ', '#', '@', '$', ' ', ' ', '#', ' ', ' ', ' '},
			[]byte{' ', ' ', '#', '#', '#', ' ', ' ', '#', ' ', ' ', ' '},
			[]byte{' ', ' ', ' ', ' ', '#', '#', '#', '#', ' ', ' ', ' '},
		}
		// 9
		// 41
		// dlddrrruuulrdddllluurrurdlllddrulurrllurr
		puzzle.root.curMap = [][]byte{
			[]byte{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
			[]byte{' ', ' ', '#', '#', '#', '#', '#', '#', ' ', ' ', ' '},
			[]byte{' ', ' ', '#', ' ', '@', '$', '.', '#', ' ', ' ', ' '},
			[]byte{' ', ' ', '#', ' ', '$', '$', '.', '#', ' ', ' ', ' '},
			[]byte{' ', ' ', '#', ' ', ' ', '#', '.', '#', ' ', ' ', ' '},
			[]byte{' ', ' ', '#', ' ', ' ', ' ', ' ', '#', ' ', ' ', ' '},
			[]byte{' ', ' ', '#', '#', '#', '#', '#', '#', ' ', ' ', ' '},
		}
		// 10
		// 23
		// uullldurrrddllulurdrdru
		puzzle.root.curMap = [][]byte{
			[]byte{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
			[]byte{' ', ' ', '#', '#', '#', '#', '#', '#', ' ', ' ', ' '},
			[]byte{' ', ' ', '#', ' ', ' ', '.', '.', '#', ' ', ' ', ' '},
			[]byte{' ', ' ', '#', '$', '$', '$', ' ', '#', ' ', ' ', ' '},
			[]byte{' ', ' ', '#', '.', ' ', ' ', '@', '#', ' ', ' ', ' '},
			[]byte{' ', ' ', '#', '#', '#', '#', '#', '#', ' ', ' ', ' '},
			[]byte{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
		}
		// 11
		// 52
		// urdrrurrdlllluldrdlurrruruulddrdlllurrurdrdlllulduld
		puzzle.root.curMap = [][]byte{
			[]byte{' ', ' ', ' ', ' ', ' ', '#', '#', '#', '#', ' ', ' '},
			[]byte{' ', ' ', ' ', ' ', ' ', '#', ' ', ' ', '#', ' ', ' '},
			[]byte{' ', ' ', '#', '#', '#', '#', ' ', ' ', '#', '#', ' '},
			[]byte{' ', '#', '#', ' ', '$', ' ', '$', ' ', ' ', '#', ' '},
			[]byte{' ', '#', ' ', '@', ' ', '$', ' ', ' ', ' ', '#', ' '},
			[]byte{' ', '#', '.', '.', '.', '#', '#', '#', '#', '#', ' '},
			[]byte{' ', '#', '#', '#', '#', '#', ' ', ' ', ' ', ' ', ' '},
		}
		// 12
		// 51
		// uulllldddrrulluurrrrddlllrrruulldurrddlldllurruulld
		puzzle.root.curMap = [][]byte{
			[]byte{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
			[]byte{' ', ' ', '#', '#', '#', '#', '#', '#', '#', ' ', ' '},
			[]byte{' ', ' ', '#', ' ', ' ', ' ', ' ', ' ', '#', ' ', ' '},
			[]byte{' ', ' ', '#', ' ', '#', ' ', '#', ' ', '#', ' ', ' '},
			[]byte{' ', ' ', '#', '.', ' ', '$', '*', '@', '#', ' ', ' '},
			[]byte{' ', ' ', '#', ' ', ' ', ' ', '#', '#', '#', ' ', ' '},
			[]byte{' ', ' ', '#', '#', '#', '#', '#', ' ', ' ', ' ', ' '},
		}
		// 13
		// 37
		// drdddlluuddrruulullldllurrrrdrrddlluu
		puzzle.root.curMap = [][]byte{
			[]byte{' ', ' ', ' ', ' ', ' ', ' ', '#', '#', '#', ' ', ' '},
			[]byte{' ', '#', '#', '#', '#', '#', '#', '@', '#', '#', ' '},
			[]byte{' ', '#', ' ', ' ', ' ', ' ', '.', '*', ' ', '#', ' '},
			[]byte{' ', '#', ' ', ' ', ' ', '#', ' ', ' ', ' ', '#', ' '},
			[]byte{' ', '#', '#', '#', '#', '#', '$', '#', ' ', '#', ' '},
			[]byte{' ', ' ', ' ', ' ', ' ', '#', ' ', ' ', ' ', '#', ' '},
			[]byte{' ', ' ', ' ', ' ', ' ', '#', '#', '#', '#', '#', ' '},
		}
		// 14
		// 25
		// urrlddrurruuldlrrdllrrdll
		puzzle.root.curMap = [][]byte{
			[]byte{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
			[]byte{' ', ' ', ' ', ' ', ' ', '#', '#', '#', '#', ' ', ' '},
			[]byte{' ', ' ', '#', '#', '#', '#', ' ', ' ', '#', ' ', ' '},
			[]byte{' ', ' ', '#', ' ', '.', '$', ' ', ' ', '#', ' ', ' '},
			[]byte{' ', ' ', '#', '@', '.', '$', ' ', ' ', '#', ' ', ' '},
			[]byte{' ', ' ', '#', ' ', '.', '$', ' ', ' ', '#', ' ', ' '},
			[]byte{' ', ' ', '#', '#', '#', '#', '#', '#', '#', ' ', ' '},
		}
		// 15
		// 71
		// ulllldddrruurullrdddlluuluurdrrrrrdllullldlddrrruurullrdddllluurduluurd
		puzzle.root.curMap = [][]byte{
			[]byte{' ', '#', '#', '#', '#', ' ', ' ', ' ', ' ', ' ', ' '},
			[]byte{' ', '#', ' ', ' ', '#', '#', '#', '#', '#', '#', ' '},
			[]byte{' ', '#', ' ', ' ', ' ', ' ', ' ', ' ', ' ', '#', ' '},
			[]byte{' ', '#', ' ', '.', '#', '$', '$', '@', ' ', '#', ' '},
			[]byte{' ', '#', ' ', ' ', '#', ' ', '#', '#', '#', '#', ' '},
			[]byte{' ', '#', ' ', '.', ' ', ' ', '#', ' ', ' ', ' ', ' '},
			[]byte{' ', '#', '#', '#', '#', '#', '#', ' ', ' ', ' ', ' '},
		}
		// 16
		// 44
		// llldlludlrrrulllrrrrruullldlddrrullrrrrdllll
		puzzle.root.curMap = [][]byte{
			[]byte{' ', ' ', ' ', '#', '#', '#', '#', '#', '#', ' ', ' '},
			[]byte{' ', '#', '#', '#', ' ', ' ', ' ', ' ', '#', ' ', ' '},
			[]byte{'#', '#', '.', ' ', '$', '#', '#', ' ', '#', '#', ' '},
			[]byte{'#', '.', '.', '$', ' ', '$', ' ', ' ', '@', '#', ' '},
			[]byte{'#', '.', '.', ' ', '$', ' ', '$', ' ', '#', '#', ' '},
			[]byte{'#', '#', '#', '#', '#', '#', ' ', ' ', '#', ' ', ' '},
			[]byte{' ', ' ', ' ', ' ', ' ', '#', '#', '#', '#', ' ', ' '},
		}
		// 17
		// 17
		// dlludluluurdrddlu
		puzzle.root.curMap = [][]byte{
			[]byte{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
			[]byte{' ', ' ', '#', '#', '#', '#', ' ', ' ', ' ', ' ', ' '},
			[]byte{' ', ' ', '#', ' ', ' ', '#', '#', '#', '#', ' ', ' '},
			[]byte{' ', ' ', '#', ' ', '.', ' ', '.', ' ', '#', ' ', ' '},
			[]byte{' ', ' ', '#', ' ', '$', '$', '#', '@', '#', ' ', ' '},
			[]byte{' ', ' ', '#', '#', ' ', ' ', ' ', ' ', '#', ' ', ' '},
			[]byte{' ', ' ', ' ', '#', '#', '#', '#', '#', '#', ' ', ' '},
		}
		// 18
		// 47
		// ldddrrululurrrurddullldddluulurrrrdrullldddrrru
		puzzle.root.curMap = [][]byte{
			[]byte{' ', ' ', ' ', ' ', ' ', '#', '#', '#', '#', '#', ' '},
			[]byte{' ', '#', '#', '#', '#', '#', ' ', '.', ' ', '#', ' '},
			[]byte{' ', '#', ' ', ' ', '@', ' ', ' ', ' ', ' ', '#', ' '},
			[]byte{' ', '#', ' ', ' ', '$', '#', ' ', '.', ' ', '#', ' '},
			[]byte{' ', '#', '#', ' ', '$', ' ', '#', ' ', '#', '#', ' '},
			[]byte{' ', ' ', '#', ' ', ' ', ' ', ' ', ' ', '#', ' ', ' '},
			[]byte{' ', ' ', '#', '#', '#', '#', '#', '#', '#', ' ', ' '},
		}
		// 19
		// 54
		// ruullurdrdddlluudrruululldrddrruurullrdddlluururdlllur
		puzzle.root.curMap = [][]byte{
			[]byte{' ', ' ', '#', '#', '#', '#', '#', '#', '#', ' ', ' '},
			[]byte{' ', ' ', '#', ' ', ' ', '*', ' ', ' ', '#', ' ', ' '},
			[]byte{' ', ' ', '#', ' ', ' ', ' ', ' ', ' ', '#', ' ', ' '},
			[]byte{' ', ' ', '#', '#', ' ', '#', '.', '#', '#', ' ', ' '},
			[]byte{' ', ' ', ' ', '#', '$', '@', ' ', '#', ' ', ' ', ' '},
			[]byte{' ', ' ', ' ', '#', ' ', ' ', ' ', '#', ' ', ' ', ' '},
			[]byte{' ', ' ', ' ', '#', '#', '#', '#', '#', ' ', ' ', ' '},
		}
		// 20
		// 29
		// urrdlulluurdldrlddrurruulluld
		puzzle.root.curMap = [][]byte{
			[]byte{' ', ' ', '#', '#', '#', '#', '#', '#', '#', ' ', ' '},
			[]byte{' ', ' ', '#', '#', ' ', ' ', '#', '#', '#', ' ', ' '},
			[]byte{' ', ' ', '#', '#', ' ', '$', '$', ' ', '#', ' ', ' '},
			[]byte{' ', ' ', '#', '#', '.', '.', '.', ' ', '#', ' ', ' '},
			[]byte{' ', ' ', '#', ' ', ' ', '@', '$', ' ', '#', ' ', ' '},
			[]byte{' ', ' ', '#', ' ', ' ', ' ', '#', '#', '#', ' ', ' '},
			[]byte{' ', ' ', '#', '#', '#', '#', '#', '#', '#', ' ', ' '},
		}
		// 21
		// 15
		// dldruuurrddluld
		puzzle.root.curMap = [][]byte{
			[]byte{' ', ' ', ' ', '#', '#', '#', '#', '#', ' ', ' ', ' '},
			[]byte{' ', ' ', ' ', '#', ' ', ' ', ' ', '#', ' ', ' ', ' '},
			[]byte{' ', ' ', '#', '#', ' ', '.', ' ', '#', '#', ' ', ' '},
			[]byte{' ', ' ', '#', ' ', '@', '$', ' ', ' ', '#', ' ', ' '},
			[]byte{' ', ' ', '#', ' ', '$', '$', '$', ' ', '#', ' ', ' '},
			[]byte{' ', ' ', '#', ' ', '.', '.', '.', ' ', '#', ' ', ' '},
			[]byte{' ', ' ', '#', '#', '#', '#', '#', '#', '#', ' ', ' '},
		}
		// 22
		// 50
		// rdldrrddlllluuurldddrrrruullduruuldlldddrrrluurull
		puzzle.root.curMap = [][]byte{
			[]byte{' ', ' ', ' ', ' ', '#', '#', '#', '#', ' ', ' ', ' '},
			[]byte{' ', ' ', '#', '#', '#', '@', ' ', '#', ' ', ' ', ' '},
			[]byte{' ', ' ', '#', '.', ' ', '$', ' ', '#', '#', ' ', ' '},
			[]byte{' ', ' ', '#', ' ', '#', '$', ' ', ' ', '#', ' ', ' '},
			[]byte{' ', ' ', '#', ' ', '#', ' ', '#', ' ', '#', ' ', ' '},
			[]byte{' ', ' ', '#', ' ', ' ', ' ', ' ', '.', '#', ' ', ' '},
			[]byte{' ', ' ', '#', '#', '#', '#', '#', '#', '#', ' ', ' '},
		}
		// 23
		// 33
		// urlddrurruuldrdlldlluurrdrullldrr
		puzzle.root.curMap = [][]byte{
			[]byte{' ', ' ', ' ', ' ', ' ', '#', '#', '#', '#', ' ', ' '},
			[]byte{' ', ' ', ' ', ' ', ' ', '#', ' ', ' ', '#', ' ', ' '},
			[]byte{' ', ' ', '#', '#', '#', '#', '.', ' ', '#', ' ', ' '},
			[]byte{' ', ' ', '#', ' ', ' ', '$', ' ', ' ', '#', ' ', ' '},
			[]byte{' ', ' ', '#', ' ', '@', '$', '.', ' ', '#', ' ', ' '},
			[]byte{' ', ' ', '#', ' ', ' ', ' ', '#', '#', '#', ' ', ' '},
			[]byte{' ', ' ', '#', '#', '#', '#', '#', ' ', ' ', ' ', ' '},
		}
		// 24
		// 21
		// dlllurdrrurulullddldr
		puzzle.root.curMap = [][]byte{
			[]byte{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
			[]byte{' ', ' ', ' ', '#', '#', '#', '#', '#', ' ', ' ', ' '},
			[]byte{' ', ' ', ' ', '#', ' ', ' ', ' ', '#', '#', ' ', ' '},
			[]byte{' ', ' ', '#', '#', '$', '.', '$', ' ', '#', ' ', ' '},
			[]byte{' ', ' ', '#', ' ', '$', '.', '@', ' ', '#', ' ', ' '},
			[]byte{' ', ' ', '#', ' ', ' ', '.', ' ', ' ', '#', ' ', ' '},
			[]byte{' ', ' ', '#', '#', '#', '#', '#', '#', '#', ' ', ' '},
		}
		// 25
		// 17
		// dururddrddlluulur
		puzzle.root.curMap = [][]byte{
			[]byte{' ', ' ', ' ', ' ', '#', '#', '#', '#', ' ', ' ', ' '},
			[]byte{' ', ' ', ' ', '#', '#', ' ', ' ', '#', ' ', ' ', ' '},
			[]byte{' ', ' ', '#', '#', '@', '$', '.', '#', '#', ' ', ' '},
			[]byte{' ', ' ', '#', ' ', '$', '$', ' ', ' ', '#', ' ', ' '},
			[]byte{' ', ' ', '#', ' ', '.', ' ', '.', ' ', '#', ' ', ' '},
			[]byte{' ', ' ', '#', '#', '#', ' ', ' ', ' ', '#', ' ', ' '},
			[]byte{' ', ' ', ' ', ' ', '#', '#', '#', '#', '#', ' ', ' '},
		}
		// 26
		// 35
		// ulldullddrururrdlllddruluruuldrrrdl
		puzzle.root.curMap = [][]byte{
			[]byte{' ', ' ', ' ', '#', '#', '#', '#', ' ', ' ', ' ', ' '},
			[]byte{' ', ' ', '#', '#', ' ', ' ', '#', '#', '#', ' ', ' '},
			[]byte{' ', ' ', '#', ' ', ' ', ' ', ' ', ' ', '#', ' ', ' '},
			[]byte{' ', ' ', '#', '.', '*', '*', '$', '@', '#', ' ', ' '},
			[]byte{' ', ' ', '#', ' ', ' ', ' ', '#', '#', '#', ' ', ' '},
			[]byte{' ', ' ', '#', '#', ' ', ' ', '#', ' ', ' ', ' ', ' '},
			[]byte{' ', ' ', ' ', '#', '#', '#', '#', ' ', ' ', ' ', ' '},
		}
		// 27
		// 34
		// drrrurruullddldllurrruurrdlulduldd
		puzzle.root.curMap = [][]byte{
			[]byte{' ', ' ', ' ', '#', '#', '#', '#', '#', '#', ' ', ' '},
			[]byte{' ', ' ', ' ', '#', ' ', ' ', ' ', ' ', '#', ' ', ' '},
			[]byte{' ', '#', '#', '#', '$', '$', '$', ' ', '#', ' ', ' '},
			[]byte{' ', '#', '@', ' ', '$', '.', '.', ' ', '#', ' ', ' '},
			[]byte{' ', '#', ' ', '$', '.', '.', '.', '#', '#', ' ', ' '},
			[]byte{' ', '#', '#', '#', '#', ' ', ' ', '#', ' ', ' ', ' '},
			[]byte{' ', ' ', ' ', ' ', '#', '#', '#', '#', ' ', ' ', ' '},
		}
		// 28
		// 30
		// drudrrurlluurdldrdlllurlurrurd
		puzzle.root.curMap = [][]byte{
			[]byte{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
			[]byte{' ', ' ', ' ', '#', '#', '#', '#', ' ', ' ', ' ', ' '},
			[]byte{' ', '#', '#', '#', ' ', ' ', '#', '#', '#', '#', ' '},
			[]byte{' ', '#', ' ', ' ', ' ', ' ', ' ', ' ', ' ', '#', ' '},
			[]byte{' ', '#', '@', '$', '*', '*', '*', '.', ' ', '#', ' '},
			[]byte{' ', '#', ' ', ' ', ' ', ' ', ' ', ' ', ' ', '#', ' '},
			[]byte{' ', '#', '#', '#', '#', '#', '#', '#', '#', '#', ' '},
		}
		// 29
		// 37
		// rrrdddllluuddlllluurrrllluurrrdrrldll
		puzzle.root.curMap = [][]byte{
			[]byte{' ', '#', '#', '#', '#', '#', '#', ' ', ' ', ' ', ' '},
			[]byte{' ', '#', ' ', ' ', ' ', ' ', '#', '#', '#', '#', '#'},
			[]byte{' ', '#', ' ', '#', '#', '$', '@', ' ', '.', ' ', '#'},
			[]byte{' ', '#', ' ', '.', ' ', '$', '$', ' ', '#', ' ', '#'},
			[]byte{' ', '#', ' ', '#', '#', '#', '.', '#', '#', ' ', '#'},
			[]byte{' ', '#', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', '#'},
			[]byte{' ', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#'},
		}
		// 30
		// 41
		// dlllurdlllurruuldrddllurdrudrrruuldlrrdll
		puzzle.root.curMap = [][]byte{
			[]byte{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
			[]byte{' ', ' ', '#', '#', '#', '#', ' ', ' ', ' ', ' ', ' '},
			[]byte{' ', ' ', '#', ' ', ' ', '#', '#', '#', '#', ' ', ' '},
			[]byte{' ', '#', '#', '$', '.', '#', ' ', ' ', '#', ' ', ' '},
			[]byte{' ', '#', ' ', '$', '.', '$', ' ', '@', '#', ' ', ' '},
			[]byte{' ', '#', ' ', ' ', '.', ' ', ' ', ' ', '#', ' ', ' '},
			[]byte{' ', '#', '#', '#', '#', '#', '#', '#', '#', ' ', ' '},
		}
	*/
	puzzle.root.curMap = [][]byte{
		[]byte{' ', '#', '#', '#', '#', ' ', ' ', ' ', ' ', ' ', ' '},
		[]byte{' ', '#', ' ', ' ', '#', '#', '#', '#', '#', '#', ' '},
		[]byte{' ', '#', ' ', ' ', ' ', ' ', ' ', ' ', ' ', '#', ' '},
		[]byte{' ', '#', ' ', '.', '#', '$', '$', '@', ' ', '#', ' '},
		[]byte{' ', '#', ' ', ' ', '#', ' ', '#', '#', '#', '#', ' '},
		[]byte{' ', '#', ' ', '.', ' ', ' ', '#', ' ', ' ', ' ', ' '},
		[]byte{' ', '#', '#', '#', '#', '#', '#', ' ', ' ', ' ', ' '},
	}

	puzzle.root.puzzle = &puzzle
	puzzle.size.y = len(puzzle.root.curMap[0])
	puzzle.size.x = len(puzzle.root.curMap)
	puzzle.maps = make(map[string]bool)

	// 初始化
	for i := 0; i < puzzle.size.x; i++ {
		for j := 0; j < puzzle.size.y; j++ {
			if puzzle.root.curMap[i][j] == '@' || puzzle.root.curMap[i][j] == '%' {
				puzzle.root.pos = vector2d{
					x: i,
					y: j,
				}
			} else if puzzle.root.curMap[i][j] == '$' {
				puzzle.root.count++
			}
		}
	}

	puzzle.queue = append(puzzle.queue, &puzzle.root)

	solution := puzzle.search()
	if solution != nil {
		fmt.Println(len(solution.solution))
		fmt.Println(solution.solution)
		fmt.Println(puzzle.root.key())
		fmt.Println(puzzle.memory)
	}
}

func (n *node) move(next vector2d) bool {
	//fmt.Println("before:")
	//fmt.Println(n.key())

	//fmt.Println(next)

	if next.x < 0 || next.x >= n.puzzle.size.x {
		return false
	}
	if next.y < 0 || next.y >= n.puzzle.size.y {
		return false
	}

	switch n.curMap[next.x][next.y] {
	case '#':
		return false
	case ' ', '.':
		if n.curMap[n.pos.x][n.pos.y] == '@' {
			n.curMap[n.pos.x][n.pos.y] = ' '
		} else { // == '%'
			n.curMap[n.pos.x][n.pos.y] = '.'
		}
		if n.curMap[next.x][next.y] == ' ' {
			n.curMap[next.x][next.y] = '@'
		} else {
			n.curMap[next.x][next.y] = '%'
		}
	case '$', '*':
		pos := vector2d{
			x: 2*next.x - n.pos.x,
			y: 2*next.y - n.pos.y,
		}

		if n.curMap[next.x][next.y] == '$' && n.curMap[pos.x][pos.y] == '.' {
			n.count--
		}
		if n.curMap[next.x][next.y] == '*' && n.curMap[pos.x][pos.y] == ' ' {
			n.count++
		}

		switch n.curMap[pos.x][pos.y] {
		case '#', '$', '*':
			return false
		case ' ', '.':
			if n.curMap[pos.x][pos.y] == ' ' {
				n.curMap[pos.x][pos.y] = '$'
			} else { // == '.'
				n.curMap[pos.x][pos.y] = '*'
			}
			if n.curMap[next.x][next.y] == '$' {
				n.curMap[next.x][next.y] = '@'
			} else { // == '*'
				n.curMap[next.x][next.y] = '%'
			}
			if n.curMap[n.pos.x][n.pos.y] == '@' {
				n.curMap[n.pos.x][n.pos.y] = ' '
			} else { // == '%'
				n.curMap[n.pos.x][n.pos.y] = '.'
			}
		default:
			panic(n.key())
		}
	default:
		panic(n.key())
	}

	//fmt.Println("after:")
	//fmt.Println(n.key())

	n.pos = next
	return true
}

func (n *node) duplicate() *node {
	ret := node{
		puzzle:   n.puzzle,
		curMap:   make([][]byte, n.puzzle.size.x),
		solution: n.solution,
		// 剩余箱子数量
		count: n.count,
		// 小人位置
		pos: n.pos,
	}

	for i := range n.curMap {
		ret.curMap[i] = make([]byte, n.puzzle.size.y)
		copy(ret.curMap[i], n.curMap[i])
	}

	return &ret
}

func (n *node) moveUp() {
	next := n.duplicate()
	next.solution += "u"

	if !next.move(vector2d{x: n.pos.x - 1, y: n.pos.y}) {
		return
	}

	b, ok := next.puzzle.maps[next.key()]
	if ok == true && b == true {
		return
	}
	next.puzzle.maps[next.key()] = true

	n.puzzle.queue = append(n.puzzle.queue, next)
}

func (n *node) moveDown() {
	next := n.duplicate()
	next.solution += "d"

	if !next.move(vector2d{x: n.pos.x + 1, y: n.pos.y}) {
		return
	}

	b, ok := next.puzzle.maps[next.key()]
	if ok == true && b == true {
		return
	}
	next.puzzle.maps[next.key()] = true

	n.puzzle.queue = append(n.puzzle.queue, next)
}

func (n *node) moveLeft() {
	next := n.duplicate()
	next.solution += "l"

	if !next.move(vector2d{x: n.pos.x, y: n.pos.y - 1}) {
		return
	}

	b, ok := next.puzzle.maps[next.key()]
	if ok == true && b == true {
		return
	}
	next.puzzle.maps[next.key()] = true

	n.puzzle.queue = append(n.puzzle.queue, next)
}

func (n *node) moveRight() {
	next := n.duplicate()
	next.solution += "r"

	if !next.move(vector2d{x: n.pos.x, y: n.pos.y + 1}) {
		return
	}

	b, ok := next.puzzle.maps[next.key()]
	if ok == true && b == true {
		return
	}
	next.puzzle.maps[next.key()] = true

	n.puzzle.queue = append(n.puzzle.queue, next)
}

func (n *node) key() string {
	var key []byte
	for i := 0; i < n.puzzle.size.x; i++ {
		for j := 0; j < n.puzzle.size.y; j++ {
			key = append(key, n.curMap[i][j])
		}
		key = append(key, '\n')
	}
	return string(key)
}

func (n *node) isSolved() bool {
	if n.count == 0 {
		return true
	}

	return false
}

func (s *sokoban) search() *node {
	for len(s.queue) != 0 {
		//fmt.Println("depth:", depth, ", count:", len(s.queue))

		n := s.queue[0]
		s.queue = s.queue[1:]

		if n.isSolved() {
			return n
		}

		n.moveUp()
		n.moveDown()
		n.moveLeft()
		n.moveRight()

		var m runtime.MemStats
		runtime.ReadMemStats(&m)
		if m.Sys > s.memory {
			s.memory = m.Sys
		}
	}

	return nil
}
