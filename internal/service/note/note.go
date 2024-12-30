package note

type NoteService struct {
	NoteRepository
}

func New(repo NoteRepository) (*NoteService, error) {
	return &NoteService{NoteRepository: repo}, nil
}
