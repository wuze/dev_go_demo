package main

import (
	"fmt"
	"regexp"
)

type myRegexp struct {
	*regexp.Regexp
}

var plist = myRegexp{regexp.MustCompile(`<div.*?>(?P<first>.*?)</div>`)}

//var plist = myRegexp{regexp.MustCompile(`\w+`)}

var str = "<div class=\"view  view-noCom\" id=\"J_ItemList\" data-spm=\"a220m.1000858.1000725\" data-area=\"\" data-atp-a=\"{p},{id},,,spu,1,spu,{user_id}\" data-atp-b=\"{p},{id},,,spu,2,spu,{user_id}\">FUCKU</div>"

//var str = "HelloWorld"

func main() {

	var ret = plist.FindString(str)
	for pos, i := range ret {
		fmt.Printf("Key:%v  Value: %v \n", pos, i)
	}

}

func (r *myRegexp) FindString(body string) map[string]string {

	captures := make(map[string]string, 0)
	matches := r.FindAllStringSubmatch(body, -1)

	if matches == nil {
		return captures
	}

	for _, match := range matches {

		//		fmt.Printf(match)

		for pos, val := range match {
			key := fmt.Sprintf("%d", pos)
			captures[key] = val
		}
	}

	return captures
}
