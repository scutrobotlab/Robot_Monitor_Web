package SerialHandle

import (
	"encoding/json"
	"errors"
	"log"
	"strings"
	"time"

	"../DataPack"

	"go.bug.st/serial.v1"
)

var MySerialPort serial.Port

type DataToSerial struct {
	Board uint8
	Act   uint8
	Type  string
	Addr  uint32
	Data  float64
}

func FindSerialPorts() []string {
	tmp, err := serial.GetPortsList()
	if err != nil {
		log.Fatal(err)
	}
	if len(tmp) == 0 {
		log.Fatal("No serial ports found!")
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
	} else {
		return errors.New("empty serial port")
	}
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

func (p *DataToSerial) GetDataFromWeb(d *DataPack.DataFromWeb_t) {
	p.Board = d.Board
	p.Act = d.Act
	p.Type = d.Type
	p.Addr = d.Addr
	p.Data = d.Data
}

func (p *DataToSerial) Send() error {
	data := make([]byte, 3)
	data[0] = byte(p.Board)
	data[1] = p.Act
	data[2] = byte(DataPack.TypeLen[p.Type])
	a := DataPack.AnyToBytes(p.Addr)
	data = append(data, a...)
	b := make([]byte, 8)
	if p.Act == DataPack.ACT_WRITE {
		switch p.Type {
		case "uint8_t":
			b = DataPack.AnyToBytes(uint8(p.Data))
		case "uint16_t":
			b = DataPack.AnyToBytes(uint16(p.Data))
		case "uint32_t":
			b = DataPack.AnyToBytes(uint32(p.Data))
		case "uint64_t":
			b = DataPack.AnyToBytes(uint64(p.Data))
		case "int8_t":
			b = DataPack.AnyToBytes(int8(p.Data))
		case "int16_t":
			b = DataPack.AnyToBytes(int16(p.Data))
		case "int32_t", "int":
			b = DataPack.AnyToBytes(int32(p.Data))
		case "int64_t":
			b = DataPack.AnyToBytes(int64(p.Data))
		case "float":
			b = DataPack.AnyToBytes(float32(p.Data))
		case "double":
			b = DataPack.AnyToBytes(float64(p.Data))
		default:
			b = DataPack.AnyToBytes(p.Data)
		}
	}
	data = append(data, b...)
	data = append(data, '\n')
	return SerialSend(data)
}

func SerialParse(jsonString chan string) {
	var b []byte
	for {
		if MySerialPort != nil {
			var chatPack DataPack.DataToChat
			var chatData DataPack.DataToChat_t
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
					chatData.Board = buff[i*16]
					addr := DataPack.BytesToUint32(buff[i*16+3 : i*16+7])
					for _, v := range DataPack.DataToRead {
						if v.Addr == addr {
							chatData.Name = v.Name
							switch v.Type {
							case "uint8_t":
								chatData.Data = float64(DataPack.BytesToUint8(buff[i*16+7 : i*16+15]))
							case "uint16_t":
								chatData.Data = float64(DataPack.BytesToUint16(buff[i*16+7 : i*16+15]))
							case "uint32_t":
								chatData.Data = float64(DataPack.BytesToUint32(buff[i*16+7 : i*16+15]))
							case "uint64_t":
								chatData.Data = float64(DataPack.BytesToUint64(buff[i*16+7 : i*16+15]))
							case "int8_t":
								chatData.Data = float64(DataPack.BytesToInt8(buff[i*16+7 : i*16+15]))
							case "int16_t":
								chatData.Data = float64(DataPack.BytesToInt16(buff[i*16+7 : i*16+15]))
							case "int32_t", "int":
								chatData.Data = float64(DataPack.BytesToInt32(buff[i*16+7 : i*16+15]))
							case "int64_t":
								chatData.Data = float64(DataPack.BytesToInt64(buff[i*16+7 : i*16+15]))
							case "float":
								chatData.Data = float64(DataPack.BytesToFloat32(buff[i*16+7 : i*16+15]))
							case "double":
								chatData.Data = float64(DataPack.BytesToFloat64(buff[i*16+7 : i*16+15]))
							default:
								chatData.Data = 0
							}
						}
					}
					chatPack.DataPack = append(chatPack.DataPack, chatData)
					b, _ = json.Marshal(chatPack)
					jsonString <- string(b)
				} else {
					log.Println("Invalid data pack")
				}
			}
		}
		time.Sleep(9 * time.Millisecond)
	}
}
