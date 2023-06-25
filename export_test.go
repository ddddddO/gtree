package gtree

var (
	ExportErrEmptyText = errEmptyText
	// TODO: fixme
	// ExportErrIncorrectFormat = errIncorrectFormat
	ExportErrIncorrectFormat = func(row string) error {
		return &inputFormatError{row: row}
	}
)
