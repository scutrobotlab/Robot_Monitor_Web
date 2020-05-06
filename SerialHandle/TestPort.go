package serialhandle

import (
	"errors"
	"log"
	"math"
	"time"

	"go.bug.st/serial.v1"

	datapack "www.scut-robotlab.cn/git/M3chD09/Robot_Monitor_Web/DataPack"
)

type testPort struct {
	readingAddresses []uint32
	createdTime      time.Time
}

func newTestPort() serial.Port {
	return &testPort{
		readingAddresses: []uint32{},
		createdTime:      time.Now(),
	}
}

func (tp *testPort) SetMode(mode *serial.Mode) error { return nil }

func testValue(x float64, addr uint32) float64 {
	waveform := addr >> 24
	freq := (float64((addr>>8)&0xFF) - 0x80) / 16.0
	freq = math.Exp(freq)
	phase := float64(addr&0xFF) / 0xFF
	amplitude := float64((addr>>16)&0xFF) / 16.0
	amplitude = math.Exp(amplitude)
	currentPhase := x*freq + phase
	var scale float64
	switch waveform {
	case 1: // square
		if currentPhase-math.Floor(currentPhase) < 0.5 {
			scale = 1.0
		} else {
			scale = -1.0
		}
	case 0: // sin
		fallthrough
	default:
		scale = math.Sin(currentPhase * 2 * math.Pi)
	}
	y := scale * amplitude
	log.Printf("Test port: Address: 0x%08X %.5f => %.5f\n", addr, x, y)
	return y
}

func (tp *testPort) Read(p []byte) (n int, err error) {
	addresses := tp.readingAddresses
	maxNumPack := len(p) / 16
	if len(addresses) > maxNumPack {
		addresses = addresses[:maxNumPack]
	}

	for i, addr := range addresses {
		s := p[16*i : 16*(i+1)]
		s[0] = 1 // board
		s[1] = 2 // ???
		s[2] = 8 // typeLen
		copy(s[3:7], datapack.AnyToBytes(addr))
		x := time.Now().Sub(tp.createdTime).Seconds()
		y := testValue(x, addr)
		copy(s[7:15], datapack.AnyToBytes(y))
		s[15] = '\n'
	}
	return len(addresses) * 16, nil
}

func (tp *testPort) Write(p []byte) (n int, err error) {
	if len(p) != 16 {
		return 0, errors.New("Invalid len")
	}
	if p[len(p)-1] != '\n' {
		return 0, errors.New("Invalid package")
	}

	board := p[0]
	if board != 1 {
		return 0, errors.New("Invalid board")
	}
	act := p[1]
	typeLen := p[2]
	if typeLen != 8 {
		return 0, errors.New("Unsupported typeLen")
	}
	address := datapack.BytesToUint32(p[3:7])

	switch act {
	case datapack.ACT_SUBSCRIBE:
		tp.readingAddresses = append(tp.readingAddresses, address)
		log.Printf("Adding address: %08X\n", address)
		break
	case datapack.ACT_UNSUBSCRIBE:
		var newAddresses []uint32
		for _, addr := range tp.readingAddresses {
			if addr != address {
				newAddresses = append(newAddresses, addr)
			}
		}
		tp.readingAddresses = newAddresses
		break
	default:
		return 0, errors.New("Invalid act")
	}

	return 16, nil
}

func (tp *testPort) ResetInputBuffer() error { return nil }

func (tp *testPort) ResetOutputBuffer() error { return nil }

func (tp *testPort) SetDTR(dtr bool) error { return errors.New("Not supported") }

func (tp *testPort) SetRTS(rts bool) error { return errors.New("Not supported") }

func (tp *testPort) GetModemStatusBits() (*serial.ModemStatusBits, error) {
	return nil, errors.New("Not supported")
}

func (tp *testPort) Close() error { return nil }
