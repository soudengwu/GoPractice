package selector

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecondTimeout = 10 * time.Second

func Racer(url1, url2 string) (winner string, err error) {

	return ConfigurableRacer(url1, url2, tenSecondTimeout)

}

func ConfigurableRacer(url1 string, url2 string, timeout time.Duration) (winner string, err error) {
	select {
	case <-ping(url1):
		return url1, nil
	case <-ping(url2):
		return url2, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", url1, url2)
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}

// func measureResponseTime(url string) time.Duration {
// 	start := time.Now()
// 	http.Get(url)
// 	urlDuration := time.Since(start)
// 	return urlDuration
// }

func nillChannel() {
	ch1 := make(chan bool, 1)
	ch1 <- true
	fmt.Println("tada")

	var ch2 chan bool
	fmt.Println("the zero value of a channel is nil", ch2)

	// needs to be non blocking coz you cant send to a nil channel
	go func() {
		ch2 <- true
		fmt.Println("you'll never see this because the above will block forever")
	}()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println(`you cant send to a nil channel, so this "wins"`)
	case <-ch2:
		fmt.Println("you'll never see this")
	}
}
