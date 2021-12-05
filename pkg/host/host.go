package host

import (
	"log"

	"github.com/jaypipes/ghw"
	"github.com/shirou/gopsutil/host"
)

type Spec struct {
	CPU *CPUInfo
	OS  *OSInfo
	RAM *RAMInfo
}
type OSInfo struct {
	Hostname string
	Name     string
	Platform string
	Version  string
}

type CPUInfo struct {
	Cores   uint32
	Threads uint32
}

type RAMInfo struct {
	Total  float32
	Usable float32
}

func Info() Spec {
	return Spec{
		OS:  getOSInfo(),
		CPU: getCPUInfo(),
		RAM: getRAMInfo(),
	}
}

func getCPUInfo() *CPUInfo {
	c, err := ghw.CPU()
	if err != nil {
		log.Printf("error reading CPU info: %s", err)
		return nil
	}

	return &CPUInfo{
		Cores:   c.TotalCores,
		Threads: c.TotalThreads,
	}
}

func getRAMInfo() *RAMInfo {
	m, err := ghw.Memory()
	if err != nil {
		log.Printf("error reading RAM info: %s", err)
		return nil
	}

	return &RAMInfo{
		Total:  bytesToGB(m.TotalPhysicalBytes),
		Usable: bytesToGB(m.TotalUsableBytes),
	}
}

func bytesToGB(b int64) float32 {
	return float32(b/1000000) / 1000
}

func getOSInfo() *OSInfo {
	h, err := host.Info()
	if err != nil {
		log.Printf("error reading host info: %s", err)
		return nil
	}

	return &OSInfo{
		Hostname: h.Hostname,
		Name:     h.OS,
		Platform: h.Platform,
		Version:  h.PlatformVersion,
	}
}
