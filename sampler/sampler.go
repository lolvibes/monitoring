package main

import (
	"fmt"
	"time"

	"github.com/prometheus/procfs"
)

func main() {
	fs, err := procfs.NewFS("/proc")
	if err != nil {
		fmt.Print(err)
	}
	stat, err := fs.Stat()
	if err != nil {
		fmt.Print(err)
	}
	fmt.Printf("cpu total user ,", stat.CPUTotal.User, "/n ")
	fmt.Printf("cpu total ", stat.CPUTotal, "/n ")
	cpu := stat.CPUTotal
	fmt.Println("=== Aggregated System CPU Values (in seconds) ===")
	fmt.Printf("User:       %.2f\n", cpu.User)
	fmt.Printf("Nice:       %.2f\n", cpu.Nice)
	fmt.Printf("System:     %.2f\n", cpu.System)
	fmt.Printf("Idle:       %.2f\n", cpu.Idle)
	fmt.Printf("Iowait:     %.2f\n", cpu.Iowait)
	fmt.Printf("IRQ:        %.2f\n", cpu.IRQ)
	fmt.Printf("SoftIRQ:    %.2f\n", cpu.SoftIRQ)
	fmt.Printf("Steal:      %.2f\n", cpu.Steal)
	fmt.Printf("Guest:      %.2f\n", cpu.Guest)
	fmt.Printf("GuestNice:  %.2f\n", cpu.GuestNice)
	idletime := cpu.Idle + cpu.Iowait
	workingtime := cpu.User + cpu.System + cpu.IRQ
	time.Sleep(1 * time.Second)
	fs2, err := procfs.NewFS("/proc")
	if err != nil {
		fmt.Print(err)
	}
	stat2, err := fs2.Stat()
	if err != nil {
		fmt.Print(err)
	}

	cpu2 := stat2.CPUTotal
	idletime1 := cpu2.Idle + cpu2.Iowait
	fmt.Println("idlet time ", idletime1)
	workingtime1 := cpu2.User + cpu2.System + cpu2.IRQ
	activetime := workingtime
	totaltime := idletime + workingtime
	fmt.Println("active time and total time", activetime, totaltime)
	activetime1 := workingtime1

	totaltime1 := idletime1 + workingtime1
	fmt.Println("active time 1 and totaltime1", activetime1, totaltime1)
	delatactive := activetime1 - activetime
	fmt.Println(delatactive)
	delatatotle := totaltime1 - totaltime
	fmt.Println(delatatotle)
	cpupercentage := (delatactive / delatatotle) * 100
	fmt.Printf("\n cpu perecentage %.3f%%", cpupercentage)
}
