package fileio

import (
	"fmt"
	"os"
	"strings"
)

// TODO - implement proper error handling

func WriteFile(path string, buf []string) ([]string, error) {
	f, err := os.Create(path)
	if err != nil {
		return buf, err
	}
	defer f.Close()

	// Remove trailing empty lines from buf
	realBufLen := len(buf)
	for i := len(buf) - 1; i >= 0; i-- {
		if strings.Trim(buf[i], " \t") != "" {
			break
		}
		realBufLen--
	}
	buf = buf[:realBufLen]

	for _, str := range buf {
		if _, err := fmt.Fprintf(f, "%s\n", str); err != nil {
			return buf, err
		}
	}
	return append(buf, ""), nil
}

func LoadFileBuf(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	stat, err := f.Stat()
	if err != nil {
		return nil, err
	}

	size := stat.Size()
	buf := make([]byte, size)
	if _, err := f.Read(buf); err != nil {
		return nil, err
	}

	res := strings.Split(string(buf), "\n")
	return res, nil
}
