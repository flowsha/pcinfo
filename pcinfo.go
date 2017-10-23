// Copyright 2017 The pcinfo Liusha. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Program provides functions for get PC system information.
package main

import (
	"crypto/sha1"
	"fmt"
	"strconv"
	"strings"
	//"time"
	//"database/sql"
	"github.com/StackExchange/wmi"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"log"
)

type Win32_BIOS struct {
	Name          string
	Manufacturer  string
	SerialNumber  string
	ReleaseDate   string
}

type Win32_ComputerSystem struct {
	Name                string
	Manufacturer        string
	Model               string
	SystemType          string
	PCSystemType        uint16
	TotalPhysicalMemory uint64
}

type Win32_Processor struct {
	Name          string
	Manufacturer  string
	SerialNumber  string
	NumberOfCores uint32
}

type Win32_PhysicalMemory struct {
	Name         string
	Manufacturer string
	Speed        uint32
	Capacity     uint64
}

type Win32_NetworkAdapterConfiguration struct {
	MACAddress string
	IPAddress  []string
}

type Win32_DiskDrive struct {
	Name          string
	Manufacturer  string
	Model         string
	SerialNumber  string
	Size          uint64
	InterfaceType string
}

type Win32_DesktopMonitor struct {
	//MonitorManufacturer string
	MonitorType string
	Name        string
}

type Win32_Printer struct {
	Name string
}

type PCInfo struct {
	ComputerName           string
	ComputerManufacturer   string
	ComputerModel          string
	SystemType             string
	PCSystemType           string
	TotalPhysicalMemory    string
	NumberOfPhysicalMemory uint16
	ProcessorName          string
	ProcessorManufacturer  string
	ProcessorNumberOfCores uint32
	NetworkAdapter         []Win32_NetworkAdapterConfiguration
	DiskDrive              []Win32_DiskDrive
	MonitorName            string
	MonitorType            string
	PrinterName            string
	BIOSName               string
	BIOSManufacturer       string
	BIOSReleaseDate        string
	BIOSSerialNumber       string
}

func SHA1(data string) string {
	t := sha1.New()
	io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}

