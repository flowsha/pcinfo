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
	"database/sql"
	"github.com/StackExchange/wmi"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"log"
)

// Win32BIOS struct defined. 
type Win32BIOS struct {
	Name         string
	Manufacturer string
	SerialNumber string
	ReleaseDate  string
}

// Win32ComputerSystem struct defined.
type Win32ComputerSystem struct {
	Name                string
	Manufacturer        string
	Model               string
	SystemType          string
	PCSystemType        uint16
	TotalPhysicalMemory uint64
}

// Win32Processor struct defined.
type Win32Processor struct {
	Name          string
	Manufacturer  string
	SerialNumber  string
	NumberOfCores uint32
}

// Win32PhysicalMemory struct defined.
type Win32PhysicalMemory struct {
	Name         string
	Manufacturer string
	Speed        uint32
	Capacity     uint64
}

// Win32NetworkAdapterConfiguration struct defined.
type Win32NetworkAdapterConfiguration struct {
	MACAddress string
	IPAddress  []string
}

// Win32DiskDrive struct defined.
type Win32DiskDrive struct {
	Name          string
	Manufacturer  string
	Model         string
	SerialNumber  string
	Size          uint64
	InterfaceType string
}

// Win32DesktopMonitor struct defined.
type Win32DesktopMonitor struct {
	//MonitorManufacturer string
	MonitorType string
	Name        string
}

// Win32Printer struct defined.
type Win32Printer struct {
	Name string
}

// PCInfo struct defined.
type PCInfo struct {
	ID                     string
	CorpName               string
	Department             string
	UserName               string
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
	NetworkAdapter         []Win32NetworkAdapterConfiguration
	DiskDrive              []Win32DiskDrive
	MonitorName            string
	MonitorType            string
	PrinterName            string
	BIOSName               string
	BIOSManufacturer       string
	BIOSReleaseDate        string
	BIOSSerialNumber       string
}

// PCInfoInterface interface defined.
type PCInfoInterface interface {
	init()
	GetPCInfo()
	GetPCSnSHA1()
}

// GetPCSnSHA1 function belongs to PCInfo.
func (pc *PCInfo) GetPCSnSHA1() string {
	t := sha1.New()
	io.WriteString(t, pc.BIOSSerialNumber)
	return fmt.Sprintf("%x", t.Sum(nil))
}

// init function belongs to PCInfo.
func (pc *PCInfo) init() {

	for {
		// Input corporation information.
		fmt.Print("请输入所属公司[50字符]：")
		fmt.Scanln(&pc.CorpName)
		pc.CorpName = strings.TrimSpace(pc.CorpName)
		strLength := len([]rune(pc.CorpName))
		if strLength > 50 {
			fmt.Println("所属公司超过50个字符！")
			continue
		} else if strLength == 0 {
			fmt.Println("所属公司不能为空！")
			continue
		}

		// Input corporation information.
		fmt.Print("请输入所属部门[50字符]：")
		fmt.Scanln(&pc.Department)
		pc.Department = strings.TrimSpace(pc.Department)
		strLength = len([]rune(pc.Department))
		if strLength > 50 {
			fmt.Println("所属部门超过50个字符！")
			continue
		} else if strLength == 0 {
			fmt.Println("所属部门不能为空！")
			continue
		}

		// Input username information.
		fmt.Print("请输入使用人姓名[20字符]：")
		fmt.Scanln(&pc.UserName)
		pc.UserName = strings.TrimSpace(pc.UserName)
		strLength = len([]rune(pc.UserName))
		if strLength > 20 {
			fmt.Println("使用人姓名超过20个字符！")
			continue
		} else if strLength == 0 {
			fmt.Println("使用人姓名不能为空！")
			continue
		}

		fmt.Println("所属公司：", pc.CorpName)
		fmt.Println("所属部门：", pc.Department)
		fmt.Println("使用人：", pc.UserName)
		fmt.Print("请确认输入信息是否正确？（yes/no）:")
		var answer string
		fmt.Scanln(&answer)
		if answer != "yes" {
			continue
		} else {
			break
		}
	}
}

