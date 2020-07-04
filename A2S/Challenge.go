package A2S

import (
	"bytes"
	"fmt"
	"net"
	"reflect"
	"time"
)

func GetChallengeNumber(address Address) []byte {

	var response = make([]byte, 4096)
	var conn, _ = net.Dial("udp", fmt.Sprintf("%s:%d", address.IP, address.Port))
	conn.Write([]byte{0xFF, 0xFF, 0xFF, 0xFF, 0x55, 0xFF, 0xFF, 0xFF, 0xFF})
	conn.SetDeadline(time.Now().Add(time.Second * 5))
	var n, err = conn.Read(response)

	if err == nil {
		var buffer = bytes.NewBuffer(response[:n])
		if reflect.DeepEqual(buffer.Next(5), []byte{0xFF, 0xFF, 0xFF, 0xFF, 0x41}) {
			return buffer.Next(4)

		}
	}
	return nil

}
