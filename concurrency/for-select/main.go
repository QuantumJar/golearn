package main

import "fmt"

//定义了一个警告接口
type Alert interface {
	Alert() string
}

//定义了一个CPU结构体模拟传感器
type CPU struct {
}

//实现方法
func (c CPU) Alert() string {
	return "CPU 温度过高"
}

//定义了一个风扇结构体模拟传感器
type Fan struct {
}

//实现方法
func (f Fan) Alert() string {
	return "风扇 温度过高"
}

func main() {

	fanChannel := make(chan string)
	CPUChannel := make(chan string)

	fan := Fan{}
	cpu := CPU{}

	go func() {
		//CPU发出了温度过高的告警
		for {

			CPUChannel <- cpu.Alert()
		}
	}()

	go func() {
		//风扇发出了温度过高的告警
		for {

			fanChannel <- fan.Alert()
		}
	}()

	for {

		select {
		case fanAlert := <-fanChannel:
			fmt.Printf("%v,请运维人员进行处理。", fanAlert)
		case cpuAlert := <-CPUChannel:
			fmt.Printf("%v,请运维人员进行处理。", cpuAlert)
		}
	}

}
