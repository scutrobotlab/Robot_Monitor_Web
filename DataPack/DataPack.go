package datapack

const MAX_PACKET_NUM = 5

const (
	_ = iota
	ACT_SUBSCRIBE
	ACT_SUBSCRIBERETURN
	ACT_UNSUBSCRIBE
	ACT_UNSUBSCRIBERETURN
	ACT_READ
	ACT_READRETURN
	ACT_WRITE
	ACT_WRITERETURN
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

type VariableT struct {
	Board uint8
	Name  string
	Type  string
	Addr  uint32
	Data  float64
}

type jsonCurrentVariablesT struct {
	Variables []VariableT
}

var CurrentVariables jsonCurrentVariablesT
var ModVariables jsonCurrentVariablesT

type DataToChartT struct {
	Board uint8
	Name  string
	Data  float64
}

type DataToChart struct {
	DataPack []DataToChartT
}
