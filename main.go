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
	Gscore int
	Hscore int
	Fscore int
	Parent *Node
}

type ByFscore []Node


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
func Fcost(n *Node){
	n.Fscore = n.Gscore + n.Hscore
}

// Parses map and returns cartesian coords for start and goal
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
func (n1 *Node) Eql (n2 *Node) bool{
	return n1.X == n2.X && n1.Y == n2.Y
}

func SortNodes(*[]Node) {
	
}
func Astar(start []int, goal []int) (path [][]int) {
	startNode := Node{
		X: start[0],
		Y: start[1],
		Gscore: 0,
	}
	/* goalNode  := Node{
		X: goal[0],
		Y: goal[1],
	} */
	//closedset := []Node{}
	openset	  := []Node{startNode}
	Fcost(&startNode)
	Hcost(&startNode, goal)

	// This is where the real action starts
	for len(openset) != 0 {

	}
	return path
}

func main() {
	start, goal := ParseMap(mapArr)
	fmt.Println(start, goal)
	path := Astar(start, goal)
	fmt.Println(path)
}
