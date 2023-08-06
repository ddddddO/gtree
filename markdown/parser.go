package markdown

import (
	"errors"
	"strings"
	"sync"
)

// 一旦簡易なパーサー
type Parser struct {
	// rootが#要素かフラグ
	// #であれば次の#までのhierarchyは+=1
	isSharpRoot bool
	mu          sync.RWMutex
	spaces      int
	sep         string
}

// TODO: 要リファクタ
func NewParser() *Parser {
	return &Parser{}
}

const (
	rootHierarchyNum uint = 1
)

var (
	ErrBlankLine       = errors.New("blank line")
	ErrEmptyText       = errors.New("empty text")
	ErrIncorrectFormat = errors.New("incorrect input format")
)

// TODO: 要リファクタ
func (p *Parser) Parse(row string) (*Markdown, error) {
	// 空行か否か
	if p.isBlank(row) {
		return nil, ErrBlankLine
	}

	if strings.HasPrefix(row, sharp) {
		p.isSharpRoot = true

		_, after, found := strings.Cut(row, sharp)
		if !found {
			return nil, ErrIncorrectFormat
		}

		text := strings.Trim(strings.TrimLeft(after, sharp), space)
		if len(text) == 0 {
			return nil, ErrEmptyText
		}

		return &Markdown{
			hierarchy: rootHierarchyNum,
			text:      text,
		}, nil
	}

	p.mu.Lock()
	defer p.mu.Unlock()

	spaceCount, afterText, err := p.separateRow(row)
	if err != nil {
		return nil, err
	}

	text := strings.TrimPrefix(afterText, space)
	if len(text) == 0 {
		return nil, ErrEmptyText
	}

	return &Markdown{
		hierarchy: p.calculateHierarchy(spaceCount),
		text:      text,
	}, nil
}

func (*Parser) isBlank(row string) bool {
	r := strings.TrimSpace(row)
	return len(r) == 0
}

var listSymbols = []string{hyphen, asterisk, plus}

func (p *Parser) separateRow(row string) (int, string, error) {
	var err error
	for _, symbol := range listSymbols {
		before, after, found := strings.Cut(row, symbol)
		if !found {
			err = ErrIncorrectFormat
			continue
		}

		if len(before) != 0 {
			sep := strings.Split(before, "")[0]
			if sep == space || sep == tab {
				if p.sep == "" {
					p.sep = sep
				}
			} else {
				err = ErrIncorrectFormat
				// p.sep = ""
				continue
			}
		} else {
			p.sep = ""
		}

		spaceCount := strings.Count(before, p.sep)
		if p.sep != "" && spaceCount != len(before) {
			err = ErrIncorrectFormat
			// p.sep = ""
			continue
		}

		if p.sep == "" {
			spaceCount = 0
		}

		// Root行だとspaceCountは0にしかならない
		// Rootの次行のスペース数を一度だけ取得し、それをMarkdown全体のパースで利用する
		if spaceCount > 0 && p.spaces == 0 {
			p.spaces = spaceCount
		}

		if e := p.validateSpaces(spaceCount); e != nil {
			err = e
			// p.sep = ""
			// p.spaces = 0
			continue
		}

		return spaceCount, after, nil
	}

	return 0, "", err
}

func (p *Parser) validateSpaces(spaceCount int) error {
	if p.spaces <= 1 {
		return nil
	}
	if spaceCount%p.spaces != 0 {
		return ErrIncorrectFormat
	}
	return nil
}

func (p *Parser) calculateHierarchy(spaceCount int) uint {
	var hierarchy uint
	if p.spaces == 0 || p.sep == "" {
		hierarchy = uint(spaceCount) + rootHierarchyNum
	} else {
		hierarchy = uint(spaceCount/p.spaces) + rootHierarchyNum
	}

	if p.isSharpRoot {
		hierarchy += 1
	}
	return hierarchy
}
