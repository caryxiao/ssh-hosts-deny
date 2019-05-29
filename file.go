package sshHostsDeny

import (
	"bufio"
	"io"
	"os"
)

func ReadFile(filePath string, handle func([]byte)) (err error) {
	f, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer f.Close()

	buf := bufio.NewReader(f)

	for {
		line, err := buf.ReadBytes('\n')
		handle(line)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
}

func WriteFile(filePath string, content string) (err error) {
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return
	}
	defer f.Close()

	writer := bufio.NewWriter(f)
	str := []byte(content + "\n")
	_, err = writer.Write(str)

	if err != nil {
		return
	}

	err = writer.Flush()
	return
}
