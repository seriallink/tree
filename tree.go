package tree

import (
	"sort"
	"strings"
)

type Tree map[string]Tree

type WalkFunc func(path string, err error) error

// New returns a new tree with the given root.
func New(root string) Tree {
	tree := Tree{}
	tree.Add(root)
	return tree
}

// Add adds the given path to the tree.
func (t Tree) Add(path string) {
	t.add(strings.Split(path, "/"))
}

// Walk walks the tree and calls the given function for each item.
func (t Tree) Walk(fn WalkFunc) {
	t.walk(fn)
}

// Item returns the key of current node.
func (t Tree) Item() string {
	for key := range t {
		return key
	}
	return ""
}

// Items returns a list of all items in the tree.
// For example, if the tree contains the following paths: root/sibling1/child1, root/sibling2/child1, root/sibling2/child2
// the returned list will be: [root, sibling1, child1, sibling2, child1, child2]
func (t Tree) Items() (list []string) {
	t.Walk(func(path string, err error) error {
		list = append(list, path)
		return nil
	})
	return
}

// SortedKeys returns a list of keys in the current node sorted in alphabetical order.
func (t Tree) SortedKeys() (list []string) {
	for k := range t {
		list = append(list, k)
	}
	sort.Strings(list)
	return
}

// AllPaths returns all possible paths for the tree.
// For example, if the tree contains the following paths: root/sibling1/childa, root/sibling2/childb, root/sibling2/childc
// the returned list will be: [root, root/sibling1, root/sibling2, root/sibling1/childa, root/sibling2/childb, root/sibling2/childc]
func (t Tree) AllPaths() (list []string) {
	t.allPaths("/", "", &list)
	sort.Strings(list)
	return
}

// FullPaths returns a list of full paths in the tree.
// For example, if the tree contains the following paths: root/sibling1/childa, root/sibling2/childb, root/sibling2/childc
// the returned list will be: [root/sibling1/childa, root/sibling2/childb, root/sibling2/childc]
func (t Tree) FullPaths() (list []string) {
	t.fullPaths("/", "", &list)
	sort.Strings(list)
	return
}

func (t Tree) add(items []string) {
	if len(items) == 0 {
		return
	}
	next, ok := t[items[0]]
	if !ok {
		next = Tree{}
		t[items[0]] = next
	}
	next.add(items[1:])
}

func (t Tree) walk(fn WalkFunc) {
	for k, v := range t {
		if err := fn(k, nil); err != nil {
			return
		}
		v.walk(fn)
	}
}

func (t Tree) allPaths(separator, prefix string, list *[]string) {
	for k, v := range t {
		*list = append(*list, prefix+k)
		v.allPaths(separator, prefix+k+separator, list)
	}
}

func (t Tree) fullPaths(separator, prefix string, list *[]string) {
	for k, v := range t {
		if len(v) == 0 {
			*list = append(*list, prefix+k)
		} else {
			v.fullPaths(separator, prefix+k+separator, list)
		}
	}
}
