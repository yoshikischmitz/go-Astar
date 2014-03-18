package main

import "testing"

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
