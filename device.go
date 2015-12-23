package main

type Device struct {
	*Site
	Host     *Device
	Name     string
	Group    string
	Serial   string
	debit    int
	credit   int
	balance  int
	isHost   bool
	isOnline bool
}

func NewDevice(name, group, serial string) *Device {
	d := new(Device)
	d.Name = name
	d.Group = group
	d.Serial = serial
	d.debit = 0
	d.credit = 0
	d.balance = 0
	d.isHost = false
	d.isOnline = false
	return d
}

func (d *Device) Debit(value int) {
	d.debit = value
	d.balance = d.balance + d.debit
}

func (d *Device) Credit(value int) {
	d.credit = value
	d.balance = d.balance - d.credit
}
