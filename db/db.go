package db

import (
	"context"
	"entgo/ent"
	"entgo/util"
	"fmt"
)

func CreateUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	u, err := client.User.
		Create().
		SetAge(30).
		SetName("테스트").
		SetPNumber("01011112222").
		Save(ctx)
	util.HandleError(err)

	fmt.Println("user was created :", u)

	return u, nil
}
