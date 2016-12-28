package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var format = "http://wpc.8c48.edgecastcdn.net/038C48/SV/480/DXDJPN0021/DXDJPN0021-480-%dK.mp4.m3u8?26snq6A8oNgg3mLxOXhBtB2vKXT6TivKr5GrS60l4R4M6ELQQH-KAWNvtFKATXP32oLOxReeDepJRBqaUao4qrs_l0liqj-4850Pm82obMM4C5voNYPgtVWYy2f_0dGLgkbhrs0P2ehd7JDXos4lIOxn"

func main() {
	urls := make(chan string, 1)
	go func() {
		for i := 1500; i <= 4000; i++ {
			log.Printf("Checking %d", i)
			url := fmt.Sprintf(format, i)
			urls <- url
		}
		close(urls)
	}()
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			for url := range urls {
				resp, err := http.Get(url)
				if err != nil {
					log.Println(err)
					continue
				}
				if resp.StatusCode == 200 {
					log.Printf("Found %dK: %s", i, url)
				}
				resp.Body.Close()
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

// 750, 1000, 1500, 2000, 2500, 3500, 4000
