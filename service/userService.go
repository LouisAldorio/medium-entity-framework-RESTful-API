package service

import (
	"context"
	"myapp/config"
	"myapp/ent"
)

type UserOps struct {
	ctx    context.Context
	client *ent.Client
}

func NewUserOps(ctx context.Context) *UserOps {
	return &UserOps{
		ctx:    ctx,
		client: config.GetClient(),
	}
}

func (r *UserOps) UsersGetAll() ([]*ent.User, error) {

	users, err := r.client.User.Query().WithTodos().All(r.ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserOps) UserGetByID(id int) (*ent.User, error) {

	user, err := r.client.User.Get(r.ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserOps) UserCreate(newUser ent.User) (*ent.User, error) {

	newCreatedUser, err := r.client.User.Create().
		SetEmail(newUser.Email).
		SetName(newUser.Name).
		Save(r.ctx)

	if err != nil {
		return nil, err
	}

	return newCreatedUser, nil
}

func (r *UserOps) UserUpdate(user ent.User) (*ent.User, error) {

	updatedUser, err := r.client.User.UpdateOneID(user.ID).
		SetEmail(user.Email).
		SetName(user.Name).Save(r.ctx)

	if err != nil  {
		return nil, err
	}

	return updatedUser, nil
}

func (r *UserOps) UserDelete(id int)(int, error) {

	err := r.client.User.
		DeleteOneID(id).
		Exec(r.ctx)

	if err != nil {
		return 0, err
	}

	return id, nil
}
