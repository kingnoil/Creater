package main

import (
	"crawler/engine"
	"crawler/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParserCityList,
	})
}


//	resp,err := http.Get( "http://www.zhenai.com/zhenghun")
//	if err != nil {
//		panic(err)
//	}
//	defer resp.Body.Close()
//	if resp.StatusCode  != http.StatusOK {
//	  fmt.Println("Errpr:status code",resp.StatusCode)
//	  return
//	}
//    e := determineEncoding(resp.Body)
//
// 	utf8Reader := transform.NewReader(resp.Body,e.NewDecoder())
//	all,err := ioutil.ReadAll(utf8Reader)
//	if err !=nil{
//		panic(err)
//	}
//	//fmt.Printf("%s\n", all)
//     printCityList(all)
//}
//func determineEncoding(r io.Reader)encoding.Encoding{
//	 bytes,err := bufio.NewReader(r).Peek(1024)
//	 if err != nil{
//	 	panic(err)
//	 }
//	e,_,_ := charset.DetermineEncoding(bytes,"")
//	return e
//}
//func printCityList(contents[]byte){
//	//<a href="http://www.zhenai.com/zhenghun/taizhou1" data-v-5e16505f>台州</a>
//	re :=regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)
//     matches :=re.FindAllSubmatch(contents,-1)
//     for _,m := range  matches{
//     		fmt.Printf("City: %s,URL: %s\n", m[2],m[1])
//	 }
//     fmt.Printf("Matches founds:%d\n",len(matches))
//}
