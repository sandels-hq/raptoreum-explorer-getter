package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

var delay int64

func main() {
	flag.Int64Var(&delay, "delay", 1, "delay in seconds")
	flag.Parse()

	response, err := http.Get("https://explorer.raptoreum.com/api/getblockcount")
	if err != nil {
		log.Fatalln("getBlockCount", err)
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln("body read", err)
	}

	block, err := strconv.Atoi(string(body))
	if err != nil {
		log.Fatalln("body convert", err)
	}

	for i := block; i >= 0; i-- {
		url := "https://explorer.raptoreum.com/block-height/" + strconv.Itoa(i)
		println(url)

		stop := false
		done := make(chan bool)

		go func() {
			print("slo")
			for {
				if stop {
					break
				}
				print("o")
				time.Sleep(100 * time.Millisecond)
			}
			done <- true
		}()

		startedAt := time.Now()
		response, _ := http.Get(url)

		stop = true
		<-done

		fmt.Printf("w, %.1fs\n", time.Since(startedAt).Seconds())
		time.Sleep(time.Duration(delay) * time.Second)

		if response.StatusCode != 200 {
			log.Fatalln("not 200", response.StatusCode)
		}
	}
}
