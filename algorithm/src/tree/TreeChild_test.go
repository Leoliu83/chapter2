package tree

import (
	"testing"
)

func TestTreeChildInitTree(t *testing.T) {
	tree := CInitTree(10)
	t.Log(tree)
}

func TestTreeChildCDestoryTree(t *testing.T) {
	tree := CInitTree(10)
	t.Logf("%v,%p,%+v", &tree, tree, tree)
	tree.Destory()
	t.Logf("%v,%p,%+v", &tree, tree, tree)
}

func TestTreeChildCleanFlag(t *testing.T) {
	tree := CInitTree(8)
	tree.debug = true
	tree.flag = 1<<8 - 1
	tree.CleanFlag(4)
}

func TestTreeChildCInsertChild(t *testing.T) {
	tree := CInitTree(8)
	tree.debug = true
	c := CCreateCBox(0)
	tree.InsertRoot(c)

	c1 := CCreateCBox(1)
	tree.InsertChild(&tree.nodes[0], c1, 0)
	tree.Print()

	c2 := CCreateCBox(2)
	tree.InsertChild(&tree.nodes[0], c2, 2)
	tree.Print()
	// c3 := CCreateCBox(3)
	// tree.InsertChild(&tree.nodes[0], c3, 1)
	// tree.Print()

	// c4 := CCreateCBox(4)
	// c5 := CCreateCBox(5)
	// c6 := CCreateCBox(6)
	// c7 := CCreateCBox(7)

	// tree.InsertChild(&tree.nodes[3], c4, 0)
	// tree.InsertChild(&tree.nodes[3], c5, 0)
	// tree.InsertChild(&tree.nodes[4], c6, 1)
	// tree.InsertChild(&tree.nodes[6], c7, 0)
	// tree.Print()

	// c8 := CCreateCBox(8)
	// tree.InsertChild(&tree.nodes[4], c8, 0)
}
