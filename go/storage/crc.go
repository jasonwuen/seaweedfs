package storage

import (
	"encoding/binary"
	"fmt"

	"github.com/klauspost/crc32"
)

var table = crc32.MakeTable(crc32.Castagnoli)

type CRC uint32

func NewCRC(b []byte) CRC {
	return CRC(0).Update(b)
}

func (c CRC) Update(b []byte) CRC {
	return CRC(crc32.Update(uint32(c), table, b))
}

func (c CRC) Value() uint32 {
	return uint32(c>>15|c<<17) + 0xa282ead8
}

func (n *Needle) Etag() string {
	bits := make([]byte, 4)
	binary.BigEndian.PutUint32(bits, uint32(n.Checksum))
	return fmt.Sprintf("\"%x\"", bits)
}
