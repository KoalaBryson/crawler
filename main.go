package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "https://movie.douban.com/cinema/nowplaying/hangzhou/"

	headers := make(http.Header)
	headers.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36 Edg/117.0.2045.35")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("build request error:%v", err)
		return
	}
	req.Header = headers

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("fetch url error:%v", err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error status code:%v", resp.StatusCode)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read content failed:%v", err)
		return
	}
	fmt.Println("body:", string(body))
}
