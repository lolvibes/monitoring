// package main
//
//import (
//	"fmt"
//	"time"
//
//	"github.com/prometheus/procfs"
//)
//
//type Sampler struct {
//	cpuPercentage float64
//}
//
//var Sample *Sampler
//
//type cpufields struct {
//	User      float64
//	Nice      float64
//	System    float64
//	Idle      float64
//	Iowait    float64
//	IRQ       float64
//	SoftIRQ   float64
//	Steal     float64
//	Guest     float64
//	GuestNice float64
//}
//
//var BaseLine cpufields
//
//func getCpuTimes(c *cpufields) (active float64, total float64) {
//	idleTime := c.Idle + c.Iowait
//
//	activeTime := c.User + c.Nice + c.System + c.IRQ + c.SoftIRQ + c.Steal
//	totaltime := idleTime + activeTime
//	return activeTime, totaltime
//}
//
//func CPUpercent(fs procfs.FS, ticker *time.Ticker) {
//	var baseline cpufields
//	var tick2 cpufields
//	baseline.setcpufields(fs)
//	for {
//
//		fmt.Printf(" baseline %+v\n", baseline)
//
//		active1, total1 := getCpuTimes(&baseline)
//
//		select {
//		case <-ticker.C:
//			tick2.setcpufields(fs)
//			fmt.Printf("tick2 %+v \n", tick2)
//			active2, total2 := getCpuTimes(&tick2)
//			deltaActive := active2 - active1
//			deltaTotal := total2 - total1
//			cpuPercentage := (deltaActive / deltaTotal) * 100
//			fmt.Println("my cpu percentageeeee", cpuPercentage)
//			Sample = &Sampler{
//				cpuPercentage: cpuPercentage,
//			}
//		}
//		baseline = tick2
//		if baseline == tick2 {
//			fmt.Print("baseline updated\n")
//		}
//	}
//}
//
//func (c *cpufields) setcpufields(fs procfs.FS) {
//	stat, err := fs.Stat()
//	if err != nil {
//		fmt.Println("err1 :=", err)
//	}
//	cpu := stat.CPUTotal
//	c.User = cpu.User
//	c.System = cpu.System
//	c.Idle = cpu.Idle
//	c.Iowait = cpu.Iowait
//	c.IRQ = cpu.IRQ
//	c.SoftIRQ = cpu.SoftIRQ
//	c.Steal = cpu.Steal
//	c.Guest = cpu.Guest
//	c.GuestNice = cpu.GuestNice
//}
//
//func main() {
//	ticker := time.NewTicker(4 * time.Second)
//	defer ticker.Stop()
//	fs, err := procfs.NewFS("/proc")
//	if err != nil {
//		fmt.Println("err1 :=", err)
//	}
//	go CPUpercent(fs, ticker)
//	select {}
//}
// :%s#^// ##
