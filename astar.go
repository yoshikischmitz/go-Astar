package astar

import (
	"math"
	"sort"
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

// An interface for sorting nodes
type ByFscore []Node

// Functions required for sort.Sort
func (ns ByFscore) Len() int           { return len(ns) }
func (ns ByFscore) Swap(i, j int)      { ns[i], ns[j] = ns[j], ns[i] }
func (ns ByFscore) Less(i, j int) bool { return ns[i].Fscore < ns[j].Fscore }

// Wrapper function for sort.Sort
func SortNodes(nodes []Node) {
	sort.Sort(ByFscore(nodes))
}

// Parses map and returns x-y coords for start and goal
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

// Assigns Fscore to Node
// Returns the new fscore
func Fcost(n *Node) int {
	n.Fscore = n.Gscore + n.Hscore
	return n.Fscore
}

// Takes a node and calculates its heuristic distance from the goal
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

// Checks equality of n1 to n2 by their x-y coordinates
func (n1 *Node) Eql(n2 *Node) bool {
	return n1.X == n2.X && n1.Y == n2.Y
}

// Checks if coordinates are within the bounds of the map.
func InMap(n Node, mapArr [][]int) bool {
	if n.X < 0 || n.X > (len(mapArr)-1) || n.Y < 0 || n.Y > (len(mapArr[0])-1) {
		return false
	}
	if mapArr[n.X][n.Y] == 2 {
		return false
	}
	return true
}

// Takes a nodelist and node, checking if the node is already in the list.
func Includes(n *Node, NodeList []Node) bool {
	for _, x := range NodeList {
		if n.Eql(&x) {
			return true
		}
	}
	return false
}

// Returns a path built by following the parent nodes from the current node.
func buildPath(cNode Node) (path [][]int) {
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
	return
}

func NeighborNodes(n *Node) (neighbors []Node) {
	for xS := -1; xS < 2; xS++ {
		x := n.X + xS
		for yS := 1; yS < 2; yS++ {
			y := n.Y + yS
			cNeighbor := Node{X: x, Y: y}
			neighbors = append(neighbors, cNeighbor)
		}
	}
	return
}

// Takes a start and goal as x,y coordinates, as well as the map, and finds the shortest path
// to goal from start using the A* pathfinding algorithm. Returns a two dimensional array of x,y
// coordinates plotting the path from start to goal if it succeeds, else it returns an empty array
// and an error.
func Astar(start []int, goal []int, mapArr [][]int) ([][]int, error) {
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
			return buildPath(cNode), nil
		}
		openSet = openSet[1:]
		closedSet = append(closedSet, cNode)
		// xS is for xShift
		for _, neighbor := range NeighborNodes(&cNode) {
			if InMap(neighbor, mapArr) == false {
				continue
			}
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
	return [][]int{}, nil
}
