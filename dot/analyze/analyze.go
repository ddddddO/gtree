package analyze

// https://qiita.com/toshiyuki_tsutsui/items/604f92dbe6e20a18a17e
type Analyzer interface {
	// url/ressをもとに感情分析
	analyze() (interface{}, error)
}

func Run(an Analyzer) error {
	_, err := an.analyze()
	if err != nil {
		return err
	}
	return nil
}
