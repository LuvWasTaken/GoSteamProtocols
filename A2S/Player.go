package A2S

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"reflect"
	"time"
)

func QueryPlayer(address Address) []byte {

	var response = make([]byte, 4096)
	var request = make([]byte, 0)
	var conn, _ = net.Dial("udp", fmt.Sprintf("%s:%d", address.IP, address.Port))

	request = append(request, []byte{0xFF, 0xFF, 0xFF, 0xFF, 0x55}...)
	request = append(request, GetChallengeNumber(address)...)

	conn.Write(request)
	conn.SetDeadline(time.Now().Add(time.Second * 5))
	var n, err = conn.Read(response)

	if err == nil {
		var buffer = bytes.NewBuffer(response[:n])
		if reflect.DeepEqual(buffer.Next(5), []byte{0xFF, 0xFF, 0xFF, 0xFF, 0x44}) {
			buffer.Next(1)
			for i := 0; i < buffer.Len(); i++ {
				var index = buffer.Next(1)[0]
				var name = nullTerminatedString(buffer)
				buffer.Next(4)
				fmt.Println(binary.LittleEndian.Uint32(buffer.Next(4)))
				fmt.Println(index)
				fmt.Println(name)

			}

		}
	}
	return nil

}
