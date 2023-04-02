// Adapter Pattern Conceptual Example
package main

import "log"

type Client struct{}

func (c *Client) InsertLightningConnector(com Computer) {
	log.Println("Client inserts lightening into computer")
	com.InsertIntoLightningPort()
}

type Computer interface {
	InsertIntoLightningPort()
}

type Mac struct{}

func (m *Mac) InsertIntoLightningPort() {
	log.Println("USB connector is plugged into mac machine")
}

type Windows struct{}

func (w *Windows) InsertIntoUSBPort() {
	log.Println("USB connector is plugged into windows machine")
}

type WindowsAdapter struct {
	windowsMachine *Windows
}

func (w *WindowsAdapter) InsertIntoLightningPort() {
	log.Println("Adapter converts lightning signal to USB")
	w.windowsMachine.InsertIntoUSBPort()
}

func notmain() {

	client := &Client{}
	mac := &Mac{}

	client.InsertLightningConnector(mac)

	windowsMachine := &Windows{}
	windowsMachineAdapter := &WindowsAdapter{
		windowsMachine: windowsMachine,
	}

	client.InsertLightningConnector(windowsMachineAdapter)
}
