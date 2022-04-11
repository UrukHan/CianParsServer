package code

import (
	"math/rand"
	"time"
)

func Response(link string, max int) map[string][]string {

	resp := map[string][]string {"title": []string{}, "subtitle": []string{}, "price": []string{}, "deadline": []string{}, "diller": []string{}, "buyer": []string{}, "quadrature": []string{}, "metro": []string{}, "geo": []string{}}
	links := Links(link)
	if len(links) < max {
		max = len(links)
	}
	for _, link := range links[:max] {
		doc := MakeRequest(link)
		dict := ParseTechnical(doc)

		duration := time.Duration(rand.Intn(5)+1)*time.Second
		time.Sleep(duration)

		for _, v := range dict {
			if _,ok := v["title"]; !ok {
				resp["title"] = append(resp["title"], "unknoun")
			} else {
				resp["title"] = append(resp["title"], v["title"])
			}
			if _,ok := v["subtitle"]; !ok {
				resp["subtitle"] = append(resp["subtitle"], "unknoun")
			} else {
				resp["subtitle"] = append(resp["subtitle"], v["subtitle"])
			}
			if _,ok := v["price"]; !ok {
				resp["price"] = append(resp["price"], "unknoun")
			} else {
				resp["price"] = append(resp["price"], v["price"])
			}
			if _,ok := v["deadline"]; !ok {
				resp["deadline"] = append(resp["deadline"], "unknoun")
			} else {
				resp["deadline"] = append(resp["deadline"], v["deadline"])
			}
			if _,ok := v["title"]; !ok {
				resp["diller"] = append(resp["diller"], "unknoun")
			} else {
				resp["diller"] = append(resp["diller"], v["diller"])
			}
			if _,ok := v["buyer"]; !ok {
				resp["buyer"] = append(resp["buyer"], "unknoun")
			} else {
				resp["buyer"] = append(resp["buyer"], v["buyer"])
			}
			if _,ok := v["quadrature"]; !ok {
				resp["quadrature"] = append(resp["quadrature"], "unknoun")
			} else {
				resp["quadrature"] = append(resp["quadrature"], v["quadrature"])
			}
			if _,ok := v["metro"]; !ok {
				resp["metro"] = append(resp["metro"], "unknoun")
			} else {
				resp["metro"] = append(resp["metro"], v["metro"])
			}
			if _,ok := v["geo"]; !ok {
				resp["geo"] = append(resp["geo"], "unknoun")
			} else {
				resp["geo"] = append(resp["geo"], v["geo"])
			}

		}

	}
	return resp
}

