package markdown

import (
	"errors"
	"strings"
)

// 一旦簡易なパーサー
type Parser struct {
	// rootが#要素かフラグ
	// #であれば次の#までのhierarchyは+=1
	isSharpRoot bool
	spaces      int
	sep         string
}

// TODO: 要リファクタ
func NewParser(spaces int) *Parser {
	sep := space
	if spaces == 1 {
		sep = tab
	}

	return &Parser{
		spaces: spaces,
		sep:    sep,
	}
}

const (
	sharp = "#"

	hyphen   = "-"
	asterisk = "*"
	plus     = "+"

	space = " "
	tab   = "\t"
)

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
	if p.isEmpty(row) {
		return nil, ErrBlankLine
	}

	if strings.HasPrefix(row, sharp) {
		p.isSharpRoot = true

		_, after, found := strings.Cut(row, sharp)
		if !found {
			return nil, ErrIncorrectFormat
		}
		text := strings.TrimLeft(after, sharp)
		text = strings.Trim(text, space)
		if len(text) == 0 {
			return nil, ErrEmptyText
		}

		return &Markdown{
			hierarchy: rootHierarchyNum,
			text:      text,
		}, nil
	}

	// #のセクション内
	if p.isSharpRoot {
		spaceCount, afterText, err := p.separateRow(row)
		if err != nil {
			return nil, err
		}

		text := strings.TrimPrefix(afterText, space)
		if len(text) == 0 {
			return nil, ErrEmptyText
		}

		hierarchy := p.calculateHierarchy(spaceCount)
		return &Markdown{
			hierarchy: hierarchy,
			text:      text,
		}, nil
	}

	spaceCount, afterText, err := p.separateRow(row)
	if err != nil {
		return nil, err
	}

	text := strings.TrimPrefix(afterText, space)
	if len(text) == 0 {
		return nil, ErrEmptyText
	}

	hierarchy := p.calculateHierarchy(spaceCount)
	return &Markdown{
		hierarchy: hierarchy,
		text:      text,
	}, nil
}

func (p *Parser) isEmpty(row string) bool {
	r := strings.TrimSpace(row)
	return len(r) == 0
}

func (p *Parser) validSpaces(spaceCount int) error {
	if p.isTab() {
		return nil
	}
	if spaceCount%p.spaces != 0 {
		return ErrIncorrectFormat
	}
	return nil
}

func (p *Parser) calculateHierarchy(spaceCount int) uint {
	var hierarchy uint
	if p.isTab() {
		hierarchy = uint(spaceCount) + rootHierarchyNum
	} else {
		hierarchy = uint(spaceCount/p.spaces) + rootHierarchyNum
	}

	if p.isSharpRoot {
		hierarchy += 1
	}
	return hierarchy
}

func (p *Parser) isTab() bool {
	return p.spaces == 1
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
		spaceCount := strings.Count(before, p.sep)
		if spaceCount != len(before) {
			err = ErrIncorrectFormat
			continue
		}
		if e := p.validSpaces(spaceCount); e != nil {
			err = e
			continue
		}

		return spaceCount, after, nil
	}

	return 0, "", err
}
