package habitsRepo

import (
	"errors"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/zura-t/observer.dev/internal/app/models"
	habitsUsecase "github.com/zura-t/observer.dev/internal/app/usecases/habits"
)

func newHabit(o *habitsUsecase.CreateHabit) *habit {
	return &habit{
		Title:       o.Title,
		Frequency:   o.Frequency,
		TargetCount: o.TargetCount,
		UserID:      o.UserID,
	}
}

func newHabitLog(o *habitsUsecase.CreateHabitLog) *habitLog {
	return &habitLog{
		HabitID:     o.HabitID,
		LogDate:     o.LogDate,
		ActualCount: o.ActualCount,
		Note:        o.Note,
	}
}

type PGErrorMap map[error]error

var habitErrMap = PGErrorMap{
    pgx.ErrNoRows: models.ErrHabitNotFound,
}

var habitLogErrMap = PGErrorMap{
    pgx.ErrNoRows: models.ErrHabitLogNotFound,
}

func convertPGError(err error, errMap PGErrorMap) error {
	if err == nil {
		return nil
	}

	log.Print(err)

	if errors.Is(err, pgx.ErrNoRows) {
		if mapped, ok := errMap[pgx.ErrNoRows]; ok {
			return mapped
		}
	}
	return err
}
