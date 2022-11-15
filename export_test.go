package gtree

var (
	ExportErrEmptyText = errEmptyText
	// TODO: fixme
	// ExportErrIncorrectFormat = errIncorrectFormat
	ExportErrIncorrectFormat = &inputFormatError{}
)
