package png

// Decoded version of PNG
type DecodedPNG struct {
	Header   IHDR
	MetaData map[string]string
	IDATData []byte
}
