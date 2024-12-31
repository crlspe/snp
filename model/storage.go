package model

type Storager interface {
	Save()
}

type Storage struct {
	Snippet Snippet
}

func (s *Storage) New(snippet Snippet) {
	s.Snippet = snippet
}
