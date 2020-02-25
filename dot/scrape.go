package dot

import (
	"errors"
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
			nodes = append(nodes,
				Node{
					Title: quote(parent),
					Url:   quoteUrl(href),
				},
			)
			edges = append(edges, []string{quote("livedoor"), quote(parent)})
			return true
		}

		nodes = append(nodes,
			Node{
				Title: quote(title),
				Url:   quoteUrl(href),
			},
		)
		edges = append(edges, []string{quote(parent), quote(title)})

		// ランキング上位のcnt個数まで、でEachを抜ける
		if i >= cnt {
			return false
		}

		return true
	})

	return nodes, edges, nil
}

func quoteUrl(u string) string {
	return `"` + u + `"`
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
