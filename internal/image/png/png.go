package png

import (
	"fmt"
	"os"
)

var PNG_SINGATURE [8]byte = [8]byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A}
var IEND [4]byte = [4]byte{0x49, 0x45, 0x4E, 0x44}

// Header of PNG
type IHDR struct {
	Width           uint32
	Height          uint32
	BithDepth       byte
	ColorSpace      byte
	CompressionType byte
	FilterType      byte
	Interlacing     byte
}

// PNG signature with chunks that make it up
type PNG struct {
	Signature [8]byte
	Chunks    []Chunk
}

// Validates that all mandiatory chunks are there
func (p PNG) ValidateChunks() error {
	seenIHDR, seenIEND, IDATData := false, false, 0

	for i, chunk := range p.Chunks {
		switch string(chunk.Type[:]) {
		case "IHDR":
			if i != 0 {
				return fmt.Errorf("IHDR must be the first chunk")
			}
			seenIHDR = true
		case "IDAT":
			IDATData++
		case "IEND":
			if i != len(p.Chunks)-1 {
				return fmt.Errorf("IEND must be last chunk")
			}
			seenIEND = true
		}
	}

	if !seenIEND {
		return fmt.Errorf("missing IEND chunk")
	}

	if !seenIHDR {
		return fmt.Errorf("missing IHDR chunk")
	}

	if IDATData == 0 {
		return fmt.Errorf("no IDAT chunks found")
	}

	return nil
}

// Take a path of PNG and decodees it
func (p PNG) DecodePNG(filePath string) (*DecodedPNG, error) {

	return nil, nil

}

// Creates a png struct provided a path to a png file
func CreatePNG(filePath string) (*PNG, error) {

	file, err := os.Open(filePath)

	if os.IsNotExist(err) {
		return nil, fmt.Errorf("file: %s was not found: %w", filePath, err)
	}

	signature, err := parseSignature(file)

	if err != nil {
		return nil, err
	}

	chunks, err := parseChunks(file)

	if err != nil {
		return nil, err
	}
	png := PNG{Signature: *signature, Chunks: chunks}

	if err := png.ValidateChunks(); err != nil {
		return nil, err
	}

	return &png, nil
}

// Parses the chunks of a png file
func parseChunks(file *os.File) ([]Chunk, error) {
	var chunks []Chunk

	for {
		var chunk Chunk

		if err := chunk.Parse(file); err != nil {
			return nil, err
		}

		chunks = append(chunks, chunk)

		if chunk.Type == IEND {
			break
		}
	}

	return chunks, nil
}

// parses the signature of a png file
func parseSignature(file *os.File) (*[8]byte, error) {
	var signature [8]byte

	if _, err := file.Read(signature[:]); err != nil {
		return nil, fmt.Errorf("failed to read png signature")
	}

	if signature != PNG_SINGATURE {
		return nil, fmt.Errorf("invalid signture for png")
	}

	return &signature, nil

}
