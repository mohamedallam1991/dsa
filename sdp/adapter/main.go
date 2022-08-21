package main

import (
	"fmt"
)

type Computer interface {
	connect()
}

type mac struct{}

func (m *mac) connect() {
	fmt.Println("connecting")
}

type windows struct{}

type adapter struct {
	adapter *windows
}

func (a adapter) connect() {
	a.adapter.windowConnection()
}

func (m *windows) windowConnection() {
	fmt.Println("connecting")
}

type client struct{}

func (c client) connectToComputer(com Computer) {
	com.connect()
}

type connection struct{}

func (c connection) connectFromoTo(com Connector) {
}

type Connector interface {
	connect()
}

type oldCon struct{}

func (o oldCon) connect() {
	fmt.Println("connected")
}

type newConnector struct {
	newCon oldCon
}

func (n newConnector) connect() {
	n.newCon.connect()
}

func (c newConnector) usingNewerPlug() {
	fmt.Println("coonnecting using new plug")
}

type LegacyPrinter interface {
	Print(s string) string
}
type MyLegacyPrinter struct{}

func (l *MyLegacyPrinter) Print(s string) (newMSG string) {
	newMSG = fmt.Sprintf("Legacy printer: %s\n", s)
	fmt.Println(newMSG)
	return
	// fmt.Println("Legacy printer, printed out the data")
	// return s
}

type NewPrint interface {
	PrinterStored() string
}
type PrinterAdater struct {
	OldPinter *MyLegacyPrinter
	Msg       string
}

func (p *PrinterAdater) PrinterStored() (newMSG string) {
	// log.Fatal(p)
	// a := MyLegacyPrinter{}
	if p.OldPinter != nil {
		newMSG := fmt.Sprintf("Adapter: %s\n", p.Msg)
		newMSG = p.OldPinter.Print(newMSG)
		// return newMSG
	} else {
		newMSG = p.OldPinter.Print(p.Msg)
		newMSG = p.Msg
	}
	return
}

// func testing() {

// 	client := &client{}
// 	mac := &mac{}
// 	client.connectToComputer(mac)
// 	windowsMachine := &windows{}

// 	windowsMachineAdapter := &adapter{
// 		adapter: windowsMachine,
// 	}
// 	client.windowsMachineAdapter(windowsMachineAdapter)
// }
func main() {
	// testing()
	// log.Fatalln()
	// mac := &mac{}
	// client.insertSquareUsbInComputer(&mac{})
	// windowsMachine := &windows{}
	// windowsMachineAdapter := &windowsAdapter{
	// 	windowMachine: &windows{},
	// }
	// client.insertSquareUsbInComputer(&windowsAdapter{
	// 	windowMachine: &windows{},
	// })

	msg := "Printing"
	printing := PrinterAdater{OldPinter: &MyLegacyPrinter{}, Msg: msg}
	printing.PrinterStored()
	printings := PrinterAdater{OldPinter: nil, Msg: msg}
	printings.PrinterStored()
	// var conn Connector = newConnector{}
	// conn.connect()
}

// type IProcess interface {
// 	process()
// }
// type Adapter struct {
// 	adaptee Adaptee
// }

// type Adaptee struct {
// 	adapterType int
// }

// func (adapter Adapter) process() {
// 	fmt.Println("process")

// 	adapter.adaptee.convert()
// }

// func (adaptee Adaptee) convert() {
// 	fmt.Println("convert")
// }

// type Computer interface {
// 	insertUSBA()
// }

// type USBA struct {
// 	USBA int
// }

// type USBC struct {
// 	adapter USBA
// }

// func (adapter USBC) insertUSBC() {
// 	fmt.Println("insertUSBC")
// 	// adapter.adapter./
// 	adapter.adapter.insertUSBA()
// }

// func (adaptee USBA) insertUSBA() {
// 	fmt.Println("adapte USBC to USBA")
// }

// // process from algo book
// // printer from the DP book
// // Windows and mac
// // old new phone
// func main() {
// 	testing()

// 	log.Fatal("")
// 	var processor IProcess = Adapter{}
// 	processor.process()

// 	var phone Computer = USBA{}
// 	phone.insertUSBA()
// 	// old := &MyLegacyPrinter{}

// 	var printing Printer = MyLegacyPrinter{}
// 	printing.PaperPrint()
// 	// l := MyLegacyPrinter{}
// 	// a := PrinterAdapter{OldPrinter: l}
// 	var newPrint Printer = PrinterAdapter{}
// 	newPrint.PaperPrint()
// 	// fmt.Println(a)
// }

// type Printer interface {
// 	PaperPrint()
// }

// func (l MyLegacyPrinter) PaperPrint() {
// 	fmt.Println("printing")
// }

// type MyLegacyPrinter struct {
// }

// type PrinterAdapter struct {
// 	OldPrinter MyLegacyPrinter
// }

// func (p PrinterAdapter) PaperPrint() {
// 	p.OldPrinter.PaperPrint()
// }

// func testing() {
// 	var request ITarget = Adapteer{}
// 	request.request()
// }

// type ITarget interface {
// 	request()
// }
// type Adapteee struct{}

// func (adapteee Adapteee) specificRequest() {
// 	fmt.Println("specificRRequest")
// }

// type Adapteer struct {
// 	Adapteee Adapteee
// }

// func (adapteer Adapteer) request() {
// 	adapteer.Adapteee.specificRequest()
// }
