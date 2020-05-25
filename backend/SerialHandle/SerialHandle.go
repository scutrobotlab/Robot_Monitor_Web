package serialhandle

import (
	"encoding/json"
	"errors"
	"log"
	"strings"
	"time"

	datapack "Robot_Monitor_Web/backend/DataPack"

	"go.bug.st/serial.v1"
)

type currentSerialPortT struct {
	Name     string
	BaudRate int
}

var rxBuff []byte
var CurrentSerialPort currentSerialPortT
var MySerialPort serial.Port
var testPortName = "Test port"

var chOpen = make(chan int)
var chClose = make(chan int)
var chTxBuff = make(chan []byte, 10)

func FindSerialPorts() []string {
	var ports []string
	ports = append(ports, testPortName)

	tmp, err := serial.GetPortsList()
	if err != nil {
		log.Println("Serial ports errors!")
	}
	if len(tmp) == 0 {
		log.Println("No serial ports found!")
	}
	for _, port := range tmp {
		if strings.Contains(port, "USB") || strings.Contains(port, "ACM") || strings.Contains(port, "COM") || strings.Contains(port, "tty.usb") {
			ports = append(ports, port)
		}
	}
	return ports
}

func OpenSerialPort(portName string, baudRate int) error {
	mode := &serial.Mode{
		BaudRate: baudRate,
	}
	if portName == testPortName {
		MySerialPort = newTestPort()
		chOpen <- 1
		return nil
	}
	var err error
	MySerialPort, err = serial.Open(portName, mode)
	if err != nil {
		return err
	}
	chOpen <- 1
	return nil
}

func CloseSerialPort() error {
	if MySerialPort != nil {
		err := MySerialPort.Close()
		if err != nil {
			return err
		}
		chClose <- 1
		return nil
	}
	return errors.New("empty serial port")
}

func SerialSend(data []byte) error {
	_, err := MySerialPort.Write(data)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func SerialReceive() ([]byte, error) {
	buff := make([]byte, 200)
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
		chTxBuff <- data
		return nil
	}
	return errors.New("No serial port")
}

func verifyBuff(buff []byte) (int, int, []byte) {
	if len(buff) < 19 {
		return 0, 0, nil
	}
	if buff[0] == datapack.BOARD_1 && buff[1] == datapack.ACT_SUBSCRIBERETURN && len(buff)%20 == 0 {
		return 0, len(buff) / 20, buff
	}
	b1 := buff[:]
	b2 := buff[:]
	var index int
	for index, v := range buff {
		if v == datapack.BOARD_1 && len(buff[index:]) > 19 && buff[index+1] == datapack.ACT_SUBSCRIBERETURN {
			b1 = b1[index:]
			break
		}
	}
	for i := len(b1) - 1; i > 18; i-- {
		if b1[i] == 10 {
			b2 = b1[:i+1]
			if len(b2)%20 == 0 {
				return index, len(b2) / 20, b2
			}
			b2 = b1[:]
		}
	}
	return 0, 0, nil
}

func SerialPraseThread(chRxBuff chan []byte, chJson chan string) {
	var chartPack datapack.DataToChart
	var chartData datapack.DataToChartT
	for {
		chartPack.DataPack = chartPack.DataPack[0:0]
		chartData.Name = ""
		b := <-chRxBuff
		rxBuff = append(rxBuff, b...)
		index, packNum, buff := verifyBuff(rxBuff)
		rxBuff = rxBuff[index+packNum*20:]
		for i := 0; i < packNum; i++ {
			chartData.Board = buff[i*20]
			chartData.Tick = datapack.BytesToUint32(buff[i*20+15 : i*20+19])
			addr := datapack.BytesToUint32(buff[i*20+3 : i*20+7])
			for _, v := range datapack.VariableRead.Variables {
				if v.Addr == addr {
					chartData.Name = v.Name
					switch v.Type {
					case "uint8_t":
						chartData.Data = float64(datapack.BytesToUint8(buff[i*20+7 : i*20+15]))
					case "uint20_t":
						chartData.Data = float64(datapack.BytesToUint16(buff[i*20+7 : i*20+15]))
					case "uint32_t":
						chartData.Data = float64(datapack.BytesToUint32(buff[i*20+7 : i*20+15]))
					case "uint64_t":
						chartData.Data = float64(datapack.BytesToUint64(buff[i*20+7 : i*20+15]))
					case "int8_t":
						chartData.Data = float64(datapack.BytesToInt8(buff[i*20+7 : i*20+15]))
					case "int20_t":
						chartData.Data = float64(datapack.BytesToInt16(buff[i*20+7 : i*20+15]))
					case "int32_t", "int":
						chartData.Data = float64(datapack.BytesToInt32(buff[i*20+7 : i*20+15]))
					case "int64_t":
						chartData.Data = float64(datapack.BytesToInt64(buff[i*20+7 : i*20+15]))
					case "float":
						chartData.Data = float64(datapack.BytesToFloat32(buff[i*20+7 : i*20+15]))
					case "double":
						chartData.Data = float64(datapack.BytesToFloat64(buff[i*20+7 : i*20+15]))
					default:
						chartData.Data = 0
					}
					break
				}
			}
			if chartData.Name != "" {
				chartPack.DataPack = append(chartPack.DataPack, chartData)
			}
		}
		b, err := json.Marshal(chartPack)
		if err != nil {
			log.Fatalln(err)
		}
		chJson <- string(b)
	}
}

func SerialReceiveThread(chRxBuff chan []byte) {
	for {
		<-chOpen
		for _, v := range datapack.VariableRead.Variables {
			SerialSendCmd(datapack.ACT_SUBSCRIBE, v)
			time.Sleep(10 * time.Millisecond)
		}
	Loop:
		for {
			select {
			case <-chClose:
				break Loop
			default:
				buff, err := SerialReceive()
				if err != nil {
					log.Println("Fail: Can't receive serial data")
				}
				chRxBuff <- buff
				time.Sleep(5 * time.Millisecond)
			}
		}
		time.Sleep(5 * time.Millisecond)
	}
}

func SerialTransmitThread() {
	for {
		b := <-chTxBuff
		err := SerialSend(b)
		if err != nil {
			log.Println(err)
		}
		time.Sleep(3 * time.Millisecond)
	}
}
