package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type ItemNotFoundError struct {
	itemName string
}

func (e ItemNotFoundError) Error() string {
	return fmt.Sprintf("List \"%s\" not found", e.itemName)
}

type ListNotFoundError struct {
	listName string
}

func (e ListNotFoundError) Error() string {
	return fmt.Sprintf("List \"%s\" not found", e.listName)
}

type State struct {
	lists        ListMap
	selectedList string
}

func (state State) GetSelected() string {
	return state.selectedList
}

func (s *State) SetSelected(listName string) error {
	if _, ok := s.lists[listName]; ok {
		s.selectedList = listName
		return nil
	}

	return ListNotFoundError{listName}
}

func (s State) GetListKeys() []string {
	var keys []string

	for k := range s.lists {
		keys = append(keys, k)
	}

	return keys
}

func (s State) GetList(listName string) ([]Item, error) {
	if s.listExists(listName) {
		return s.lists[listName], nil
	}

	return nil, ListNotFoundError{listName}
}

func (s *State) AddList(listName string) error {
	if _, ok := s.lists[listName]; ok {
		return fmt.Errorf("List \"%s\" already exists", listName)
	}

	s.lists[listName] = []Item{}

	return nil
}

func (s *State) listExists(listName string) bool {
	_, ok := s.lists[listName]

	return ok
}

func (s *State) RemoveList(listName string) error {
	if s.listExists(listName) {
		delete(s.lists, listName)
		return nil
	}

	return ListNotFoundError{listName}
}

func (s *State) AddItemToList(listName, itemName string) error {
	if s.listExists(listName) {
		item := NewItem(itemName)
		s.lists[listName] = append(s.lists[listName], item)
		return nil
	}

	return ListNotFoundError{listName}
}

func (s *State) MarkItemComplete(listName, itemName string) error {
	if !s.listExists(listName) {
		return ListNotFoundError{listName}
	}

	for i := range s.lists[listName] {
		item := &s.lists[listName][i]
		if item.Name == itemName {
			item.Completed = time.Now()
			return nil
		}
	}

	return ItemNotFoundError{itemName}
}

func (s *State) RemoveItemFromList(listName, itemName string) error {
	if !s.listExists(listName) {
		return ListNotFoundError{listName}
	}

	for i, v := range s.lists[listName] {
		if v.Name == itemName {
			s.lists[listName] = append(s.lists[listName][:i], s.lists[listName][i+1:]...)
			return nil
		}
	}

	return ItemNotFoundError{itemName}
}

func (state *State) SaveToFile(path string) error {
	bytes, err := json.Marshal(state.lists)
	if err != nil {
		return err
	}

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(bytes)
	if err != nil {
		return err
	}

	return nil
}

func NewStateFromFile(path string) (*State, error) {
	var state = &State{}

	loaded, err := NewListMapFromFile(path)
	if err != nil {
		return state, err
	}

	state.lists = loaded

	err = state.SetSelected(state.lists.First())
	if err != nil {
		return state, err
	}
	fmt.Printf("Setting selected to %s\n", state.selectedList)

	return state, nil
}
