package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type IO struct {
	reader* bufio.Reader
}

func NewIO() IO {
	return IO{
		reader: bufio.NewReader(os.Stdin),
	}
}

func (io *IO) ReadString(question string) (string, error) {
	fmt.Print(question)
	text, err := io.reader.ReadString('\n')
	return strings.TrimRight(text, "\n"), err
}

func (io *IO) ReadInt(question string) (int, error) {
	text, err := io.ReadString(question)
	if err != nil {
		return 0, err
	}

	if len(text) == 0 {
		return 0, nil
	}

	num, err := strconv.Atoi(text)
	if err != nil {
		return 0, err
	}

	return num, nil
}
