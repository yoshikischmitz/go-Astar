package main

import "fmt"
import "math"

var (
	mapArr [][]int = [][]int{
		{0, 0, 0, 0, 0, 0, 0},
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
	HScore int
	GScore int
	FScore int
	Parent *Node
}

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

// Parses map and returns cartesian coords for start and goal
// Takes a node and calculates its heuristic distance from the goal(or any x y coordinates)
func Hcost(n *Node, goal []int) int {
	// Take the value of
	x := math.Abs(float64(n.X - goal[0]))
	y := math.Abs(float64(n.Y - goal[1]))
	return int(x + y)
}

// Takes two adjacent nodes n1 and n2 and returns the movement cost from n1 to n2
func Gcost(n1 *Node, n2 *Node) int {
	// if they're on the same X or Y axis movement cost is 10 else 14
	if n1.X == n2.X || n1.Y == n2.Y {
		return 10
	} else {
		return 14
	}
}

func Astar(start []int, goal []int) (path [][]int) {
	return path
}

func main() {
	start, goal := ParseMap(mapArr)
	fmt.Println(start, goal)
	path := Astar(start, goal)
	fmt.Println(path)
}
