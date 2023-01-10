package main

import (
	// f "patterns/creational/factory_method"
	"fmt"
	af "patterns/creational/abstract_factory"
	g "patterns/creational/factory_method"
	"reflect"
)

func main() {

	a, _ := af.GetFurnitureCreator("ArtDeco")
	fmt.Println(a)
	aChair := a.CreateChair()

	fmt.Println(reflect.TypeOf(aChair).Name())
	fmt.Println("----")
	fmt.Println(reflect.TypeOf(aChair).Elem())
	fmt.Println(reflect.TypeOf(&aChair).Elem())

	gpuMaker, _ := g.GetGpuAssembler("msi")

	msiGpu := gpuMaker.CreateGpu()

	msiGpu.Assemble(100, 78)
	fmt.Println(msiGpu.CheckAssembly())

}
