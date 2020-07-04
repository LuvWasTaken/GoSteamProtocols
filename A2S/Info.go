package A2S

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"reflect"
	"time"
)

func nullTerminatedString(buffer *bytes.Buffer) string {
	var data bytes.Buffer
	var char = buffer.Next(1)[0]
	for char != 0 {
		data.WriteString(string(char))
		char = buffer.Next(1)[0]
	}
	return data.String()
}

func QueryInfo(address Address) Info {
	var response = make([]byte, 4096)
	var request = make([]byte, 0)
	var serverData Info
	var conn, _ = net.Dial("udp", fmt.Sprintf("%s:%d", address.IP, address.Port))

	request = append(request, []byte{255, 255, 255, 255, 84}...)
	request = append(request, []byte("Source Engine Query")...)
	request = append(request, byte(0))
	conn.Write(request)
	conn.SetDeadline(time.Now().Add(time.Second * 5))

	var n, err = conn.Read(response)
	if err == nil {
		var buffer = bytes.NewBuffer(response[:n])
		if reflect.DeepEqual(buffer.Next(4), []byte{255, 255, 255, 255}) {
			serverData = Info{
				Header:     buffer.Next(1)[0],
				Protocol:   buffer.Next(1)[0],
				Name:       nullTerminatedString(buffer),
				GameMap:    nullTerminatedString(buffer),
				Folder:     nullTerminatedString(buffer),
				Game:       nullTerminatedString(buffer),
				ID:         binary.LittleEndian.Uint16(buffer.Next(2)),
				Players:    buffer.Next(1)[0],
				MaxPlayers: buffer.Next(1)[0],
				Bots:       buffer.Next(1)[0],
				ServerType: buffer.Next(1)[0],
				Enviroment: buffer.Next(1)[0],
				Visibility: buffer.Next(1)[0],
				Vac:        buffer.Next(1)[0],
				Version:    nullTerminatedString(buffer),
				EDF:        buffer.Next(1)[0],
				GamePort:   binary.LittleEndian.Uint16(buffer.Next(2))}
		}

	}
	conn.Close()
	return serverData
}
