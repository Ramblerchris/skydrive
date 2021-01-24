package utils

import (
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"time"
)

func GetCpuPercent() float64 {
	percent, _:= cpu.Percent(time.Second, false)
	return percent[0]
}

func GetMemPercent()(total uint64,percent float64) {
	memInfo, _ := mem.VirtualMemory()
	return  memInfo.Total,memInfo.UsedPercent
}

func GetSwapMemoryPercent()(total uint64,percent float64){
	memInfo, _ := mem.SwapMemory()
	return  memInfo.Total,memInfo.UsedPercent
}

func GetDiskPercent() (total uint64,percent float64) {
	parts, _ := disk.Partitions(true)
	diskInfo, _ := disk.Usage(parts[0].Mountpoint)
	return diskInfo.Total,diskInfo.UsedPercent
}