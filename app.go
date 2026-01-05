package main

import (
	"encoding/binary"
	"fmt"
	"hash/crc32"
	"math/rand"
	"net"
	"time"
)

func main() {
	c := make(chan Data)
	go Gamepad(c)
	//return
	serverAddr, _ := net.ResolveUDPAddr("udp", "192.168.0.133:26760")
	conn, err := net.DialUDP("udp", nil, serverAddr)
	if err != nil {
		panic(err)
	}
	clientid := rand.Uint32()
	defer conn.Close()
	send1 := []byte{}
	// Header
	send1 = append(send1, []byte("DSUC")...)

	send1 = binary.LittleEndian.AppendUint16(send1, 1001) // protocol version
	send1 = binary.LittleEndian.AppendUint16(send1, 12)   // packet length
	send1 = binary.LittleEndian.AppendUint32(send1, 0)    // zero CRC32 field
	//send1 = binary.LittleEndian.AppendUint32(send1, 1932838430)
	send1 = binary.LittleEndian.AppendUint32(send1, clientid) // client id

	// Message
	send1 = binary.LittleEndian.AppendUint32(send1, 0x100001) // Event type
	var controllerCount int32 = 4
	send1 = binary.LittleEndian.AppendUint32(send1, uint32(controllerCount)) // controller count
	send1 = binary.LittleEndian.AppendUint32(send1, 50462976)                // idk
	// calculate crc32
	crc := crc32.ChecksumIEEE(send1)
	binary.LittleEndian.PutUint32(send1[8:12], crc)
	fmt.Printf("% X\n", send1)
	conn.Write(send1)
	//header
	send2 := []byte{}
	send2 = append(send2, []byte("DSUC")...)
	send2 = binary.LittleEndian.AppendUint16(send2, 1001) // protocol version
	send2 = binary.LittleEndian.AppendUint16(send2, 12)   // packet length
	send2 = binary.LittleEndian.AppendUint32(send2, 0)    // zero CRC32 field

	//send1 = binary.LittleEndian.AppendUint32(send1, 1932838430)
	send2 = binary.LittleEndian.AppendUint32(send2, clientid) // client id
	//Message
	send2 = binary.LittleEndian.AppendUint32(send2, 0x100002)
	send2 = binary.LittleEndian.AppendUint64(send2, 0)

	crc1 := crc32.ChecksumIEEE(send2)
	binary.LittleEndian.PutUint32(send2[8:12], crc1)
	fmt.Printf("% X\n", send2)
	conn.Write(send2)
	buf := make([]byte, 1024)
	go func() {
		for {
			time.Sleep(100 * time.Millisecond)
			conn.Write(send1)
			conn.Write(send2)
		}
	}()
	for {
		n, err := conn.Read(buf)

		if err != nil {
			fmt.Println("Ошибка чтения:", err)
			continue
		}
		data := buf[:n]
		if n < 50 {
			continue
		}

		fmt.Printf("\r %08b", data[37])
		fmt.Printf(" %08b", data[36])
		fmt.Printf(" %3d ", data[40])
		fmt.Printf(" %3d ", data[41])
		fmt.Printf(" %3d ", data[42])
		fmt.Printf(" %3d ", data[43])
		p := fmt.Sprintf("%08b", data[37])
		o := fmt.Sprintf("%08b", data[36])
		if p[0] == '1' {
			print("X")
		} else {
			print(" ")
		}
		if p[1] == '1' {
			print("A")
		} else {
			print(" ")
		}
		if p[2] == '1' {
			print("B")
		} else {
			print(" ")
		}
		if p[3] == '1' {
			print("Y")
		} else {
			print(" ")
		}
		if p[4] == '1' {
			print("RB")
		} else {
			print("  ")
		}
		if p[5] == '1' {
			print("LB")
		} else {
			print("  ")
		}
		if p[6] == '1' {
			print("RT")
		} else {
			print("  ")
		}
		if p[7] == '1' {
			print("LT")
		} else {
			print("  ")
		}
		c <- Data{
			leftX:         int(data[40]),
			leftY:         int(data[41]),
			rightX:        int(data[42]),
			rightY:        int(data[43]),
			FirstButtons:  p,
			SecondButtons: o,
		}
		time.Sleep(10 * time.Millisecond)

	}
	//send2 =

	//fmt.Printf("%% X:  % X\n", send1)
}
