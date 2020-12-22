package main

import (
	"context"
	"myseeder/firestore/seeder"
	"os"

	"cloud.google.com/go/firestore"
)

func main() {
	ctx := context.Background()

	// for local
	db := "spartech-infa-vaccine"
	os.Setenv("FIRESTORE_EMULATOR_HOST", "0.0.0.0:8080")

	// // for cloud
	// db := "spartech-282005"
	// os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "./dev-auth.json")

	client, err := firestore.NewClient(ctx, db)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	seeder.SeedDistributorPermissions(client, ctx)
	seeder.SeedDistributorRoles(client, ctx)
	seeder.SeedDistributorDistributors(client, ctx)
	seeder.SeedDistributorUser(client, ctx)
	seeder.SeedDistributorVaccines(client, ctx)
	seeder.SeedDistributorAlokasi(client, ctx)

	seeder.SeedFaskesData(client, ctx)
	seeder.SeedOperatorData(client, ctx)

}
