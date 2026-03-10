package userUsecase

import "context"

func (u *usecase) DeleteUser(ctx context.Context, id uint64) error {
	return u.userRepo.DeleteUser(ctx, id)
}
