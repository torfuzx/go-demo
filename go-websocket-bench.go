package main

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("==\n==", "Go Websocket Bench", "\n==")

	host := "http://ws.kefu-dev.hotpu.cn/employee/ws"

	req, err := http.NewRequest("GET", host, nil)

	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	var standard WebsocketStandard = new(WebsocketStandard_13)

	standard.AddHeaders(req)

	//req.Header.Add("Origin", "go-websocket-bench")

	fmt.Println(req)

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	fmt.Println(resp)

	if standard.VerifyResponse(resp) != true {
		fmt.Println("==\n==", "Wrong accept from server", "\n==")
	}
}

type WebsocketStandard interface {
	AddHeaders(req *http.Request)
	VerifyResponse(resp *http.Response) bool
}

type WebsocketStandard_13 struct {
	Key string
}

func (s *WebsocketStandard_13) AddHeaders(req *http.Request) {
	req.Header.Add("Upgrade", "websocket")
	req.Header.Add("Connection", "Upgrade")
	req.Header.Add("Sec-WebSocket-Version", "13")

	buf := make([]byte, 16)
	_, err := io.ReadFull(rand.Reader, buf)

	if err != nil {
		panic(fmt.Errorf("ERROR", err))
	}

	s.Key = base64.StdEncoding.EncodeToString(buf)

	req.Header.Add("Sec-WebSocket-Key", s.Key)
}

func (s *WebsocketStandard_13) VerifyResponse(resp *http.Response) (b bool) {
	h := sha1.New()

	p := []byte(s.Key + "258EAFA5-E914-47DA-95CA-C5AB0DC85B11")

	_, err := h.Write(p)

	if err != nil {
		panic(fmt.Errorf("ERROR", err))
	}

	expected := base64.StdEncoding.EncodeToString(h.Sum(nil))

	return resp.Header.Get("Sec-WebSocket-Accept") == expected
}
