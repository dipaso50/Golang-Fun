package memorywarehouse

import (
	"Go_notes/server/notes"
)

type MemoryWarehouse struct {
	allNotes map[string]notes.Note
}

func NewMemoryWarehouse() MemoryWarehouse {
	return MemoryWarehouse{make(map[string]notes.Note)}
}

func (m MemoryWarehouse) Add(note notes.Note) bool {
	m.allNotes[note.ID] = note
	return true
}

func (m MemoryWarehouse) GetAll() []notes.Note {
	v := make([]notes.Note, 0, len(m.allNotes))

	for _, value := range m.allNotes {
		v = append(v, value)
	}

	return v
}

func (m MemoryWarehouse) Delete(noteID string) bool {
	delete(m.allNotes, noteID)
	return true
}

func (m MemoryWarehouse) GetById(noteID string) notes.Note {
	return m.allNotes[noteID]
}
