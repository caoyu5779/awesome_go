package parser

import (
	"selfLearning/crawler/engine"
	"regexp"
)

var profileRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
var cityUrlRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)

func ParseCity(content []byte) engine.ParseResult {
	re := profileRe
	matches := re.FindAllSubmatch(content, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		name := string(m[2])
		result.Items = append(
			result.Items, "User " + name)
		result.Requests = append(
			result.Requests, engine.Request{
				Url : string(m[1]),
				ParserFunc: func(contents []byte) engine.ParseResult {
					return ParseProfile(contents, name)
				},
			})
	}

	cityMatches := cityUrlRe.FindAllSubmatch(content, -1)

	for _,m := range cityMatches {
		result.Requests = append(result.Requests, engine.Request{
			Url : string(m[1]),
			ParserFunc:ParseCity,
		})
	}
	return result
}