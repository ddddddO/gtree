package dot

import (
	"errors"
	"log"
	"net/http"

	gq "github.com/PuerkitoBio/goquery"
)

const livedoor = "https://blog.livedoor.com/ranking/blog/"

func Scrape(cnt int) ([]Node, [][]string, error) {
	resp, err := http.Get(livedoor)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, nil, errors.New("status code is not 200")
	}

	doc, err := gq.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, nil, err
	}

	liSel := doc.Find("#lb-container .ranking-inner > div > div > ul").First().Find("li a")

	var (
		nodes  []Node
		edges  [][]string
		parent string
	)

	// 一番親ノード
	nodes = append(nodes, Node{
		Title:   quote("livedoor"),
		ToolTip: quoteNonShrink("livedoor"),
		Url:     quoteNonShrink(livedoor),
		Shape:   "doubleoctagon",
	})

	// ul内のli要素(100) * a要素(3)だけ繰り返す
	liSel.EachWithBreak(func(i int, sel *gq.Selection) bool {
		if (i % 4) == 0 {
			parent = ""
		}

		title := sel.Find(".text").Text()
		href, _ := sel.Attr("href")
		// 個々の大元のまとめ
		if len(title) == 0 {
			parent = sel.Text()
			nodes = append(nodes, Node{
				Title:   quote(parent),
				ToolTip: quoteNonShrink(parent),
				Url:     quoteNonShrink(href),
				Shape:   "box",
			})
			edges = append(edges, []string{quote("livedoor"), quote(parent)})
			return true
		}

		Thread(href)

		nodes = append(nodes, Node{
			Title:   quote(title),
			ToolTip: quoteNonShrink(title),
			Url:     quoteNonShrink(href),
			Shape:   "ellipse",
		})
		edges = append(edges, []string{quote(parent), quote(title)})

		// ランキング上位のcnt個数まで、でEachを抜ける
		if i >= cnt {
			return false
		}

		return true
	})

	return nodes, edges, nil
}

func quoteNonShrink(s string) string {
	return `"` + s + `"`
}

func quote(s string) string {
	return `"` + shrink(s) + `"`
}

func shrink(s string) string {
	chars := []rune(s)
	if len(chars) > 10 {
		return string(chars[:10]) + "..."
	}
	return s
}

func Thread(url string) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		panic("status code is not 200")
		//return nil, nil, errors.New("status code is not 200")
	}

	doc, err := gq.NewDocumentFromReader(resp.Body)
	if err != nil {
		panic(err)
	}

	log.Println(url)
	log.Println(finder(doc))
}

var selectors = []string{
	".article-body-more",
	".article-body",
	"#articlebody",
	".entrybody",
	".more_body",
}

func finder(doc *gq.Document) int {
	for _, selector := range selectors {
		n := doc.Find(selector).Length()
		if n != 0 {
			log.Println(selector)
			return n
		}
	}
	log.Println("not match selector")
	return 0
}
