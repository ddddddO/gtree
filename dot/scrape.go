package dot

import (
	"bufio"
	"errors"
	"io"
	"log"
	"net/http"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding/htmlindex"

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
	//url = "http://blog.livedoor.jp/dqnplus/archives/1998364.html" // euc-jp
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		panic("status code is not 200")
		//return nil, nil, errors.New("status code is not 200")
	}

	// NOTE: 文字化け対策
	// https://github.com/PuerkitoBio/goquery/wiki/Tips-and-tricks
	utfBody, err := DecodeHTMLBody(resp)
	if err != nil {
		panic(err)
	}

	doc, err := gq.NewDocumentFromReader(utfBody)
	if err != nil {
		panic(err)
	}

	log.Println(url)
	extractRes(finder(doc))
	log.Println("----")
}

var selectors = map[string][]string{
	".article-body-more": []string{".t_b", "t_b b", "strong span", "b"},
	".article-body":      []string{".t_b"},
	"#articlebody":       []string{".t_b", "b"},
	".more_body":         []string{".t_b"},
	//".entrybody": []string{},
}

func finder(doc *gq.Document) *gq.Selection {
	for threadSelector, resSelectors := range selectors {
		threadDoc := doc.Find(threadSelector)
		if threadDoc.Length() != 0 {
			for _, resSelector := range resSelectors {
				resDoc := threadDoc.Find(resSelector)
				if resDoc.Length() != 0 {
					return resDoc
				}
			}
			log.Println("no match res selector")
			return nil
		}
	}
	log.Println("no match thread selector")
	return nil
}

func extractRes(resDoc *gq.Selection) {
	if resDoc == nil {
		return
	}

	resDoc.Each(func(i int, sel *gq.Selection) {
		log.Println(sel.Text())
	})
}

func detectContentCharset(resp *http.Response) string {
	r := bufio.NewReader(resp.Body)
	if data, err := r.Peek(1024); err == nil {
		if _, name, ok := charset.DetermineEncoding(data, resp.Header.Get("content-type")); ok {
			return name
		}
	}
	return "utf-8"
}

// DecodeHTMLBody returns an decoding reader of the html Body for the specified `charset`
// If `charset` is empty, DecodeHTMLBody tries to guess the encoding from the content
func DecodeHTMLBody(resp *http.Response) (io.Reader, error) {
	charset := detectContentCharset(resp)
	log.Println("charset:", charset)

	e, err := htmlindex.Get(charset)
	if err != nil {
		return nil, err
	}

	if name, _ := htmlindex.Name(e); name != "utf-8" {
		log.Println("name:", name)
		return e.NewDecoder().Reader(resp.Body), nil
	}
	return resp.Body, nil
}
