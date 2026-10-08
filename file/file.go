package file

import (
	"errors"
	"io"
	"os"
	"strings"
	// "fmt"
)

func ReadFile(file *os.File) ([]byte, error) {
  buffer := make([]byte, 0, 512)
  for {
    if len(buffer) == cap(buffer) {
      buffer = append(buffer, 0)[:len(buffer)]
    }
    offset, err := file.Read(buffer[len(buffer):cap(buffer)])
    buffer = buffer[:len(buffer)+offset]
    if err != nil {
      if errors.Is(err, io.EOF) {
        err = nil
      }
      return buffer, err
    }
  }
}

func CheckContent(data string) error {
	keyError := errors.New("key isn't 'user', 'mode' or 'node'")
	lines := strings.Split(data, "\n")
	for _, l := range lines {
		l = strings.TrimSpace(l)
		if l == ""{
			continue
		}
		parts := strings.Split(l, ":")
		if len(parts) < 2 {
			return keyError
		}

		key := strings.TrimSpace(parts[0])

		if key != "user" && key != "mode" && key != "node"{
			return keyError
		}
	}
	return nil
}
