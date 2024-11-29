package serialport

import (
	"github.com/tarm/serial"
	"log"
)

type SerialPort struct {
	port *serial.Port
}

type ReceiveDataListener func(s SerialPort, data []byte)

var Run = true

func NewSerialPort(path string, baud int, listener ReceiveDataListener) SerialPort {
	config := &serial.Config{
		Name: path,
		Baud: baud,
	}
	port, err := serial.OpenPort(config)
	if err != nil {
		log.Fatal("打开串口失败:", err)
		return SerialPort{}
	}
	serialPort := SerialPort{
		port: port,
	}
	Run = true
	go func() {
		buf := make([]byte, 512)
		for Run {
			n, err := port.Read(buf)
			if err != nil {
				continue
			}
			bytes := buf[:n]
			if listener != nil {
				listener(serialPort, bytes)
			}
		}
	}()
	return serialPort
}

func (p SerialPort) Send(data []byte) {
	_, err := p.port.Write(data)
	if err != nil {
		return
	}
}

func (p SerialPort) Close() {
	p.Close()
	Run = false
}
