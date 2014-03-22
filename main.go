package main

import (
	"fmt"
	"math"
	"sort"
)

var (
	mapArr [][]int = [][]int{
		{0, 0, 0, 2, 0, 0, 0},
		{0, 0, 0, 2, 0, 0, 0},
		{0, 1, 0, 2, 0, 3, 0},
		{0, 0, 0, 2, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
	}
)

// In this program it is understood that GScore refers to the value returned by the function
// GCost. The same applies to HScore and FScore.
type Node struct {
	X      int
	Y      int
	Gscore int
	Hscore int
	Fscore int
	Parent *Node
}

type ByFscore []Node

func (ns ByFscore) Len() int           { return len(ns) }
func (ns ByFscore) Swap(i, j int)      { ns[i], ns[j] = ns[j], ns[i] }
func (ns ByFscore) Less(i, j int) bool { return ns[i].Fscore < ns[j].Fscore }

// Parses map and returns cartesian coords for start and goal
func ParseMap(Map [][]int) (start []int, goal []int) {
	start_val := 1
	end_val := 3
	for xIdx, xArr := range Map {
		for yIdx, ySqr := range xArr {
			if ySqr == start_val {
				start = []int{xIdx, yIdx}
			}
			if ySqr == end_val {
				goal = []int{xIdx, yIdx}
			}
		}
	}
	return
}

// Assigns the Nodes Fcost
// Returns the new fscore
func Fcost(n *Node) int {
	n.Fscore = n.Gscore + n.Hscore
	return n.Fscore
}

// Takes a node and calculates its heuristic distance from the goal(or any x y coordinates)
// Assigns the h_score to the node as a side effect
func Hcost(n *Node, goal []int) int {
	// Take the value of
	x := math.Abs(float64(n.X - goal[0]))
	y := math.Abs(float64(n.Y - goal[1]))
	hscore := int(x + y)
	n.Hscore = hscore
	return hscore
}

// Takes two adjacent nodes n1 and n2 and returns the movement cost from n1 to n2
// Returns the g_score as an int and assigns it to the node as a side-effect.
func Gcost(n1 *Node, n2 *Node) int {
	// if they're on the same X or Y axis movement cost is 10 else 14
	if n1.X == n2.X || n1.Y == n2.Y {
		n1.Gscore = 10
		return 10
	} else {
		n1.Gscore = 14
		return 14
	}
}

// Compares one node to another
func (n1 *Node) Eql(n2 *Node) bool {
	return n1.X == n2.X && n1.Y == n2.Y
}

func SortNodes(nodes []Node) {
	sort.Sort(ByFscore(nodes))
}

func InMap(x int, y int, mapArr [][]int) bool {
	if x < 0 || x > (len(mapArr)-1) || y < 0 || y > (len(mapArr[0])-1) {
		return false
	}
	if mapArr[x][y] == 2 {
		return false
	}
	return true
}

func Includes(n *Node, NodeList []Node) bool {
	for _, x := range NodeList {
		if n.Eql(&x) {
			return true
		}
	}
	return false
}

func Astar(start []int, goal []int) (path [][]int) {
	startNode := Node{
		X:      start[0],
		Y:      start[1],
		Gscore: 0,
	}
	goalNode := Node{
		X: goal[0],
		Y: goal[1],
	}
	closedSet := []Node{}
	openSet := []Node{startNode}
	Hcost(&startNode, goal)
	Fcost(&startNode)
	// This is where the real action starts
	for len(openSet) != 0 {
		SortNodes(openSet)
		cNode := openSet[0]
		if cNode.Eql(&goalNode) {
			at_start := false
			pNode := cNode
			for at_start == false {
				path = append(path, []int{pNode.X, pNode.Y})
				pNode = *pNode.Parent
				if pNode.Parent == nil {
					path = append(path, []int{pNode.X, pNode.Y})
					break
				}
			}
			return path
		}
		openSet = openSet[1:]
		closedSet = append(closedSet, cNode)
		// xS is for xShift
		for xS := -1; xS < 2; xS++ {
			x := cNode.X + xS
			for yS := 1; yS < 2; yS++ {
				y := cNode.Y + yS
				if InMap(x, y, mapArr) == false {
					continue
				}
				neighbor := Node{X: x, Y: y}
				if Includes(&neighbor, closedSet) {
					continue
				}
				tentativeGscore := Gcost(&cNode, &neighbor) + cNode.Gscore
				if Includes(&neighbor, openSet) == false || tentativeGscore < neighbor.Gscore {
					neighbor.Parent = &cNode

					neighbor.Gscore = tentativeGscore
					Hcost(&neighbor, goal)
					Fcost(&neighbor)
					openSet = append(openSet, neighbor)
				}
			}
		}
	}
	return path
}

func main() {
	start, goal := ParseMap(mapArr)
	fmt.Println(Astar(start, goal))
}
