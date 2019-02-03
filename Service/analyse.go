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

//URLData is for saving every distinct url avgResponse time struct
type URLData struct {
	URL        string
	AvgResTime int
	Quantity   int
}
type urlDataArr []URLData

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
	urlmap := make(map[string]URLData)
	var urlData urlDataArr

	line := 0 //to track which line occur problems
	for {
		line++
		lineData, err := rd.ReadString('\n')
		if err != nil || err == io.EOF {
			fmt.Println(err)
			break
		}
		//skip \r and \n
		lineData = strings.TrimSpace(lineData)
		if len(lineData) == 0 {
			fmt.Printf("skip this this line: %d", line)
			continue
		}
		s := strings.Split(string(lineData), ",")
		if len(s) != 3 {
			fmt.Printf("err format problem:%v on line %d", s, line)
			break
		}

		//filter URL ending with .gif and with another method
		skip, url := filterURL(s[0])

		if skip {
			continue
		}
		//filter status without 200
		ok := filterStatus(s[2])
		if !ok {
			continue
		}
		//filter response time
		resTime := filterResTime(s[1])
		urlExists, ok := urlmap[url]
		if !ok { //new one
			var tmp URLData
			tmp.URL = url
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
		res[i] = urlData[i].URL
	}

	return res
}

func filterURL(s string) (bool, string) {
	allowSuffix := ".gif"
	allowMethod := "get" //here assume that begin with "/" or "get" is GET method Request
	url := strings.ToLower(s)
	fmt.Println("1", url)
	url = strings.TrimSpace(strings.TrimPrefix(url, allowMethod)) //TrimSpace doesn't consider buffer with 0
	fmt.Println("2", url)
	skip := false
	if !strings.HasPrefix(url, "/") { //another method
		// url = subString(s, '/')
		skip = true
	}

	return strings.HasSuffix(url, allowSuffix) || skip, url
}

// no need to record another method
// func subString(s string, firstByte byte) string {
// 	pos := strings.IndexByte(s, firstByte)
// 	res := make([]byte, 0)
// 	if pos >= 0 {
// 		res = []byte(s)[pos:len(s)]
// 	}

// 	return string(res)
// }
func filterStatus(status string) bool {

	statusReg := regexp.MustCompile(`[0-9]+`)
	statusData := statusReg.FindAllString(status, -1)[0]

	if strings.Compare(statusData, "200") != 0 {
		fmt.Println("status:", statusData)
		return false
	}
	return true
}

func filterResTime(time string) int {
	time1 := strings.TrimSpace(time)
	data, _ := strconv.ParseFloat(strings.TrimSuffix(time1, "s"), 32)
	resTime := int(data * 1000)
	return resTime
}