// GetPCInfo function belongs to PCInfo.
func (pc *PCInfo) GetPCInfo() {
	fmt.Print("\n开始获取电脑系统配置...\n\n")
	// get Win32_ComputerSystem information
	var computer []Win32ComputerSystem
	//q := wmi.CreateQuery(&computer, "")
	q := "SELECT * FROM Win32_ComputerSystem"
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
	var bios []Win32BIOS
	//q = wmi.CreateQuery(&bios, "WHERE PrimaryBIOS=TRUE")
	q = "SELECT * FROM Win32_BIOS WHERE PrimaryBIOS=TRUE"
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
	var processor []Win32Processor
	//q = wmi.CreateQuery(&processor, "")
	q = "SELECT * FROM Win32_Processor"
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
		fmt.Println("CPU核数：", pc.ProcessorNumberOfCores)
	}

	// get Win32_PhysicalMemory information
	var memory []Win32PhysicalMemory
	//q = wmi.CreateQuery(&memory, "")
	q = "SELECT * FROM Win32_PhysicalMemory"
	err = wmi.Query(q, &memory)
	if err != nil {
		log.Fatal(err)
	}
	for i, v := range memory {
		pc.NumberOfPhysicalMemory++
		fmt.Printf("内存频率[%d]：%s\n", i+1, strconv.FormatUint(uint64(v.Speed), 10)+"MHz")
		fmt.Printf("内存容量[%d]：%s\n", i+1, strconv.FormatUint(v.Capacity/1000000000, 10)+"G")
	}
	fmt.Println("内存数量：", pc.NumberOfPhysicalMemory)

	// get Win32_NetworkAdapterConfiguration information
	var network []Win32NetworkAdapterConfiguration
	//q = wmi.CreateQuery(&network, "WHERE IPEnabled=TRUE")
	q = "SELECT * FROM Win32_NetworkAdapterConfiguration WHERE IPEnabled=TRUE"
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
	var disk []Win32DiskDrive
	//q = wmi.CreateQuery(&disk, "")
	q = "SELECT * FROM Win32_DiskDrive"
	err = wmi.Query(q, &disk)
	if err != nil {
		log.Fatal(err)
	}
	pc.DiskDrive = disk
	for i, v := range disk {
		fmt.Printf("硬盘型号[%d]：%s\n", i+1, strings.TrimSpace(v.Model))
		fmt.Printf("硬盘序列号[%d]：%s\n", i+1, strings.TrimSpace(v.SerialNumber))
		fmt.Printf("硬盘容量[%d]：%s\n", i+1, strconv.FormatUint(v.Size/1000000000, 10)+"G")
		fmt.Printf("硬盘接口类型[%d]：%s\n", i+1, v.InterfaceType)
	}

	// get Win32_DesktopMonitor information
	var monitor []Win32DesktopMonitor
	//q = wmi.CreateQuery(&monitor, "")
	q = "SELECT * FROM Win32_DesktopMonitor"
	err = wmi.Query(q, &monitor)
	if err != nil {
		log.Fatal(err)
	}
	for i, v := range monitor {
		pc.MonitorName = strings.TrimSpace(v.Name)
		pc.MonitorType = strings.TrimSpace(v.MonitorType)
		fmt.Printf("显示器名称[%d]：%s\n", i+1, pc.MonitorName)
		fmt.Printf("显示器类型[%d]：%s\n", i+1, pc.MonitorType)
	}

	// get Win32_Printer information
	var printer []Win32Printer
	//q = wmi.CreateQuery(&printer, "WHERE Default=TRUE AND Local=TRUE AND Network=FALSE")
	q = "SELECT * FROM Win32_Printer WHERE Default=TRUE AND Local=TRUE AND Network=FALSE"
	err = wmi.Query(q, &printer)
	if err != nil {
		log.Fatal(err)
	}
	for i, v := range printer {
		pc.PrinterName = strings.TrimSpace(v.Name)
		fmt.Printf("打印机名称[%d]：%s\n", i+1, pc.PrinterName)
	}
}

// The entry of main program.
func main() {

	var pc PCInfo

	fmt.Print("Hello, This is pcinfo program by liusha.\n\n")
	pc.init()
	pc.GetPCInfo()
	//fmt.Println(pc.GetPCSnSHA1())

	db, err := sql.Open("mysql", "root:xtcdma@tcp(127.0.0.1:3306)/zhwh?charset=utf8")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	//var sn string
	rows, err := db.Query("SELECT V_SN_SHA1 FROM tb_pcinfo WHERE V_SN_SHA1 = ?", pc.GetPCSnSHA1())

    if rows.Next() {
    	println(rows)
	}

	if err != nil {
		log.Fatal(err)
	}
	db.Close()

	fmt.Print("\n已经成功获取电脑配置信息，并录入数据库！\n\n")
	fmt.Print("请按任意键退出...")
	fmt.Scanln()
	return
}
