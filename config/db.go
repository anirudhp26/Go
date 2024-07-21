package config

import (
	"fmt"
	"hello/prisma/db"
)

func SetupClient() (*db.PrismaClient, error) {
	client := db.NewClient()
	fmt.Println(("New Client Created"))
	if err := client.Prisma.Connect(); err != nil {
		return nil, err
	}
	return client, nil
}
