package main

import (
	"context"
	"myseeder/firestore/seeder"
	"os"

	"cloud.google.com/go/firestore"
)

func main() {
	ctx := context.Background()

	host := "0.0.0.0:8080"
	db := "spartech-infa-vaccine"

	os.Setenv("FIRESTORE_EMULATOR_HOST", host)

	client, err := firestore.NewClient(ctx, db)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	seeder.SeedDistributorPermissions(client, ctx)
	seeder.SeedDistributorRoles(client, ctx)
	seeder.SeedDistributorUser(client, ctx)

}
