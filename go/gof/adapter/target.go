package adapter

// Target役(必要となっているメソッドを定めている役)
type Target interface {
	OutputSharpFrame()
	OutputHyphenFrame()
}
