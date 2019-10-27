package DataPack

import (
	"bytes"
	"encoding/binary"
	"math"
)

func AnyToBytes(i interface{}) []byte {
	buf := bytes.NewBuffer([]byte{})
	binary.Write(buf, binary.LittleEndian, i)
	return buf.Bytes()
}

func BytesToInt8(i []byte) int8 {
	var o int8
	buf := bytes.NewReader(i)
	binary.Read(buf, binary.LittleEndian, &o)
	return o
}

func BytesToInt16(i []byte) int16 {
	var o int16
	buf := bytes.NewReader(i)
	binary.Read(buf, binary.LittleEndian, &o)
	return o
}

func BytesToInt32(i []byte) int32 {
	var o int32
	buf := bytes.NewReader(i)
	binary.Read(buf, binary.LittleEndian, &o)
	return o
}

func BytesToInt64(i []byte) int64 {
	var o int64
	buf := bytes.NewReader(i)
	binary.Read(buf, binary.LittleEndian, &o)
	return o
}

func BytesToUint8(i []byte) uint8 {
	var o uint8
	buf := bytes.NewReader(i)
	binary.Read(buf, binary.LittleEndian, &o)
	return o
}

func BytesToUint16(i []byte) uint16 {
	var o uint16
	buf := bytes.NewReader(i)
	binary.Read(buf, binary.LittleEndian, &o)
	return o
}

func BytesToUint32(i []byte) uint32 {
	var o uint32
	buf := bytes.NewReader(i)
	binary.Read(buf, binary.LittleEndian, &o)
	return o
}

func BytesToUint64(i []byte) uint64 {
	var o uint64
	buf := bytes.NewReader(i)
	binary.Read(buf, binary.LittleEndian, &o)
	return o
}

func BytesToFloat32(bytes []byte) float32 {
	bits := binary.LittleEndian.Uint32(bytes)
	return math.Float32frombits(bits)
}

func BytesToFloat64(bytes []byte) float64 {
	bits := binary.LittleEndian.Uint64(bytes)
	return math.Float64frombits(bits)
}
