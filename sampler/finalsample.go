package main

import (
	"fmt"
	"time"

	"github.com/prometheus/procfs"
)

type Sampler struct {
	cpuPercentage float64
}

var Sample *Sampler

type cpufields struct {
	User      float64
	Nice      float64
	System    float64
	Idle      float64
	Iowait    float64
	IRQ       float64
	SoftIRQ   float64
	Steal     float64
	Guest     float64
	GuestNice float64
}

func getCpuTimes(c *cpufields) (float64, float64) {
	idleTime := c.Idle + c.Iowait

	activeTime := c.User + c.Nice + c.System + c.IRQ + c.SoftIRQ + c.Steal
	totaltime := idleTime + activeTime
	return activeTime, totaltime
}

func (c *cpufields) ReadMetric(fs procfs.FS) {
	stat, err := fs.Stat()
	if err != nil {
		fmt.Printf(" error at stat %q", err)
	}

	cpu := stat.CPUTotal
	c.User = cpu.User
	c.Nice = cpu.Nice
	c.System = cpu.System
	c.Idle = cpu.Idle
	c.Iowait = cpu.Iowait
	c.IRQ = cpu.IRQ
	c.SoftIRQ = cpu.SoftIRQ
	c.Steal = cpu.Steal
	c.Guest = cpu.Guest
	c.GuestNice = cpu.GuestNice
}

func calculationOfMetric(current *cpufields, previous *cpufields) {
	active1, total1 := getCpuTimes(previous)
	active2, total2 := getCpuTimes(current)
	deltaActive := active2 - active1
	deltaTotal := total2 - total1
	cpuPercentage := (deltaActive / deltaTotal) * 100
	fmt.Println("my cpu percentageeeee", cpuPercentage)
	Sample = &Sampler{
		cpuPercentage,
	}
}

var (
	snap cpufields
	done bool = false
)

func main() {
	go func() {
		for {
			ticker := time.NewTicker(5 * time.Second)
			fs, err := procfs.NewFS("/proc")
			if err != nil {
				fmt.Println("can't make obj of NewFS %q", err)
			}
			if done != true {
				snap.ReadMetric(fs)
				done = true
			}

			var snap2 cpufields

			select {
			case <-ticker.C:
				snap2.ReadMetric(fs)
			}

			currentSnap := snap2
			fmt.Printf("\n previous snap \n\n %+v", &snap)

			fmt.Printf("\n currentsnap \n\n %+v", currentSnap)
			calculationOfMetric(&currentSnap, &snap)
			snap = currentSnap
		}
	}()
	select {}
}
