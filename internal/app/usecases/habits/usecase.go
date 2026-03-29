package habitsUsecase

import (
	"context"

	"github.com/zura-t/observer.dev/internal/app/models"
	"github.com/zura-t/observer.dev/internal/config"
)

type HabitsUsecase interface {
	GetHabits(context context.Context, userID uint64) (*[]models.Habit, error)
	GetHabit(context context.Context, id uint64, userID uint64) (*models.Habit, error)
	CreateHabit(context context.Context, habit *CreateHabit) (*models.Habit, error)
	UpdateHabit(context context.Context, habit *UpdateHabit) (*models.Habit, error)
	DeleteHabit(context context.Context, id uint64, userID uint64) error
	GetHabitLogs(context context.Context, filter *GetHabitLogsFilter) (*[]models.HabitLog, error)
	GetHabitLog(context context.Context, id uint64, userID uint64) (*models.HabitLog, error)
	CreateHabitLog(context context.Context, habitLog *CreateHabitLog) (*models.HabitLog, error)
	UpdateHabitLog(context context.Context, habitLog *UpdateHabitLog) (*models.HabitLog, error)
	DeleteHabitLog(context context.Context, id uint64, userID uint64) error
	DeleteHabitLogs(context context.Context, habitID uint64, userID uint64) error
	DeleteAllHabitLogs(context context.Context, userID uint64) error
}

var (
	_ HabitsUsecase = (*usecase)(nil)
)

type usecase struct {
	habitsRepo HabitsRepositoryInterface
	config     *config.Config
}

func New(habitsRepo HabitsRepositoryInterface, config *config.Config) *usecase {
	return &usecase{
		habitsRepo: habitsRepo,
		config:     config,
	}
}
