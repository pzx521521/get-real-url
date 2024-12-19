package main

import (
	"fmt"
	webview "github.com/webview/webview_go"
	"os"
	"strings"
	"sync"
)

var urlMap map[string]string
var uri string
var wg sync.WaitGroup

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("can't find file:%v\n", err)
		return
	}
	urls := strings.Split(string(file), "\n")
	m := getRealUrl(urls)
	var wTexts []string
	for k, v := range m {
		wTexts = append(wTexts, fmt.Sprintf("%v\t%v\n", k, v))
	}
	err = os.WriteFile("output.txt", []byte(strings.Join(wTexts, "")), os.ModePerm)
	if err != nil {
		fmt.Printf("can't write file:%v\n", err)
		return
	}
}
func getRealUrl(urls []string) map[string]string {
	w := webview.New(false)
	w.Init(`setInterval(() => {
					window.log(window.location.href);
				}, 1000);`)
	err := w.Bind("log", jsLog)
	if err != nil {
		return nil
	}
	urlMap = map[string]string{}
	w.SetTitle("Bind Example")
	w.SetSize(480*2, 320*2, webview.HintNone)
	go func() {
		for _, url := range urls {
			if url == "" {
				continue
			}
			wg.Add(1)
			uri = url
			w.Dispatch(func() {
				w.Navigate(url)
			})
			wg.Wait()
		}
		w.Dispatch(func() {
			w.Terminate()
		})
	}()
	w.Run()
	return urlMap
}
func jsLog(logText string) {
	if urlMap[uri] == logText {
		fmt.Printf("确认%s的链接是%s\n", uri, logText)
		wg.Done()
		return
	}
	fmt.Printf("目前%s的链接是%s\n", uri, logText)
	urlMap[uri] = logText

}