func main() {

	var (
		department string
		username   string
		pc         PCInfo
	)

	for {
		fmt.Print("Hello, This is pcinfo!\n\n")

		// Input department and username information.
		fmt.Print("请输入所属部门[50字符]：")
		fmt.Scanln(&department)
		department = strings.TrimSpace(department)
		if len([]rune(department)) > 50 {
			fmt.Println("所属部门超过50个字符")
			continue
		}
		fmt.Print("请输入使用人姓名[20字符]：")
		fmt.Scanln(&username)
		username = strings.TrimSpace(username)
		if len([]rune(username)) > 20 {
			fmt.Println("使用人姓名超过20个字符")
			continue
		}
		fmt.Println("所属部门：", department)
		fmt.Println("使用人：", username)
		fmt.Print("\n开始获取电脑系统配置...\n\n")

		// get Win32_ComputerSystem information
		var computer []Win32_ComputerSystem
		q := wmi.CreateQuery(&computer, "")
		err := wmi.Query(q, &computer)
		if err != nil {
			log.Fatal(err)
		}
		for _, v := range computer {
			pc.ComputerName = strings.TrimSpace(v.Name)
			pc.ComputerManufacturer = strings.TrimSpace(v.Manufacturer)
			pc.ComputerModel = strings.TrimSpace(v.Model)
			pc.SystemType = strings.TrimSpace(v.SystemType)
			switch v.PCSystemType {
			case 1:
				pc.PCSystemType = "台式电脑"
			case 2:
				pc.PCSystemType = "笔记本"
			default:
				pc.PCSystemType = "其他"
			}
			pc.TotalPhysicalMemory = strconv.FormatUint(v.TotalPhysicalMemory/1000000000, 10) + "G"
			fmt.Println("主机名称：", pc.ComputerName)
			fmt.Println("生产厂商：", pc.ComputerManufacturer)
			fmt.Println("主机型号：", pc.ComputerModel)
			fmt.Println("物理内存：", pc.TotalPhysicalMemory)
			fmt.Println("系统架构：", pc.SystemType)
			fmt.Println("设备类型：", pc.PCSystemType)
		}

		// get Win32_BIOS information
		var bios []Win32_BIOS
		q = wmi.CreateQuery(&bios, "WHERE PrimaryBIOS=TRUE")
		err = wmi.Query(q, &bios)
		if err != nil {
			log.Fatal(err)
		}
		for _, v := range bios {
			pc.BIOSName = strings.TrimSpace(v.Name)
			pc.BIOSManufacturer = strings.TrimSpace(v.Manufacturer)
			pc.BIOSSerialNumber = strings.TrimSpace(v.SerialNumber)
			pc.BIOSReleaseDate = string([]rune(strings.TrimSpace(v.ReleaseDate))[0:8])
			fmt.Println("BIOS名称：", pc.BIOSName)
			fmt.Println("BIOS生产商：", pc.BIOSManufacturer)
			fmt.Println("BIOS序列号：", pc.BIOSSerialNumber)
			fmt.Println("BIOS出厂日期：", pc.BIOSReleaseDate)
		}

		// get Win32_Processor information
		var processor []Win32_Processor
		q = wmi.CreateQuery(&processor, "")
		err = wmi.Query(q, &processor)
		if err != nil {
			log.Fatal(err)
		}
		for _, v := range processor {
			pc.ProcessorName = strings.TrimSpace(v.Name)
			pc.ProcessorManufacturer = strings.TrimSpace(v.Manufacturer)
			pc.ProcessorNumberOfCores = v.NumberOfCores
			fmt.Println("CPU型号：", pc.ProcessorName)
			fmt.Println("CPU生产商：", pc.ProcessorManufacturer)
			//fmt.Println(v.SerialNumber)
			fmt.Println("CPU核数：", pc.ProcessorNumberOfCores)
		}

		// get Win32_PhysicalMemory information
		var memory []Win32_PhysicalMemory
		q = wmi.CreateQuery(&memory, "")
		err = wmi.Query(q, &memory)
		if err != nil {
			log.Fatal(err)
		}
		for i, v := range memory {
			pc.NumberOfPhysicalMemory += 1
			//fmt.Println(v.Name)
			//fmt.Println(v.Manufacturer)
			fmt.Printf("内存频率[%d]：%s\n", i+1, strconv.FormatUint(uint64(v.Speed), 10)+"MHz")
			fmt.Printf("内存容量[%d]：%s\n", i+1, strconv.FormatUint(v.Capacity/1000000000, 10)+"G")
		}
		fmt.Println("内存数量：", pc.NumberOfPhysicalMemory)

		// get Win32_NetworkAdapterConfiguration information
		var network []Win32_NetworkAdapterConfiguration
		q = wmi.CreateQuery(&network, "WHERE IPEnabled=TRUE")
		err = wmi.Query(q, &network)
		if err != nil {
			log.Fatal(err)
		}
		pc.NetworkAdapter = network
		for i, v := range network {
			fmt.Printf("网卡[%d]MAC: %s\n", i+1, v.MACAddress)
			for _, ip := range v.IPAddress {
				fmt.Printf("网卡[%d]IP: %s\n", i+1, ip)
			}
		}

		// get Win32_DiskDrive information
		var disk []Win32_DiskDrive
		q = wmi.CreateQuery(&disk, "")
		err = wmi.Query(q, &disk)
		if err != nil {
			log.Fatal(err)
		}
		pc.DiskDrive = disk
		for i, v := range disk {
			//fmt.Printf("硬盘名称[%d]：%s\n", i+1, v.Name)
			//fmt.Printf("硬盘厂商[%d]：%s\n", i+1, v.Manufacturer)
			fmt.Printf("硬盘型号[%d]：%s\n", i+1, strings.TrimSpace(v.Model))
			fmt.Printf("硬盘序列号[%d]：%s\n", i+1, strings.TrimSpace(v.SerialNumber))
			fmt.Printf("硬盘容量[%d]：%s\n", i+1, strconv.FormatUint(v.Size/1000000000, 10)+"G")
			fmt.Printf("硬盘接口类型[%d]：%s\n", i+1, v.InterfaceType)
		}

		// get Win32_DesktopMonitor information
		var monitor []Win32_DesktopMonitor
		q = wmi.CreateQuery(&monitor, "")
		err = wmi.Query(q, &monitor)
		if err != nil {
			log.Fatal(err)
		}
		for i, v := range monitor {
			pc.MonitorName = strings.TrimSpace(v.Name)
			pc.MonitorType = strings.TrimSpace(v.MonitorType)
			fmt.Printf("显示器名称[%d]：%s\n", i+1, pc.MonitorName)
			//fmt.Println(v.MonitorManufacturer)
			fmt.Printf("显示器类型[%d]：%s\n", i+1, pc.MonitorType)
		}

		// get Win32_Printer information
		var printer []Win32_Printer
		q = wmi.CreateQuery(&printer, "WHERE Default=TRUE")
		err = wmi.Query(q, &printer)
		if err != nil {
			log.Fatal(err)
		}
		for i, v := range printer {
			pc.PrinterName = strings.TrimSpace(v.Name)
			fmt.Printf("打印机名称[%d]：%s\n", i+1, pc.PrinterName)
		}

		fmt.Println("MAC SHA1 Hash:", SHA1(pc.NetworkAdapter[0].MACAddress))
		//db, err := sql.Open("mysql", "root:@/test?charset=utf8")

		fmt.Print("\n已经成功获取电脑配置信息，并录入数据库！\n\n")
		fmt.Print("请按任意键退出...")
		fmt.Scanln()
		return
	}
}
