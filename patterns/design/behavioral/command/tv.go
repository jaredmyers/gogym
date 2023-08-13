package main

import "fmt"

type Command interface {
	execute()
}

type Button struct {
	command Command
}

func (b *Button) press() {
	b.command.execute()
}

type OnCommand struct {
	device Device
}

func (c *OnCommand) execute() {
	c.device.on()
}

type OffCommand struct {
	device Device
}

func (c *OffCommand) execute() {
	c.device.off()
}

type ChannelUpCommand struct {
	device Device
}

func (c *ChannelUpCommand) execute() {
	c.device.changeChannelUp()
}

type ChannelDownCommand struct {
	device Device
}

func (c *ChannelDownCommand) execute() {
	c.device.changeChannelDown()
}

type Device interface {
	on()
	off()
	changeChannelUp()
	changeChannelDown()
}

type Tv struct {
	isRunning bool
	channel   int
}

func (t *Tv) on() {
	t.isRunning = true
	fmt.Println("turning tv on")
}

func (t *Tv) off() {
	t.isRunning = false
	fmt.Println("turning tv off")
}

func (t *Tv) changeChannelUp() {
	t.channel++
	fmt.Println(t.channel)
}
func (t *Tv) changeChannelDown() {
	t.channel--
	fmt.Println(t.channel)
}

func main() {
	tv := &Tv{}

	onCommand := &OnCommand{
		device: tv,
	}

	offCommand := &OffCommand{
		device: tv,
	}

	onButton := &Button{
		command: onCommand,
	}

	offButton := &Button{
		command: offCommand,
	}

	upCommand := &ChannelUpCommand{
		device: tv,
	}

	upButton := &Button{
		command: upCommand,
	}

	downCommand := &ChannelDownCommand{
		device: tv,
	}

	downButton := &Button{
		command: downCommand,
	}

	onButton.press()
	offButton.press()
	upButton.press()
	upButton.press()
	upButton.press()

	downButton.press()
	downButton.press()
	downButton.press()

	fmt.Println("------")

	tv.changeChannelUp()
	tv.changeChannelUp()
	tv.changeChannelUp()

	tv.changeChannelDown()
	tv.changeChannelDown()
	tv.changeChannelDown()
}
