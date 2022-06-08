package main

import "fmt"

//练习2：用多态来实现 将 移动硬盘或者U盘或者MP3插到电脑上进行读写数据（分析类，接口，方法）
//定义接口

// USBer 1.定义接口
type USBer interface {
	Read()
	Write()
}

// USBDev 2.创建对象
type USBDev struct {
	id     int
	name   string
	rspeed int
	wspeed int
}

type Mobile struct {
	USBDev
	call string
}
type UDisk struct {
	USBDev
}

//3.方法实现
func (m *Mobile) Read() {
	fmt.Printf("正在读取%s手机数据，速度为：%d\n", m.name, m.rspeed)
}
func (m *Mobile) Write() {
	fmt.Printf("正在写入%s手机数据，速度为：%d\n", m.name, m.wspeed)
}
func (u *UDisk) Read() {
	fmt.Printf("正在读取%sU盘数据，速度为：%d\n", u.name, u.rspeed)
}
func (u *UDisk) Write() {
	fmt.Printf("正在写入%sU盘数据，速度为：%d\n", u.name, u.wspeed)
}

// UseDev 4.多态实现 将接口作为函数参数
func UseDev(usb USBer) {
	usb.Read()
	usb.Write()
}
func main04() {
	//接口类型
	var usb USBer

	var m Mobile = Mobile{
		USBDev: USBDev{
			id:     101,
			name:   "华为",
			rspeed: 100,
			wspeed: 200,
		},
		call: "华为5G通话",
	}
	usb = &m
	UseDev(usb)

	fmt.Println("====================================")

	var u UDisk = UDisk{USBDev{
		id:     102,
		name:   "U盘",
		rspeed: 1000,
		wspeed: 2000,
	}}
	usb = &u
	UseDev(usb)
}
