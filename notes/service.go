package notes

import (
	"sync"

	"github.com/google/uuid"
)

type NoteID string

func NewNoteID() NoteID {
	return NoteID(uuid.New().String())
}

type Note struct {
	ID     string `json:"id"`
	Author string `json:"author"`
	Text   string `json:"text"`
}

type Service struct {
	sync.Mutex
	notes map[NoteID]Note
}

func NewService() *Service {
	return &Service{
		notes: make(map[NoteID]Note),
	}
}

func (n *Service) Create(note Note) Note {
	n.Lock()
	defer n.Unlock()

	id := NewNoteID()
	note.ID = string(id)
	n.notes[id] = note
	return note
}

func (n *Service) ReadAll() []Note {
	n.Lock()
	defer n.Unlock()

	notes := make([]Note, 0, len(n.notes))
	for _, note := range n.notes {
		notes = append(notes, note)
	}
	return notes
}

func (n *Service) Read(id NoteID) (Note, bool) {
	n.Lock()
	defer n.Unlock()

	note, ok := n.notes[id]
	return note, ok
}

func (n *Service) Update(id NoteID, note Note) (Note, bool) {
	n.Lock()
	defer n.Unlock()

	_, ok := n.notes[id]
	if !ok {
		return Note{}, false
	}
	note.ID = string(id)
	n.notes[id] = note
	return note, true
}

func (n *Service) Delete(id NoteID) bool {
	n.Lock()
	defer n.Unlock()

	_, ok := n.notes[id]
	if !ok {
		return false
	}
	delete(n.notes, id)
	return true
}
