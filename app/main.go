package main

import (
	"fmt"
	"net"
	"strings"

	"github.com/codecrafters-io/dns-server-starter-go/util"
)

func main() {
	udpAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:2053")
	if err != nil {
		fmt.Println("Failed to resolve UDP address:", err)
		return
	}

	udpConn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		fmt.Println("Failed to bind to address:", err)
		return
	}
	defer udpConn.Close()

	buf := make([]byte, 512)

	for {
		size, source, err := udpConn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("Error receiving data:", err)
			break
		}

		receivedData := string(buf[:size])
		fmt.Printf("Received %d bytes from %s: %s\n", size, source, receivedData)

		// fmt.Printf("Received buffer: %x\n", buf)
		fmt.Printf("Received buffer: %s\n", FormatBytes(buf))

		first12Bytes := buf[:12]
		var first12BytesArray [12]byte
		copy(first12BytesArray[:], first12Bytes)

		////////////////////
		// Header Section //
		////////////////////
		header, err := util.NewHeader(first12BytesArray)
		if err != nil {
			fmt.Println(err)
		}
		// fmt.Println("Received Buf in Header section")
		// fmt.Printf("%x\n", first12BytesArray)
		// fmt.Println("Received Header")
		// fmt.Println("qdcount: ", header.GetQdcount())
		// fmt.Println("ancount: ", header.GetAncount())
		// fmt.Println("nscount: ", header.GetNscount())
		// fmt.Println("arcount: ", header.GetArcount())
		// util.DebugHeader(header)
		// headerWithId, err := util.NewHeaderWithQdcountAndAncount(header, 1, 1)
		// if err != nil {
		// fmt.Println(err)
		// }
		answerHeader := util.Reply(header)
		// fmt.Println("Answeered Header")
		// fmt.Println("qdcount: ", answerHeader.GetQdcount())
		// fmt.Println("ancount: ", answerHeader.GetAncount())
		// fmt.Println("nscount: ", answerHeader.GetNscount())
		// fmt.Println("arcount: ", answerHeader.GetArcount())
		answerBytes := util.HeaderToBytes(answerHeader)

		response := answerBytes[:]

		//////////////////////
		// Question Section //
		//////////////////////
		numQuestion := answerHeader.GetQdcount()
		fmt.Println(numQuestion)
		questionBufByte, err := util.QuestionBytes(buf, int(numQuestion))
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("question buf")
		fmt.Printf("%x\n", questionBufByte)
		fmt.Printf("%c\n", questionBufByte)

		questions, err := util.NewQuestionsFromByte(questionBufByte, int(numQuestion))

		if err != nil {
			fmt.Println(err)
		}
		questionByte := util.QuestionToBytes(questions)
		response = append(response, questionByte...)

		////////////////////
		// Answer Section //
		////////////////////
		domainStrs, err := util.DomainsInQuestion(questions)
		if err != nil {
			fmt.Println(err)
		}

		var rrs []util.ResourceRecord
		for i := 0; i < len(domainStrs); i++ {
			rr, err := util.NewResourceRecord(domainStrs[i])
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("RR")
			fmt.Printf("%0x\n", rr)
			rrs = append(rrs, rr)
		}

		// fmt.Println("rrs")
		// fmt.Println(len(rrs))
		answerSection := util.NewAnswer(rrs)
		answerSectionBytes := util.AnswerToBytes(answerSection)

		response = append(response, answerSectionBytes...)

		///////////
		// Debug //
		///////////
		// fmt.Printf("Response buffer: %x\n", response)
		fmt.Printf("Response buffer: %s\n", FormatBytes(response))

		_, err = udpConn.WriteToUDP(response, source)
		if err != nil {
			fmt.Println("Failed to send response:", err)
		}
	}
}

func FormatBytes(data []byte) string {
	var sb strings.Builder
	for i, b := range data {
		if i > 0 && i%2 == 0 {
			sb.WriteString(" ")
		}
		sb.WriteString(fmt.Sprintf("%02x", b))
	}
	return sb.String()
}
