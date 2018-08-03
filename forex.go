package main

import (
	"fmt"
	"sync"
	"time"
	"github.com/go-resty/resty"
	"strings"
	"os"
	"path"
)

var (
	gmt, _ = time.LoadLocation("GMT")
)

func GetDukascopyLinks(symbol string, from, till time.Time) (list []string) {
	count := int(till.Sub(from) / time.Hour)
	list = make([]string, 0, count)
	for from.Before(till) {
		hour := from.Hour()
		year, month, day := from.Date()
		link := fmt.Sprintf(
			"http://datafeed.dukascopy.com/datafeed/%s/%04d/%02d/%02d/%02dh_ticks.bi5",
			symbol, year, month-1, day, hour,
		)
		list = append(list, link)
		from = from.Add(time.Hour)
	}
	return
}

func ParseTime(str string) time.Time {
	const layout = "2006.01.02"
	tm, err := time.Parse(layout, str)
	if err != nil {
		fmt.Println(err)
		return time.Time{}
	}
	return tm.In(gmt)
}

func worker(wg *sync.WaitGroup, linkChan <-chan string) {
	defer wg.Done()

	dataPath := "datafeed"

	cli := resty.New()
	cli.SetOutputDirectory(dataPath)

	for link := range linkChan {

		fileName := strings.TrimPrefix(link, "http://datafeed.dukascopy.com/datafeed/")

		fmt.Println(link)
		fmt.Println(fileName)

		if _, err := os.Stat(path.Join(dataPath, fileName)); err == nil {
			fmt.Println("skip")
			continue
		}

		resp, err := cli.R().SetOutput(fileName).Get(link)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println(resp.StatusCode())
	}

	fmt.Println("done1")
}

func main() {

	wg := sync.WaitGroup{}
	linkChan := make(chan string)
	workerCount := 8

	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go worker(&wg, linkChan)
	}

	from := ParseTime("2017.03.01")
	till := ParseTime("2018.02.05")

	list := GetDukascopyLinks("EURUSD", from, till)

	for _, link := range list {
		linkChan <- link
	}

	close(linkChan)

	wg.Wait()

	fmt.Println("done2")
}

