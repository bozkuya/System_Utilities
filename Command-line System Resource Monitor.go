package main

import (
	"fmt"
	"time"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/disk"
)

func main() {
	for {
		cpuPercent, _ := cpu.Percent(time.Second, false)
		vmStat, _ := mem.VirtualMemory()
		diskStat, _ := disk.Usage("/")

		fmt.Printf("CPU Usage: %.2f%%\n", cpuPercent[0])
		fmt.Printf("Memory Usage: %.2f%%\n", vmStat.UsedPercent)
		fmt.Printf("Disk Usage: %.2f%%\n", diskStat.UsedPercent)

		time.Sleep(time.Second * 5)
	}
}
