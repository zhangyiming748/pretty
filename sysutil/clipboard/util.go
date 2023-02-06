package clipboard

// clipboard writer, reader program names
const (
	WriterOnMac = "pbcopy"
	WriterOnWin = "clip" // clip only support write contents to clipboard.
	WriterOnLin = "xsel"

	ReaderOnMac = "pbpaste"
	ReaderOnWin = "powershell get-clipboard" // read should use: powershell get-clipboard
	ReaderOnLin = "xclip"
)

// std instance
var std = New()

// Reset clipboard data
func Reset() error {
	return std.Reset()
}

// Available clipboard available check
func Available() bool {
	return std.Available()
}

// ReadString contents from clipboard
func ReadString() (string, error) {
	return std.ReadString()
}

// WriteString contents to clipboard and flush
func WriteString(s string) error {
	if _, err := std.WriteString(s); err != nil {
		return err
	}
	return std.Flush()
}
