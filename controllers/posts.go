package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"hello/prisma/db"
)

func CreatePost(c *db.PrismaClient) (*db.PostModel, error) {
	client := c
	ctx := context.Background()
	createdPost, err := client.Post.CreateOne(
		db.Post.Title.Set("QUYWEI!"),
		db.Post.Published.Set(true),
		db.Post.Desc.Set("kdfskfhjjk"),
	).Exec(ctx)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	result, _ := json.MarshalIndent(createdPost, "", "  ")
	fmt.Printf("created post: %s\n", result)
	return createdPost, nil
}
