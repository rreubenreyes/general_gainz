package history

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/rreubenreyes/general_gainz/internal/training"
)

type Kind string

const (
	Exercise Kind = "exercise"
)

type Event struct {
	Id        uuid.UUID `json:"id"`
	HashKey   string    `json:"hash_key"`
	RangeKey  string    `json:"range_key"`
	CreatedAt string    `json:"created_at"`
	Kind      Kind      `json:"kind"`
	Data      []byte    `json:"data"`
}

func fromExercise(e training.Exercise) (*Event, error) {
	now := time.Now().Format(time.RFC3339)
	data, err := json.Marshal(e)
	if err != nil {
		return nil, err
	}

	evt := &Event{
		Id:        uuid.New(),
		HashKey:   e.Name + string(e.Tier),
		RangeKey:  now,
		CreatedAt: now,
		Kind:      Exercise,
		Data:      data,
	}

	return evt, nil
}

func From(i any) (*Event, error) {
	switch t := i.(type) {
	case training.Exercise:
		return fromExercise(t)
	default:
		err := fmt.Errorf("cannot call From with unsupported type %T", t)
		return nil, err
	}
}
