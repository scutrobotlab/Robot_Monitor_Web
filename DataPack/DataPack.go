package DataPack

const MAX_PACKET_NUM = 5

const (
	ACT_READ   = 0x01
	ACT_WRITE  = 0x03
	ACT_UNREAD = 0x05
)

var TypeLen = map[string]int{
	"uint8_t":  1,
	"uint16_t": 2,
	"uint32_t": 4,
	"uint64_t": 8,
	"int8_t":   1,
	"int16_t":  2,
	"int32_t":  4,
	"int64_t":  8,
	"int":      4,
	"float":    4,
	"double":   8,
}

type JsonSerialPort struct {
	Ports []string
}

type DataFromWeb_t struct {
	Board uint8
	Name  string
	Act   uint8
	Type  string
	Addr  uint32
	Data  float64
}

type DataToRead_t struct {
	Board uint8
	Name  string
	Type  string
	Addr  uint32
	Data  float64
}

type DataToChat_t struct {
	Board uint8
	Name  string
	Data  float64
}

type DataToChat struct {
	DataPack []DataToChat_t
}

func (p *DataToRead_t) GetWebData(d *DataFromWeb_t) {
	p.Board = d.Board
	p.Name = d.Name
	p.Type = d.Type
	p.Addr = d.Addr
	p.Data = d.Data
}

var DataToRead []DataToRead_t
