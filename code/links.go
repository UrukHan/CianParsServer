package code

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strconv"
)

func Links(url string) []string{

	links := []string {}
	txt := ""
	count := ""

	doc := MakeRequest(url)
	doc.Find("div").Find("h5").Each(func(n int, selector *goquery.Selection) {
		class, _ := selector.Attr("class")
		if class == "_93444fe79c--color_black_100--kPHhJ _93444fe79c--lineHeight_20px--tUURJ _93444fe79c--fontWeight_bold--ePDnv _93444fe79c--fontSize_14px--TCfeJ _93444fe79c--display_block--pDAEx _93444fe79c--text--g9xAG _93444fe79c--text_letterSpacing__normal--xbqP6" {
			txt = selector.Text()
		}
	})

	for _, i := range txt {
		for _, j := range []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"} {
			if string(i) == j {
				count = count + j
			}
		}
	}

	intCount, err := strconv.Atoi(count)
	if err != nil {
		fmt.Println(err)
	}

	for i := 1; i <= intCount/28; i++  {
		strI := strconv.Itoa(i)
		links = append(links, url + "p=" + strI)
	}
	//fmt.Printf("%v", links)
	return links
}

