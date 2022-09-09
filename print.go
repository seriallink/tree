package tree

import (
	"fmt"
	"io"
)

type BoxType int

const (
	Regular BoxType = iota
	Last
	AfterLast
	Between
)

func (t Tree) Fprint(w io.Writer, root bool, padding string) {
	if t == nil {
		return
	}
	index := 0
	for k, v := range t {
		fmt.Fprintf(w, "%s%s\n", padding+getPadding(root, getBoxType(index, len(t))), k)
		v.Fprint(w, false, padding+getPadding(root, getBoxTypeExternal(index, len(t))))
		index++
	}
}

// FprintSorted prints in alphabetical order
// For example, if the tree contains the following paths: root/sibling1/childa, root/sibling2/childb, root/sibling2/childc
// the returned list will be:
// root
// ├── sibling1
// │   └── childa
// └── sibling2
//
//	├── childb
//	└── childc
func (t Tree) FprintSorted(w io.Writer, root bool, padding string) {
	if t == nil {
		return
	}
	index := 0
	for _, k := range t.SortedKeys() {
		v := t[k]
		fmt.Fprintf(w, "%s%s\n", padding+getPadding(root, getBoxType(index, len(t))), k)
		v.Fprint(w, false, padding+getPadding(root, getBoxTypeExternal(index, len(t))))
		index++
	}
}

func (boxType BoxType) String() string {
	switch boxType {
	case Regular:
		return "\u251c" // ├
	case Last:
		return "\u2514" // └
	case AfterLast:
		return " "
	case Between:
		return "\u2502" // │
	default:
		panic("invalid box type")
	}
}

func getBoxType(index int, len int) BoxType {
	if index+1 == len {
		return Last
	} else if index+1 > len {
		return AfterLast
	}
	return Regular
}

func getBoxTypeExternal(index int, len int) BoxType {
	if index+1 == len {
		return AfterLast
	}
	return Between
}

func getPadding(root bool, boxType BoxType) string {
	if root {
		return ""
	}

	return boxType.String() + " "
}
