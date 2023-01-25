package prototype

import "fmt"

// -- Prototype Interface -- //
type Inode interface {
	Print(string)
	Clone() Inode
}

// -- Concrete Prototype -- //
type file struct {
	name string
}

func CreateFile(name string) Inode {
	return &file{name: name}
}

func (f *file) Print(indentation string) {
	fmt.Println(indentation + f.name)
}

func (f *file) Clone() Inode {
	return &file{name: f.name + "_clone"}
}

// -- Concrete Prototype -- //
type folder struct {
	children []Inode
	name     string
}

func CreateFolder(name string, children []Inode) Inode {
	return &folder{name: name, children: children}

}

func (f *folder) Print(indentation string) {
	fmt.Println(indentation + f.name)
	for _, i := range f.children {
		i.Print(indentation + indentation)
	}
}

func (f *folder) Clone() Inode {
	cloneFolder := &folder{name: f.name + "_clone"}
	var tempChildren []Inode
	for _, i := range f.children {
		copy := i.Clone()
		tempChildren = append(tempChildren, copy)
	}
	cloneFolder.children = tempChildren
	return cloneFolder
}
