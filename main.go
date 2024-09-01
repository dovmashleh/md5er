package main

import (
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	mymd5 "md5er/internal/md5"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}

func someTets() {
	test := "ZuluPapa234567890-frghjkl;dfgbhnm,dertyul;d5rftyuijodsauihfdasufrauigreijgvbifdajvipoafdsjgioraetguiorehtuiofafidsajfiopesaufiowe4qhasdipofs"
	myres := mymd5.AsByteArray([]byte(test))
	h := md5.New()
	io.WriteString(h, test)
	fmt.Printf("real: %x \n", h.Sum(nil))
	fmt.Printf("my: %x \n", myres)
}
