package parser

import (
	"crawler/engine"
	"crawler/model"
	"regexp"
	"strconv"
)
var ageRe =regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>([\d]+)岁</div>`)
var marriageRe =regexp.MustCompile(`"marriageString":"([^" ]+)"`)
var Gender = regexp.MustCompile(`"genderString":"([^"]+)"`)
var Height = regexp.MustCompile(`"heightString":"([\d]+)cm"`)
var Education = regexp.MustCompile(`"educationString":"([^"]+)"`)
var BaseInfo = regexp.MustCompile(`"basicInfo":\[([^\[]+)\]`)
var Money =  regexp.MustCompile(`月收入:([^<]+)`)
func ParseProfile(contents []byte, name string) engine.ParserResult{
	profile := model.Profile{}
	profile.Name = name
	 age, err := strconv.Atoi(extractString(contents,ageRe))
	 if err == nil{
	     profile.Age= age
	 }
	 profile.Marriage = extractString(contents,marriageRe)
	 profile.Gender = extractString(contents,Gender)
	 height, _ := strconv.Atoi(extractString(contents,Height))
	 //log.Print(height)
	 profile.Height =height
	 profile.Education = extractString(contents,Education)
	 profile.BasicInfo = extractString(contents,BaseInfo)
	 profile.Money = extractString(contents,Money)
	// log.Print(profile.BasicInfo)
	 result := engine.ParserResult{ Items :[]interface{}{profile},}

     return result
}
func extractString(contents []byte, re *regexp.Regexp)string{
	match := re.FindSubmatch(contents)
    if len(match) >=2{
        return string(match[1])
	}else{
		return ""
	}
}