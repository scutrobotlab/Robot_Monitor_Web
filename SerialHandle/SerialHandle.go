package serialhandle

import (
	"encoding/json"
	"errors"
	"log"
	"strings"
	"time"

	"../datapack"

	"go.bug.st/serial.v1"
)

type currentSerialPortT struct {
	Name     string
	BaudRate int
}

var CurrentSerialPort currentSerialPortT
var MySerialPort serial.Port

func FindSerialPorts() []string {
	tmp, err := serial.GetPortsList()
	if err != nil {
		log.Fatal(err)
	}
	if len(tmp) == 0 {
		log.Println("No serial ports found!")
		return nil
	}
	var ports []string
	for _, port := range tmp {
		if strings.Contains(port, "USB") || strings.Contains(port, "ACM") || strings.Contains(port, "COM") {
			ports = append(ports, port)
		}
	}
	return ports
}

func OpenSerialPort(portName string, baudRate int) error {
	mode := &serial.Mode{
		BaudRate: baudRate,
	}
	var err error
	MySerialPort, err = serial.Open(portName, mode)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func CloseSerialPort() error {
	if MySerialPort != nil {
		return MySerialPort.Close()
	}
	return errors.New("empty serial port")
}

func SerialSend(data []byte) error {
	_, err := MySerialPort.Write(data)
	if err != nil {
		return err
	}
	return nil
}

func SerialReceive() ([]byte, error) {
	buff := make([]byte, 160)
	n, err := MySerialPort.Read(buff)
	if err != nil {
		return nil, err
	}
	if n == 0 {
		return buff[0:0], nil
	}
	return buff[:n], nil
}

func SerialSendCmd(act uint8, variable datapack.VariableT) error {
	if MySerialPort != nil && CurrentSerialPort.Name != "" {
		data := make([]byte, 3)
		data[0] = byte(variable.Board)
		data[1] = act
		data[2] = byte(datapack.TypeLen[variable.Type])
		a := datapack.AnyToBytes(variable.Addr)
		data = append(data, a...)
		b := make([]byte, 8)
		if act == datapack.ACT_WRITE {
			b = datapack.SpecToBytes(variable.Type, variable.Data)
		}
		data = append(data, b...)
		data = append(data, '\n')
		return SerialSend(data)
	}
	return errors.New("No serial port")
}

func SerialParse(jsonString chan string) {
	var b []byte
	for {
		if MySerialPort != nil && CurrentSerialPort.Name != "" {
			var chartPack datapack.DataToChart
			var chartData datapack.DataToChartT
			buff, err := SerialReceive()
			if err != nil {
				log.Println("Fail: Can't receive serial data")
			}
			if len(buff)%16 != 0 {
				log.Println("Invalid data length")
			}
			packNum := len(buff) / 16
			for i := 0; i < packNum; i++ {
				if buff[i*16+1] == 2 {
					chartData.Board = buff[i*16]
					addr := datapack.BytesToUint32(buff[i*16+3 : i*16+7])
					for _, v := range datapack.CurrentVariables.Variables {
						if v.Addr == addr {
							chartData.Name = v.Name
							switch v.Type {
							case "uint8_t":
								chartData.Data = float64(datapack.BytesToUint8(buff[i*16+7 : i*16+15]))
							case "uint16_t":
								chartData.Data = float64(datapack.BytesToUint16(buff[i*16+7 : i*16+15]))
							case "uint32_t":
								chartData.Data = float64(datapack.BytesToUint32(buff[i*16+7 : i*16+15]))
							case "uint64_t":
								chartData.Data = float64(datapack.BytesToUint64(buff[i*16+7 : i*16+15]))
							case "int8_t":
								chartData.Data = float64(datapack.BytesToInt8(buff[i*16+7 : i*16+15]))
							case "int16_t":
								chartData.Data = float64(datapack.BytesToInt16(buff[i*16+7 : i*16+15]))
							case "int32_t", "int":
								chartData.Data = float64(datapack.BytesToInt32(buff[i*16+7 : i*16+15]))
							case "int64_t":
								chartData.Data = float64(datapack.BytesToInt64(buff[i*16+7 : i*16+15]))
							case "float":
								chartData.Data = float64(datapack.BytesToFloat32(buff[i*16+7 : i*16+15]))
							case "double":
								chartData.Data = float64(datapack.BytesToFloat64(buff[i*16+7 : i*16+15]))
							default:
								chartData.Data = 0
							}
						}
					}
					chartPack.DataPack = append(chartPack.DataPack, chartData)
					b, _ = json.Marshal(chartPack)
					jsonString <- string(b)
				} else {
					log.Println("Invalid data pack")
				}
			}
		}
		time.Sleep(9 * time.Millisecond)
	}
}
