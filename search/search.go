package search

import (
	"strconv"
	"github.com/PuerkitoBio/goquery"
	"strings"
	"net/url"
	"math"
)

func CalcScore(num int64, numArr []int64) float64 {
	t1 := float64(num) / float64(numArr[0])
	t2 := float64(num) / float64(numArr[1])
	score1 := math.Log2(1.0 + t1 + t2)
	//score2 := math.Log2(1.0 + t1) + math.Log2(1.0 + t2)
	return score1
}

func SearchInBaidu(searchStr string) (int64, error) {
	key := url.QueryEscape(searchStr)
	url := "http://www.baidu.com/baidu?&ie=utf-8&word=" + key
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return 0, err
	}
	tmpRet := doc.Find(".nums").Text()
	tmpRet = strings.Trim(strings.Split(tmpRet,"约")[1], "个")
	resultStr := strings.Join(strings.Split(tmpRet,","),"")
	result, _ := strconv.ParseInt(resultStr,10,64)
	return result, nil
}

func SearchInSogou(searchStr string) (int64, error) {
	key := url.QueryEscape(searchStr)
	url := "https://www.sogou.com/web?query=" + key + "&ie=utf8"
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return 0, err
	}
	tmpRet := doc.Find(".num-tips").Text()
	tmpRet = strings.Trim(strings.Split(tmpRet,"约")[1], "条相关结果")
	resultStr := strings.Join(strings.Split(tmpRet,","),"")
	result, _ := strconv.ParseInt(resultStr,10,64)
	return result, nil
}

func SearchIn360(searchStr string) (int64, error) {
	key := url.QueryEscape(searchStr)
	url := "https://www.so.com/s?q=" + key
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return 0, err
	}
	tmpRet := doc.Find(".nums").Text()
	tmpRet = strings.Trim(strings.Split(tmpRet,"约")[1], "个")
	resultStr := strings.Join(strings.Split(tmpRet,","),"")
	result, _ := strconv.ParseInt(resultStr,10,64)
	return result, nil
}

//func SearchInBing(searchStr string) (int64, error) {
//	key := url.QueryEscape(searchStr)
//	url := "https://cn.bing.com/search?q=" + key
//	doc, err := goquery.NewDocument(url)
//	if err != nil {
//		return 0, err
//	}
//	tmpRet := doc.Find(".sb_count").Text()
//	tmpRet = strings.Trim(tmpRet, " 条结果")
//	resultStr := strings.Join(strings.Split(tmpRet,","),"")
//	result, _ := strconv.ParseInt(resultStr,10,64)
//	return result, nil
//}

func SearchInChinaSo(searchStr string) (int64, error) {
	key := url.QueryEscape(searchStr)
	url := "http://www.chinaso.com/search/pagesearch.htm?q=" + key
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return 0, err
	}
	tmpRet := doc.Find(".pageTotal").Text()
	tmpRet = strings.Trim(strings.Split(tmpRet,"约")[1], "条相关结果")
	resultStr := strings.Join(strings.Split(tmpRet,","),"")
	result, _ := strconv.ParseInt(resultStr,10,64)
	return result, nil
}

func SearchFromAll(engine string, searchSlice []string) ([]string, string, int, float64) {
	retBox := make([]string, 0)
	allRetBox := make([]int64, 0)
	allRetBoxEptMax := make([]int64, 0)
	var (
		tmpNum int64
		retNum int64
		maxIdx int
		score float64
	)
	for idx, item := range searchSlice {
		tag := strconv.Itoa(idx + 1) + " -> "
		switch engine {
		case "Baidu":
			retNum, _ = SearchInBaidu(item)
			allRetBox = append(allRetBox, retNum)
		case  "Sogou" :
			retNum, _ = SearchInSogou(item)
			allRetBox = append(allRetBox, retNum)
		case "360" :
			retNum, _ = SearchIn360(item)
			allRetBox = append(allRetBox, retNum)
		//case "Bing":
		//	retNum, _ = SearchInBing(item)
		//	allRetBox = append(allRetBox, retNum)
		case "ChinaSo":
			retNum, _ = SearchInChinaSo(item)
			allRetBox = append(allRetBox, retNum)
		default:
			retNum, _ = SearchInBaidu(item)
			allRetBox = append(allRetBox, retNum)
		}
		if retNum > tmpNum {
			tmpNum = retNum
			maxIdx = idx
		}
		qAndA := tag + item + " :" + strconv.FormatInt(retNum,10)
		retBox = append(retBox, qAndA)
	}
	for i, v := range allRetBox {
		if i != maxIdx {
			if v == 0 {
				allRetBoxEptMax = append(allRetBoxEptMax, 1)
			}
			allRetBoxEptMax = append(allRetBoxEptMax, v)
		}
	}
	score = CalcScore(tmpNum, allRetBoxEptMax)
	choice := "由于第" + strconv.Itoa(maxIdx + 1) + "项结果最多，" + "选择" + strconv.Itoa(maxIdx + 1)
	return retBox, choice, maxIdx + 1, score
}

