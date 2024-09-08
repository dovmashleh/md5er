package main

import (
	"crypto/md5"
	"fmt"
	"md5er/internal/servers/http2"
	"md5er/internal/servers/tcp"
	"sync"

	"io"
	mymd5 "md5er/internal/md5"
)

func main() {
	md5service := mymd5.New()
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		HttpServer := http2.New(md5service)
		HttpServer.Start()
		wg.Done()
	}()
	go func() {
		tcpServer := tcp.New(md5service)
		tcpServer.Start()
		wg.Done()
	}()
	wg.Wait()
}

func someTets() {
	test := "ZuluPapa234567890-frghjkl;dfgbhnm,dertyul;d5rftyuijodsauihfdasufrauigreijgvbifdajvipoafdsjgioraetguiorehtuiofafidsajfiopesaufiowe4qhasdipofs"
	md5er := mymd5.New()
	myres := md5er.AsByteArray([]byte(test))
	h := md5.New()
	io.WriteString(h, test)
	fmt.Printf("real: %x \n", h.Sum(nil))
	fmt.Printf("my: %x \n", myres)
}
