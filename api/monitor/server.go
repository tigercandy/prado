package monitor

import (
	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"github.com/tigercandy/prado/internal/models/common/response"
	"net/http"
	"runtime"
)

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

func ServerInfo(c *gin.Context) {
	osDict := make(map[string]interface{}, 0)
	osDict["goOS"] = runtime.GOOS
	osDict["arch"] = runtime.GOARCH
	osDict["mem"] = runtime.MemProfileRate
	osDict["compiler"] = runtime.Compiler
	osDict["version"] = runtime.Version()
	osDict["numGoroutine"] = runtime.NumGoroutine()

	myDisk, _ := disk.Usage("/")
	diskTotalGB := int(myDisk.Total) / GB
	diskFreeGB := int(myDisk.Free) / GB
	diskDict := make(map[string]interface{}, 0)
	diskDict["total"] = diskTotalGB
	diskDict["free"] = diskFreeGB

	myMem, _ := mem.VirtualMemory()
	memUsedMB := int(myMem.Used) / GB
	memTotalMB := int(myMem.Total) / GB
	memFreeMB := int(myMem.Free) / GB
	memUsedPercent := int(myMem.UsedPercent)
	memDict := make(map[string]interface{}, 0)
	memDict["total"] = memTotalMB
	memDict["used"] = memUsedMB
	memDict["free"] = memFreeMB
	memDict["usage"] = memUsedPercent

	cpuDict := make(map[string]interface{}, 0)
	cpuDict["cpuNum"], _ = cpu.Counts(false)

	response.Custom(c, gin.H{
		"code": http.StatusOK,
		"os":   osDict,
		"disk": diskDict,
		"mem":  memDict,
		"cpu":  cpuDict,
	})
}
