package result

import (
	"fmt"
	"sort"
	"strconv"
	"search_crawl/search"
)

type AnsStruct struct {
	option int
	score float64
}

func GetFinalRet(answerBox []AnsStruct) ([]AnsStruct, string) {
	var (
		finalRet AnsStruct
		sum1 float64
		sum2 float64
		sum3 float64
	)
	for _, v := range answerBox {
		switch v.option {
		case 1:
			sum1 += v.score
		case 2:
			sum2 += v.score
		case 3:
			sum3 += v.score
		}
	}
	scoreBox := []float64{sum1, sum2, sum3}
	sort.Float64s(scoreBox)
	maxScore := scoreBox[len(scoreBox) - 1]
	if maxScore == sum1 {
		finalRet.option = 1
		finalRet.score = sum1
	}
	if maxScore == sum2 {
		finalRet.option = 2
		finalRet.score = sum2
	}
	if maxScore == sum3 {
		finalRet.option = 3
		finalRet.score = sum3
	}
	finalRetBox := []AnsStruct{{1, sum1},{2, sum2},{3,sum3}}
	AnswerStr := "由于" + strconv.Itoa(finalRet.option) + "的综合得分最高，为：" + fmt.Sprintf("%0.5f",finalRet.score) + "\n所以选择" + strconv.Itoa(finalRet.option) + "."
	return finalRetBox, AnswerStr
}

func GetResult(searchSlice []string) {
	searchEngine := []string{"Baidu", "Sogou", "360", "Bing"}
	answerBox := []AnsStruct{}
	for _, item := range searchEngine {
		tmpBox := make([]string, 0)
		var (
			choice string
			ans int
			score float64
		)
		if item == "Baidu" {
			tmpBox, choice, ans, score = search.SearchFromAll(item, searchSlice)
			fmt.Println("————————————————————————————————————————————")
			fmt.Println("百度搜索结果如下：")
			for _, x := range tmpBox{
				fmt.Println(x)
			}
			fmt.Println(choice)
			tmpAns := AnsStruct{ans,score}
			answerBox = append(answerBox, tmpAns)
		}
		if item == "Sogou" {
			tmpBox, choice, ans, score = search.SearchFromAll(item, searchSlice)
			fmt.Println("————————————————————————————————————————————")
			fmt.Println("搜狗搜索结果如下：")
			for _, x := range tmpBox{
				fmt.Println(x)
			}
			fmt.Println(choice)
			tmpAns := AnsStruct{ans,score}
			answerBox = append(answerBox, tmpAns)
		}
		if item == "360" {
			tmpBox, choice, ans, score = search.SearchFromAll(item, searchSlice)
			fmt.Println("————————————————————————————————————————————")
			fmt.Println("360搜索结果如下：")
			for _, x := range tmpBox{
				fmt.Println(x)
			}
			fmt.Println(choice)
			tmpAns := AnsStruct{ans,score}
			answerBox = append(answerBox, tmpAns)
		}
		if item == "Bing" {
			tmpBox, choice, ans, score = search.SearchFromAll(item, searchSlice)
			fmt.Println("————————————————————————————————————————————")
			fmt.Println("必应搜索结果如下：")
			for _, x := range tmpBox{
				fmt.Println(x)
			}
			fmt.Println(choice)
			tmpAns := AnsStruct{ans,score}
			answerBox = append(answerBox, tmpAns)
		}
	}
	finalRetBox, AnswerStr := GetFinalRet(answerBox)
	fmt.Println("————————————————————————————————————————————")
	fmt.Println("最后三选项得分情况为：")
	for _, v := range finalRetBox{
		fmt.Print(v.option)
		fmt.Print(" : ")
		fmt.Print(v.score)
		fmt.Printf("\n")
	}
	fmt.Println(AnswerStr)
}
