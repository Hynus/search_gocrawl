package main

import (
	"search_crawl/result"
)


func main() {
	searchSlice := []string{"李世民是 唐玄宗", "李世民是 唐太宗",  "李世民是 唐高宗"}
	result.GetResult(searchSlice)
}