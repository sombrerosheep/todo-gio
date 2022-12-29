package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type ListMap map[string][]Item

func NewListMapFromFile(path string) (ListMap, error) {
	f, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var list ListMap
	err = json.Unmarshal(f, &list)
	if err != nil {
		return nil, err
	}

	return list, err
}

func (lm ListMap) First() string {
	for k := range lm {
		return k
	}

	return ""
}

type Item struct {
	Name      string    `json:"name"`
	Created   time.Time `json:"created"`
	Completed time.Time `json:"completed"`
}

func NewItem(name string) Item {
	item := Item{
		Name:      name,
		Created:   time.Now(),
		Completed: time.Time{},
	}

	return item
}

func (i Item) String() string {
	status := func(i Item) string {
		if i.Completed.IsZero() {
			return "Incomplete"
		}

		return "Complete"
	}(i)

	return fmt.Sprintf("%s: %s (%s)",
		i.Name,
		i.Created.Format(time.RFC3339Nano),
		status,
	)
}

func (i *Item) Complete() {
	i.Completed = time.Now()
}
