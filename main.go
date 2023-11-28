package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

func main() {
	r := gin.Default()

	// root
	r.GET("", func(c *gin.Context) {
		c.String(http.StatusOK, "localhost:5649/go")
	})

	// LTE Router의 상태를 전송한다.
	r.GET("/go", func(c *gin.Context) {
		target := "192.168.10.254"
		if pingCheck(target) {
			c.String(http.StatusOK, "Alive")
		} else {
			c.String(http.StatusInternalServerError, "Down")
		}
	})

	r.Run(":5649") // 서버가 실행 되고 0.0.0.0:5649 에서 요청을 기다립니다.
}

func pingCheck(target string) bool {
	out, err := exec.Command("ping", target, "-c 3").Output()
	if err != nil {
		log.Println("ERR: ", err)
	}
	log.Println(string(out))
	if strings.Contains(string(out), "Destination Host Unrechable") || strings.Contains(string(out), "errors") {
		log.Println("Down")
		return false
	}
	if strings.Contains(string(out), "0% packet loss") {
		n := strings.Index(string(out), "0% packet loss")
		if string(out[n-1]) == "0" && string(out[n-2]) == "1" {
			log.Println("Down")
			return false
		}
	} else {
		log.Println("Down")
		return false
	}

	log.Println("Alive")
	return true
}
