package code

import (
	"github.com/PuerkitoBio/goquery"
)

func ParseTechnical(doc *goquery.Document) map[int]map[string]string {

	data := map[int]map[string]string{}

	doc.Find("div").Each(func(n int, selector *goquery.Selection) {
		class, _ := selector.Attr("class")
		if class == "_93444fe79c--card--ibP42 _93444fe79c--wide--gEKNN" || class == "_93444fe79c--card--ibP42"{
			tmp := map[string]string{}
			selector.Find("span").Each(func(i int, s *goquery.Selection) {
				attr, _ := s.Attr("data-mark")
				if attr == "OfferTitle" {
					tmp["title"] = s.Text()
				}
				if attr == "OfferSubtitle" {
					tmp["subtitle"] = s.Text()
				}
				if attr == "MainPrice" {
					tmp["price"] = s.Text()
				}
				if attr == "Deadline" {
					tmp["deadline"] = s.Text()
				}
				cl, _ := s.Attr("class")
				if cl == "_93444fe79c--color_gray60_100--MlpSF _93444fe79c--lineHeight_4u--fix4Q _93444fe79c--fontWeight_bold--ePDnv _93444fe79c--fontSize_10px--Ccu1C _93444fe79c--display_block--pDAEx _93444fe79c--text--g9xAG _93444fe79c--text_textTransform__uppercase--sMMpu" {
					tmp["diller"] = s.Text()
				}
				if cl == "_93444fe79c--color_current_color--gpi6p _93444fe79c--lineHeight_6u--A1GMI _93444fe79c--fontWeight_bold--ePDnv _93444fe79c--fontSize_16px--RB9YW _93444fe79c--display_block--pDAEx _93444fe79c--text--g9xAG" {
					tmp["buyer"] = s.Text()
				}
			})
			selector.Find("p").Each(func(i int, s *goquery.Selection) {
				attr, _ := s.Attr("data-mark")
				if attr == "PriceInfo" {
					tmp["quadrature"] = s.Text()
				}
			})
			selector.Find("div").Each(func(i int, s *goquery.Selection) {
				cl, _ := s.Attr("class")
				if cl == "_93444fe79c--container--w7txv" {
					tmp["metro"] = s.Text()
				}
				if cl == "_93444fe79c--labels--L8WyJ" {
					tmp["geo"] = s.Text()
				}
			})
			data[n] = tmp
		}

	})

	return data
}
