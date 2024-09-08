package main

import (
	"crypto/md5"
	"fmt"

	"io"
	mymd5 "md5er/internal/md5"
	"md5er/internal/servers/tcp"
)

func main() {
	server := tcp.New()
	server.Start()
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
