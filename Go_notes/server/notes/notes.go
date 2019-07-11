package notes

import "github.com/rs/xid"

//Note represents object type note
type Note struct {
	ID      string `json:"id"`
	Msg     string `json:"message"`
	Ranking int    `json:"ranking"`
}

//NoteWarehouse operations that all warehouse must implement
type NoteWarehouse interface {
	Add(note Note) bool
	Delete(noteID string) bool
	GetAll() []Note
	GetById(noteID string) Note
}

//NewNote creates a object Note with message
func NewNote(noteMsg string) Note {
	return Note{
		ID:      xid.New().String(),
		Msg:     noteMsg,
		Ranking: 0,
	}
}
