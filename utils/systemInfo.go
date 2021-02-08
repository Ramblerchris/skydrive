package utils

import (
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"github.com/skydrive/logger"
	"time"
)

func GetCpuPercent() float64 {
	percent, error:= cpu.Percent(time.Second, false)
	if error!=nil{
		logger.Info("cpuInfo error:",error)
		return 0.00001
	}
	logger.Info("cpuInfo",percent)
	return percent[0]
}
func IOCounters() net.IOCountersStat {
	percent, _:= net.IOCounters(false)
	logger.Info("IOCounters",percent)
	return percent[0]
}

func HostInfo() host.InfoStat {
	percent, _:= host.Info()
	logger.Info("HostInfo",percent)
	return *percent
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