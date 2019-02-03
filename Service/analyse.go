package anaylse

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type UrlData struct {
	Url        string
	AvgResTime int
	Quantity   int
}
type urlDataArr []UrlData

func (s urlDataArr) Len() int {
	return len(s)
}
func (s urlDataArr) Less(i, j int) bool { //decreasing
	return s[i].AvgResTime > s[j].AvgResTime
}
func (s urlDataArr) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func Analyze_requests(path string, returnSize int) []string {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	size, statErr := f.Stat()
	if statErr != nil {
		panic(statErr)
	}

	defer f.Close()

	rd := bufio.NewReaderSize(f, int(size.Size()>>2))
	urlmap := make(map[string]UrlData)
	var urlData urlDataArr

	// tmpurl := make([]UrlData, 0)
	// sort.Ints(res)
	line := 0 //to check which line
	for {
		line++
		lineData, _, err := rd.ReadLine()

		if err != nil || err == io.EOF {
			fmt.Println(err)
			break
		}
		fmt.Println(string(lineData))
		s := strings.Split(string(lineData), ",")
		// fmt.Println(s)
		if len(s) != 3 {
			fmt.Println("err format problem:%v on line %d", s, line)
			break
		}

		//filter URL ending with .gif and with another method
		skip, url := filterURL(s[0])
		// fmt.Println("url", url)
		if skip {
			continue
		}
		//filter status without 200
		ok := filterStatus(s[2])
		if !ok {
			continue
		}
		//trim response time
		data, _ := strconv.ParseFloat(strings.TrimSuffix(s[1], "s"), 32)
		resTime := int(data * 1000)
		urlExists, ok := urlmap[url]
		if !ok { //new one
			var tmp UrlData
			tmp.Url = url
			tmp.AvgResTime = resTime
			tmp.Quantity = 1

			urlmap[url] = tmp

		} else { //existing
			tmp := urlExists
			tmp.AvgResTime = (tmp.AvgResTime*tmp.Quantity + resTime) / (tmp.Quantity + 1)
			tmp.Quantity = tmp.Quantity + 1
			urlmap[url] = tmp
		}
	}
	if line == 0 {
		res := make([]string, 0)
		return res
	}
	for _, value := range urlmap {
		urlData = append(urlData, value)
	}

	sort.Sort(urlData) //decresingly sorting

	if len(urlData) < returnSize {
		returnSize = len(urlData) //maximum =returnsize
	}
	res := make([]string, returnSize)
	for i := 0; i < returnSize; i++ {
		res[i] = urlData[i].Url
	}

	return res
}

func filterURL(s string) (bool, string) {
	allowSuffix := ".gif"
	allowMethod := "get"
	url := strings.ToLower(s)

	return strings.HasSuffix(url, allowSuffix) || !strings.HasPrefix(url, allowMethod), url
}
func filterStatus(status string) bool {

	statusReg := regexp.MustCompile(`[0-9]+`)
	statusData := statusReg.FindAllString(status, -1)[0]

	if strings.Compare(statusData, "200") != 0 {
		fmt.Println("status:", statusData)
		return false
	}
	return true
}
