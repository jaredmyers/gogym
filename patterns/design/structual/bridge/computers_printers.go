// Bridge Pattern Example
package main

import "log"

// abstract computer
type Computer interface {
	Print()
	SetPrinter(Printer)
}

// abstract printer
type Printer interface {
	PrintFile()
}

// concrete computer 1
type Mac struct {
	printer Printer
}

func (m *Mac) Print() {
	log.Println("Print request for mac")
	m.printer.PrintFile()
}

func (m *Mac) SetPrinter(p Printer) {
	m.printer = p
}

// concrete computer 2
type Windows struct {
	printer Printer
}

func (w *Windows) Print() {
	log.Println("Print request for windows")
	w.printer.PrintFile()
}

func (w *Windows) SetPrinter(p Printer) {
	w.printer = p
}

// Concrete Printers

type Hp struct{}

func (p *Hp) PrintFile() {
	log.Println("Printing by an HP printer")
}

type Epson struct{}

func (p *Epson) PrintFile() {
	log.Println("Printing by Epson printer")
}

func main() {

	hpPrinter := &Hp{}
	epsonPrinter := &Epson{}

	mac := &Mac{}
	mac.SetPrinter(epsonPrinter)
	mac.Print()

	win := &Windows{}
	win.SetPrinter(hpPrinter)
	win.Print()

	win.SetPrinter(epsonPrinter)
	win.Print()
}
