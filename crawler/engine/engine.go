package engine

import (
	"crawler/fetcher"
	"log"
)

func Run(seeds ...Request){
	var requests[]Request
	for _,r := range seeds{
		requests = append(requests,r)

	}


	for len(requests)> 0{

		r := requests[0]
		requests = requests[1:]
		log.Printf("Fetching %s",r.Url)
		body,err := fetcher.Fetch(r.Url)
		if err != nil{
			log.Printf("Fetcher: error"+"fetching url %s:%v",r.Url,err)
  			continue
		}
	 ParserResult := r.ParserFunc(body)
	 requests = append(requests,ParserResult.Requests...)
     for _, items := range ParserResult.Items{
		 log.Printf("Got item %v", items )
	 }

	}

}
