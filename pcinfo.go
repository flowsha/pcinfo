// Copyright 2017 The pcinfo Liusha. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Program provides functions for get PC system information.
package main

import (
	"fmt"
	//"database/sql"
	"github.com/StackExchange/wmi"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type Win32_BIOS struct {
	InstallDate string
}

type Win32_ComputerSystem struct {
	Name string
	Manufacturer string
	Model string
	SystemType string
	PCSystemType uint16
	TotalPhysicalMemory uint64
}

type Win32_Processor struct {
	Name string
	Manufacturer string
	SerialNumber string
	NumberOfCores uint32
}

type Win32_PhysicalMemory struct {
	Name string
	Manufacturer string
	Speed uint32
	Capacity uint64
}

type Win32_NetworkAdapterConfiguration struct {
	MACAddress string
	IPAddress []string
}

type Win32_DiskDrive struct {
	Name string
	Manufacturer string
	Model string
	SerialNumber string
	Size uint64
	InterfaceType string
}

type Win32_DesktopMonitor struct {
	//MonitorManufacturer string
	MonitorType string
	Name string
}

type PCInfo struct {
	ComputerName string
	ComputerManufacturer string
	ComputerModel string
	SystemType string
	PCSystemType uint16
	TotalPhysicalMemory uint64

}

func main() {

	var (
	    department string
	    username string
	)

	for {
		fmt.Println("Hello, This is pcinfo!")
		fmt.Print("请输入所属部门[50字符]：")
		fmt.Scanln(&department)
		if len([]rune(department)) > 50 {
			fmt.Println("所属部门超过50个字符")
			continue
		}
		fmt.Print("请输入使用人姓名[20字符]：")
		fmt.Scanln(&username)
		if len([]rune(username)) > 20 {
			fmt.Println("使用人姓名超过20个字符")
			continue
		}

		fmt.Println("所属部门：", department)
		fmt.Println("使用人：", username)

		//get Win32_ComputerSystem information
		var computer []Win32_ComputerSystem
		q := wmi.CreateQuery(&computer, "")
		err := wmi.Query(q, &computer)
		if err != nil {
			log.Fatal(err)
		}
		for _, v := range computer {
			fmt.Println(v.Name)
			fmt.Println(v.Manufacturer)
			fmt.Println(v.Model)
			fmt.Println(v.TotalPhysicalMemory)
			fmt.Println(v.SystemType)
			fmt.Println(v.PCSystemType)
		}

		//get Win32_Processor information
		var processor []Win32_Processor
		q = wmi.CreateQuery(&processor, "")
		err = wmi.Query(q, &processor)
		if err != nil {
			log.Fatal(err)
		}
		for _, v := range processor {
			fmt.Println(v.Name)
			fmt.Println(v.Manufacturer)
			fmt.Println(v.SerialNumber)
			fmt.Println(v.NumberOfCores)
		}

		//get Win32_PhysicalMemory information
		var memory []Win32_PhysicalMemory
		q = wmi.CreateQuery(&memory, "")
		err = wmi.Query(q, &memory)
		if err != nil {
			log.Fatal(err)
		}
		for _, v := range memory {
			fmt.Println(v.Name)
			fmt.Println(v.Manufacturer)
			fmt.Println(v.Speed)
			fmt.Println(v.Capacity)
		}

		//get Win32_NetworkAdapterConfiguration information
		var network []Win32_NetworkAdapterConfiguration
		q = wmi.CreateQuery(&network, "WHERE IPEnabled=TRUE")
		err = wmi.Query(q, &network)
		if err != nil {
			log.Fatal(err)
		}
		for _, v := range network {
			fmt.Println(v.MACAddress)
			for _, ip := range v.IPAddress {
				fmt.Println(ip)
			}
		}

		//get Win32_DiskDrive information
		var disk []Win32_DiskDrive
		q = wmi.CreateQuery(&disk, "")
		err = wmi.Query(q, &disk)
		if err != nil {
			log.Fatal(err)
		}
		for _, v := range disk {
			fmt.Println(v.Name)
			fmt.Println(v.Manufacturer)
			fmt.Println(v.Model)
			fmt.Println(v.SerialNumber)
			fmt.Println(v.Size)
			fmt.Println(v.InterfaceType)
		}

		//get Win32_DesktopMonitor information
		var monitor []Win32_DesktopMonitor
		q = wmi.CreateQuery(&monitor, "")
		err = wmi.Query(q, &monitor)
		if err != nil {
			log.Fatal(err)
		}
		for _, v := range monitor {
			fmt.Println(v.Name)
			//fmt.Println(v.MonitorManufacturer)
			fmt.Println(v.MonitorType)
		}


		//db, err := sql.Open("mysql", "root:@/test?charset=utf8")



		fmt.Print("请按任意键退出...")
		fmt.Scanln()
		return
	}

}

