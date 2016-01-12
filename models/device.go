package models

//type DeviceType int
//
//const (
//	BOX DeviceType = iota
//	VENDOR
//	PARKING_LOCK
//)

type Device struct {
	*Site
	//	DeviceType
	ID       int
	Host     *Device
	Name     string
	Serial   string
	debit    int
	credit   int
	balance  int
	IsHost   bool
	IsOnline bool
}
type Box struct {
	Device
	Cash       int
	IsLockOpen bool //ตู้เซฟเก็บเงินเปิดอยู่หรือไม่
}

type Vendor struct {
	Device
}

//func LoadDevice(DeviceType) []Device {
//	var d []Device
//	return d
//}

func NewDevice(name, serial string) *Device {
	d := new(Device)
	d.Name = name
	//	d.DeviceType = deviceType
	d.Serial = serial
	d.debit = 0
	d.credit = 0
	d.balance = 0
	d.IsHost = false
	d.IsOnline = false
	return d
}

func (d *Device) Debit() int {
	return d.debit
}

func (d *Device) SetDebit(value int) {
	d.debit += value
	d.balance += value
}

func (d *Device) Credit() int {
	return d.credit
}

func (d *Device) SetCredit(value int) {
	d.credit -= value
	d.balance -= value
}

func (d *Device) Balance() int {
	return d.balance
}

func (d *Device) SetBalance(value int) {
	d.balance = value
}
