package astar

import (
	"math/rand"
	"reflect"
	"testing"
)

var (
	testMap [][]int = [][]int{
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 2, 0, 0, 0},
		{0, 1, 0, 2, 0, 3, 0},
		{0, 0, 0, 2, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
	}
)

func EqualPosition(xy1 []int, xy2 []int) bool {
	if xy1[0] != xy2[0] || xy1[1] != xy2[1] {
		return false
	}
	return true
}

func TestMap(t *testing.T) {
	start, goal := ParseMap(testMap)
	expected_start := []int{2, 1}
	expected_goal := []int{2, 5}
	if EqualPosition(start, expected_start) == false {
		t.Errorf("start is %d, should be %d", start, expected_start)
	}
	if EqualPosition(goal, expected_goal) == false {
		t.Errorf("goal is %d, should be %d", goal, expected_goal)
	}
}

func TestGcost(t *testing.T) {
	// Create two nodes at the same X position, should return 10
	Node1 := Node{0, 2, 0, 0, 0, nil}
	Node2 := Node{0, 1, 0, 0, 0, nil}
	if Gcost(&Node1, &Node2) != 10 {
		t.Errorf("G cost for nodes on the same X-Axis should be 10")
	}
	// See if the side effect applied properly
	if Node1.Gscore != 10 {
		t.Errorf("struct Node's Gscore not updating")
	}
	// Create two nodes at the same Y positions, should return 10
	Node1 = Node{2, 0, 0, 0, 0, nil}
	Node2 = Node{1, 0, 0, 0, 0, nil}
	if Gcost(&Node1, &Node2) != 10 {
		t.Errorf("G cost for nodes on the same Y-Axis should be 10")
	}
	// Create two nodes at different X and Y positions, should return 14
	Node1 = Node{0, 2, 0, 0, 0, nil}
	Node2 = Node{1, 0, 0, 0, 0, nil}
	if Gcost(&Node1, &Node2) != 14 {
		t.Errorf("G cost for nodes with different X and Y values should be 14")
	}
}

func TestHcost(t *testing.T) {
	// Create Node at x 0 y 2
	Node := Node{0, 2, 0, 0, 0, nil}
	// Create goal at x 5 y 4
	goal := []int{5, 4}

	//Test for when Node and goal are separated by a distance of 7
	if Hcost(&Node, goal) != 7 {
		t.Errorf("distance to goal should be 7")
	}
	// Check side effect
	if Node.Hscore != 7 {
		t.Errorf("struct Node's Gscore not updating")
	}
	// Test for when Node == Goal
	Node.X = 5
	Node.Y = 4
	if Hcost(&Node, goal) != 0 {
		t.Errorf("distance to goal should be 0")
	}

	// Test at negatives
	Node.X = -1
	Node.Y = -2
	if Hcost(&Node, goal) != 12 {
		t.Errorf("distance to goal should be 12")
	}
}

// Returns a slice of nodes with randomly generated fscores
func GenerateRandomNodes(size int) (nodes []Node) {
	for i := 0; i < size; i++ {
		newNode := Node{Fscore: rand.Int()}
		nodes = append(nodes, newNode)
	}
	return
}

// Takes a list of nodes and checks if they are sorted by F Score.
func IsSorted(nodeList []Node) bool {
	for i, n := range nodeList {
		// If our index is at the edge of the nodelist, break.
		if i == len(nodeList)-1 {
			break
		}
		// Check that the current element's Fscore is less than the next.
		if n.Fscore > nodeList[i+1].Fscore {
			return false
		}
	}
	return true
}

// Generate a random set of nodes and see if they sort properly
func TestSort(t *testing.T) {
	nodeList := GenerateRandomNodes(40)
	SortNodes(nodeList)
	if IsSorted(nodeList) == false {
		t.Errorf("Node list not sorted")
	}

	// Check if a pre-sorted list stays sorted
	SortNodes(nodeList)

	if IsSorted(nodeList) == false {
		t.Errorf("pre-sorted list returning unsorted")
	}
}

// Helper function for TestAstar, to reduce code duplication
func AstarWorking(mapArr [][]int, expected_path [][]int) ([][]int, bool) {
	start, goal := ParseMap(mapArr)
	path, _ := Astar(start, goal, mapArr)
	if reflect.DeepEqual(path, expected_path) == false {
		return path, false
	}
	return path, true
}

func TestAstar(t *testing.T) {
	mapArr := [][]int{
		{0, 0, 0, 2, 0, 0, 0},
		{0, 0, 0, 2, 0, 0, 0},
		{0, 1, 0, 2, 0, 3, 0},
		{0, 0, 0, 2, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
	}
	expected_path := [][]int{
		{2, 5},
		{3, 4},
		{4, 3},
		{3, 2},
		{2, 1},
	}
	if path, working := AstarWorking(mapArr, expected_path); working == false {
		t.Errorf("Astar returned path %d, should be %d", path, expected_path)
	}
}
