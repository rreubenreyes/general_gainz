package log

import (
	"bufio"
	"encoding/json"
	"os"

	"github.com/rreubenreyes/general_gainz/internal/history"
)

type Appender interface {
	Append(*history.Event) ([]byte, error)
	Head() (*history.Event, error)
}

type FS struct {
	Path string
}

func (f *FS) Append(e *history.Event) ([]byte, error) {
	file, err := os.OpenFile(f.Path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(e)
	if err != nil {
		return nil, err
	}
	file.Write(data)

	return data, nil
}

func (f *FS) Head() (*history.Event, error) {
	file, err := os.OpenFile(f.Path, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}

	r := bufio.NewReader(file)
	raw, _, err := r.ReadLine()
	if err != nil {
		return nil, err
	}

	var data *history.Event
	err = json.Unmarshal(raw, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
