package data

import (
	"fmt"
	"log"
	"runtime"
	"strconv"

	sigar "github.com/cloudfoundry/gosigar"
	"github.com/docker/go-units"
	"github.com/klauspost/cpuid"
	"github.com/prometheus/procfs"
)

func getProcData() map[string]string {
	data := map[string]string{}

	/* cpuid */
	data["VirtCores"] = strconv.Itoa(cpuid.CPU.LogicalCores)

	/* sigar */
	mem := sigar.Mem{}
	mem.Get()
	data["MemTotal"] = units.BytesSize(float64(mem.Total))

	/* Runtime MemStats */
	ms := new(runtime.MemStats)
	runtime.ReadMemStats(ms)
	data["MemUseVirtualRuntime"] = units.BytesSize(float64(ms.Sys))
	data["GcRuns"] = fmt.Sprintf("%d (%d forced)", ms.NumGC, ms.NumForcedGC)

	/* procfs */
	if proc, err := procfs.Self(); err == nil {
		if stat, err := proc.NewStat(); err == nil {
			data["MemUseVirtualTotal"] = units.BytesSize(float64(stat.VirtualMemory()))
			data["MemUsePhysical"] = units.BytesSize(float64(stat.ResidentMemory()))

			data["CpuSelfTime"] = strconv.FormatFloat(stat.CPUTime(), 'f', 2, 64) + "s"
		} else {
			log.Printf("Can't get proc stat: %v", err)
		}
	} else {
		log.Printf("Can't get proc info: %v", err)
	}

	return data
}
