package health

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"net/http"
)

const (
	B = 1
	KB = 2014 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

// 心跳检测
func Pong(ctx *gin.Context) {
	var (
		message string
	)
	message = "ok"
	ctx.String(http.StatusOK, message)
}



// 磁盘使用率检测
func DiskCheck(ctx *gin.Context) {
	var (
		u *disk.UsageStat
		err error
		useMB int
		useGB int
		totalMB int
		totalGB int
		usedPercent float64
		result string
		status int
	)
	status = http.StatusOK

	if u, err = disk.Usage("/"); err != nil {
		result = fmt.Sprintf("获取服务器磁盘信息失败：%v", err)
		ctx.String(status,result)
		return
	}

	result = "正常"

	useMB = int(u.Used) / MB
	useGB = int(u.Used) / GB
	totalMB = int(u.Total) / MB
	totalGB = int(u.Total) / GB

	usedPercent = u.UsedPercent

	if usedPercent > 95 {
		result = "危险"
	}else if usedPercent > 90 {
		result = "警告"
	}



	result = fmt.Sprintf("%s - Free space: %dMB (%dGB) / %dMB (%dGB) | Used:%.2f", result, useMB, useGB, totalMB, totalGB, usedPercent)

	ctx.String(status, result)

	return

}


// 内存使用率检测
func MemoryCheck(ctx *gin.Context) {
	var (
		memory *mem.VirtualMemoryStat
		err error
		status int
		result string
		usedMB int
		usedGB int
		totalMB int
		totalGB int
		usedPercent float64
	)
	status = http.StatusOK
	if memory, err = mem.VirtualMemory(); err != nil {
		result = fmt.Sprintf("获取服务器内存信息失败：%v", err)
		ctx.String(status, result)
		return
	}

	usedMB = int(memory.Used) / MB
	usedGB = int(memory.Used) / GB
	totalMB = int(memory.Total) / MB
	totalGB = int(memory.Total) / GB
	usedPercent = memory.UsedPercent

	if usedPercent > 95 {
		result = "危险"
	} else if usedPercent > 90 {
		result = "警告"
	}

	result = fmt.Sprintf("%s - Free space: %dMB (%dGB) / %dMB (%dGB) | Used:%.2f", result, usedMB, usedGB, totalMB, totalGB, usedPercent)

	ctx.String(status, result)
	return

}


// cpu使用检测
func CpuCheck(ctx *gin.Context) {
	var (
		cores int
		err error
		status int
		result string
		a *load.AvgStat
		everyOne float64 // 最近1分钟中的平均负载
		everyFive float64
		everyFifteen float64
	)
	status = http.StatusOK
	// false 物理cpu个数  true逻辑cpu个数
	if cores, err = cpu.Counts(false); err != nil {
		result = "获取服务器cpu信息失败！"
		ctx.String(status, result)
		return
	}
	if a, err = load.Avg(); err != nil {
		result = "获取服务器cpu平均负载失败！"
		ctx.String(status, result)
		return
	}

	result = "OK"
	everyOne = a.Load1
	everyFive = a.Load5
	everyFifteen = a.Load15

	result = fmt.Sprintf("%s - Load average: 最近1分钟的平均负载：%.2f,最近5分钟的平均负载 %.2f, ：最近15分钟的平均负载：%.2f | 物理CPU核数: %d", result, everyOne, everyFive, everyFifteen, cores)

	ctx.String(status, result)

}
