package main

type DeviceType int

const (
	BOX DeviceType = iota
	VENDOR
	PARKING_LOCK
)

type Device struct {
	*Site
	DeviceType
	ID       int
	Host     *Device
	Name     string
	Serial   string
	debit    int
	credit   int
	balance  int
	isHost   bool
	isOnline bool
}

func NewDevice(name, deviceType, serial string) *Device {
	d := new(Device)
	d.Name = name
	d.DeviceType = deviceType
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
