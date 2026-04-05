package notesUsecase

import "github.com/zura-t/observer.dev/internal/config"

type NotesUsecase interface{}

var (
	_ NotesUsecase = (*usecase)(nil)
)

type usecase struct {
	notesRepo NotesRepositoryInterface
	config    *config.Config
}

func New(notesRepo NotesRepositoryInterface, config *config.Config) *usecase {
	return &usecase{
		notesRepo: notesRepo,
		config:    config,
	}
}
