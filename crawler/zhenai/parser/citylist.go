package parser

import (
	"crawler/engine"
	"regexp"
)
const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`
func ParserCityList( contents []byte)engine.ParserResult{

		//<a href="http://www.zhenai.com/zhenghun/taizhou1" data-v-5e16505f>台州</a>
		re :=regexp.MustCompile(cityListRe)
		matches :=re.FindAllSubmatch(contents,-1)
		result := engine.ParserResult{}
		limit := 10
		for _,m := range  matches{
			result.Items= append(result.Items,"City "+ string(m[2]))
			result.Requests = append(
				result.Requests,
				engine.Request{Url:string(m[1]),
					ParserFunc:ParseCity,
				})
			limit--
            if limit == 0{
				break
			}
		}
		 return result
}
