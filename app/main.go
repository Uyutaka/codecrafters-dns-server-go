package main

import (
	"fmt"
	"net"
	"os"

	"github.com/codecrafters-io/dns-server-starter-go/util"
)

func main() {
	args := os.Args
	var resolver util.Resolver
	if len(args) == 3 {
		host, port, err := net.SplitHostPort(args[2])
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		ip := net.ParseIP(host)
		if ip == nil {
			fmt.Println("Invalid IP address")
			return
		}
		resolver = util.NewResolver(ip, port)
		fmt.Println(resolver)
	}

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
		fmt.Printf("Received buffer: %s\n", util.FormatBytes(buf))

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
		responseHeader := util.Reply(header)
		responseHeaderBytes := util.HeaderToBytes(responseHeader)

		response := responseHeaderBytes[:]

		//////////////////////
		// Question Section //
		//////////////////////
		numQuestion := responseHeader.GetQdcount()
		questionBufByte, err := util.QuestionBytes(buf, int(numQuestion))
		if err != nil {
			fmt.Println(err)
		}

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
			var rr util.ResourceRecord
			if util.IsResolver(resolver) {
				rr, err = util.ForwardDNSRequest(resolver, header, questions[i], domainStrs[i])
			} else {
				rr, err = util.NewResourceRecord(domainStrs[i])
			}
			if err != nil {
				fmt.Println(err)
			}

			rrs = append(rrs, rr)
		}

		answerSection := util.NewAnswer(rrs)
		answerSectionBytes := util.AnswerToBytes(answerSection)

		response = append(response, answerSectionBytes...)

		///////////
		// Debug //
		///////////
		// fmt.Printf("Response buffer: %x\n", response)
		fmt.Printf("Response buffer: %s\n", util.FormatBytes(response))

		_, err = udpConn.WriteToUDP(response, source)
		if err != nil {
			fmt.Println("Failed to send response:", err)
		}
	}
}
