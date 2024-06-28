package main

import (
	"fmt"
	"net"

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

		fmt.Printf("Received buffer: %x\n", buf)
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
		headerWithId, err := util.NewHeaderWithQdcountAndAncount(header, 1, 1)
		if err != nil {
			fmt.Println(err)
		}
		answerHeader := util.Reply(headerWithId)
		answerBytes := util.HeaderToBytes(answerHeader)

		response := answerBytes[:]

		//////////////////////
		// Question Section //
		//////////////////////
		questionBufByte, err := util.QuestionBytes(buf)
		if err != nil {
			fmt.Println(err)
		}
		question, err := util.NewQuestionFromByte(questionBufByte)

		if err != nil {
			fmt.Println(err)
		}
		questionByte := util.QuestionToBytes(question)
		response = append(response, questionByte...)

		////////////////////
		// Answer Section //
		////////////////////
		domainStr, err := util.DomainInQuestion(question)
		if err != nil {
			fmt.Println(err)
		}
		rr, err := util.NewResourceRecord(domainStr)
		if err != nil {
			fmt.Println(err)
		}

		answerSection := util.NewAnswer(rr)
		answerSectionBytes := util.AnswerToBytes(answerSection)

		response = append(response, answerSectionBytes...)

		///////////
		// Debug //
		///////////
		fmt.Printf("Response buffer: %x\n", response)

		_, err = udpConn.WriteToUDP(response, source)
		if err != nil {
			fmt.Println("Failed to send response:", err)
		}
	}
}
