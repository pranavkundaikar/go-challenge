package main

import (
"fmt"
"flag"
"time"
"strconv"
"github.com/shirou/gopsutil/cpu"
)

func printCPULoad(){
         cpuStat, err := cpu.Percent(0, true)

         if err != nil {
            fmt.Println(err)
         }

         for idx, cpupercent := range cpuStat {
                 fmt.Println("CPU " + strconv.Itoa(idx) + " utilisation: " + strconv.FormatFloat(cpupercent, 'f', 2, 64) + "%")
         }
         fmt.Println("\n")
}

func main () {
     help :=  "Interval value for CPU load check. Default value 1 second.\n Syntax: ./cpu-load-observer -interval 20s"
     default_interval := time.Duration(1)*time.Second
     var ip = flag.Duration("interval", default_interval, help)
     flag.Parse()

     for {
         printCPULoad()
         time.Sleep(*ip)
     }

}