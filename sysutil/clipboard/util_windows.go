//go:build windows
// +build windows

package clipboard

// GetWriterBin program name
func GetWriterBin() string {
	return WriterOnWin
}

// GetReaderBin program name
func GetReaderBin() string {
	return ReaderOnWin
}
