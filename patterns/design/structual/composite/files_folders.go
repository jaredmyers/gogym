// Composite pattern example
package main

import "log"

type Component interface {
	search(string)
}

type File struct {
	name string
}

func (f *File) search(keyword string) {
	log.Printf("Searching for keyword %s in file %s\n", keyword, f.name)
}

func (f *File) getName() string {
	return f.name
}

type Folder struct {
	components []Component
	name       string
}

func (f *Folder) search(keyword string) {
	log.Printf("Searching recursively for keyword %s in folder %s.\n", keyword, f.name)
	for _, composite := range f.components {
		composite.search(keyword)
	}
}

func (f *Folder) add(c Component) {
	f.components = append(f.components, c)
}

func main() {

	file1 := &File{name: "file1"}
	file2 := &File{name: "file2"}
	file3 := &File{name: "file3"}

	folder1 := &Folder{
		name: "Folder1",
	}
	folder1.add(file1)

	folder2 := &Folder{
		name: "Folder2",
	}

	folder2.add(file2)
	folder2.add(file3)

	folder2.add(folder1)

	folder2.search("something")
}
