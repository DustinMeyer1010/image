package png

import (
	"encoding/binary"
	"fmt"
	"hash/crc32"
	"os"
)

// Represents a chunk of a png
type Chunk struct {
	Length uint32
	Type   [4]byte
	Data   []byte
	CRC    uint32
}

// Computes crc32 Checksum on chunk to make sure it is valid
func (c *Chunk) ValidCRC() error {

	table := crc32.MakeTable(crc32.IEEE)
	computedCRCR := crc32.Checksum(append(c.Type[:], c.Data...), table)

	if c.CRC != computedCRCR {
		return fmt.Errorf("invalid CRC for chunk: %s", string(c.Type[:]))
	}

	return nil
}

// Parse a chunk from file
// Order of steps does matter for parsing chunk
func (c *Chunk) Parse(file *os.File) error {

	steps := []func(*os.File) error{
		c.parseLength,
		c.parseType,
		c.parseData,
		c.parseCRC,
	}

	for _, step := range steps {
		if err := step(file); err != nil {
			return nil
		}
	}

	return c.ValidCRC()
}

// Pulls the length from file chunk
func (c *Chunk) parseLength(file *os.File) error {
	var length [4]byte

	n, err := file.Read(length[:])

	if err != nil {
		return fmt.Errorf("failed to read chunk length: %w", err)
	}

	if n != 4 {
		return fmt.Errorf("chunk length should be 4 bytes")
	}

	c.Length = binary.BigEndian.Uint32(length[:])

	return nil

}

// Pull the type from the file chunk
func (c *Chunk) parseType(file *os.File) error {

	n, err := file.Read(c.Type[:])

	if err != nil {
		return fmt.Errorf("failed to read chunk type: %w", err)
	}

	if n != 4 {
		return fmt.Errorf("chunk type should be 4 bytes")
	}

	return nil
}

// Pull the data from the file chunk
func (c *Chunk) parseData(file *os.File) error {

	c.Data = make([]byte, c.Length)
	n, err := file.Read(c.Data)

	if err != nil {
		return fmt.Errorf("faild to read chunk data: %w", err)
	}

	if n != int(c.Length) {
		return fmt.Errorf("chunk data length did not match chunk length")
	}

	return nil
}

// Pull the CRC from the file chunk
func (c *Chunk) parseCRC(file *os.File) error {
	CRC := make([]byte, 4)
	n, err := file.Read(CRC)

	if err != nil {
		return fmt.Errorf("failed to read chunk CRC: %w", err)
	}

	if n != 4 {
		return fmt.Errorf("chunk CRC should be 4 bytes")
	}

	c.CRC = binary.BigEndian.Uint32(CRC)

	return nil
}
