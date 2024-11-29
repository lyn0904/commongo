package main

import (
	"github.com/lyn0904/commongo/common/other"
	"github.com/lyn0904/commongo/common/serialport"
	"log"
)

func main() {
	_ = serialport.NewSerialPort("/dev/ttyS6", 115200, func(s serialport.SerialPort, data []byte) {
		log.Println("串口接收到的值:", serialport.BytesToHeX(data))
	})
	other.Blocking()
}
