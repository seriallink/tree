package tree

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyTree(t *testing.T) {
	tree := Tree{}
	tree.Fprint(os.Stdout, true, "")
}

func TestOneLevel(t *testing.T) {
	tree := New("root")
	tree.Fprint(os.Stdout, true, "")
}

func TestTwoLevels(t *testing.T) {
	tree := Tree{}
	tree.Add("root/sibling")
	tree.Fprint(os.Stdout, true, "")
}

func TestTwoSiblings(t *testing.T) {
	tree := Tree{}
	tree.Add("root/sibling1")
	tree.Add("root/sibling2")
	tree.Fprint(os.Stdout, true, "")
}

func TestManyLevels(t *testing.T) {
	tree := Tree{}
	tree.Add("root")
	tree.Add("root/sibling0")
	tree.Add("root/sibling1")
	tree.Add("root/sibling2")
	tree.Add("root/sibling1/childa")
	tree.Add("root/sibling2/childb")
	tree.Add("root/sibling2/childc")
	tree.Add("root/sibling2/childc/deeper")
	tree.FprintSorted(os.Stdout, false, "")
}

func TestTreeWalk(t *testing.T) {
	tree := Tree{}
	var keys []string
	tree.Add("root")
	tree.Add("root/sibling1")
	tree.Add("root/sibling2")
	tree.Add("root/sibling1/childa")
	tree.Add("root/sibling2/childb")
	tree.Add("root/sibling2/childc")
	tree.Walk(func(path string, err error) error {
		keys = append(keys, path)
		return nil
	})
	fmt.Println(keys)
	assert.Len(t, keys, 6, "should have 6 keys")
}

func TestTreeAllPaths(t *testing.T) {
	tree := Tree{}
	tree.Add("root")
	tree.Add("root/sibling0")
	tree.Add("root/sibling1")
	tree.Add("root/sibling2")
	tree.Add("root/sibling1/childa")
	tree.Add("root/sibling2/childb")
	tree.Add("root/sibling2/childc")
	assert.Equal(t, tree.AllPaths(), []string{
		"root",
		"root/sibling0",
		"root/sibling1",
		"root/sibling1/childa",
		"root/sibling2",
		"root/sibling2/childb",
		"root/sibling2/childc"})
}

func TestTreeItems(t *testing.T) {
	tree := Tree{}
	tree.Add("root")
	tree.Add("root/sibling1")
	tree.Add("root/sibling2")
	tree.Add("root/sibling1/childa")
	tree.Add("root/sibling2/childb")
	tree.Add("root/sibling2/childc")
	fmt.Println(tree.Items())
	assert.Equal(t, tree.Items(), []string{
		"root",
		"sibling1",
		"childa",
		"sibling2",
		"childb",
		"childc"})
}
