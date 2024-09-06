package smbus

import (
	"fmt"
	"log"
	"strings"
)

func (d *smBus) WriteRegister(register uint8, buf []byte) error {
	debug("write", buf)
	return d.conn.WriteBlockData(d.i2cAddr, register, buf)
}

func (d *smBus) WriteRegisterUint8(register, value uint8) error {
	return d.WriteRegister(register, []byte{value})
}

func debug(act string, buf []byte) {
	var s []string
	for _, v := range buf {
		s = append(s, fmt.Sprintf("0x%02X", v))
	}
	log.Printf("%s: [%s]\n", act, strings.Join(s, ","))
}
