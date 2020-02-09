package arprint

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type format struct {
	name, magic string
	magicOffset int
	reader      NewReader
}

type NewReader func(*os.File) (io.Reader, error)

var formats []format

func RegisterFormat(name, magic string, magicOffset int, f NewReader) {
	formats = append(formats, format{name, magic, magicOffset, f})
}

func Open(file *os.File) (io.Reader, error) {
	var found *format
	r := bufio.NewReader(file)
	for _, f := range formats {
		p, err := r.Peek(f.magicOffset + len(f.magic))
		if err != nil {
			continue
		}
		if string(p[f.magicOffset:]) == f.magic {
			found = &f
			break
		}
	}
	if found == nil {
		return nil, fmt.Errorf("open archive: can't determine format")
	}
	_, err := file.Seek(0, os.SEEK_SET)
	if err != nil {
		return nil, fmt.Errorf("open archive: %s", err)
	}
	return found.reader(file)
}
