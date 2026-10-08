package file

import (
	"os"
	"errors"
	"io"
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
