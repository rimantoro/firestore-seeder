package seeder

import (
	"context"
	"fmt"
	"log"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/gosimple/slug"
)

func SeedDistributorVaccines(client *firestore.Client, ctx context.Context) (err error) {

	docDistA := client.Doc("distributors/distributor-a")
	docSnap, _ := docDistA.Get(ctx)
	data := docSnap.Data()

	rows := []map[string]interface{}{
		{
			"sku":         "???",
			"brand":       "Pfizer Inc.",
			"vaccinetype": "BNT162b2",
			"status":      1,
			"distributor": map[string]interface{}{
				"id":           docDistA.ID,
				"company_name": data["company_name"],
			},
			"stock": map[string]interface{}{
				"available": 100000,
				"hold":      0,
			},
		},
		{
			"sku":         "???",
			"brand":       "Sinovac Biotech",
			"vaccinetype": "CoronaVac",
			"status":      1,
			"distributor": map[string]interface{}{
				"id":           docDistA.ID,
				"company_name": data["company_name"],
			},
			"stock": map[string]interface{}{
				"available": 500000,
				"hold":      0,
			},
		},
		{
			"sku":         "???",
			"brand":       "Moderna",
			"vaccinetype": "mRNA-1273",
			"status":      1,
			"distributor": map[string]interface{}{
				"id":           docDistA.ID,
				"company_name": data["company_name"],
			},
			"stock": map[string]interface{}{
				"available": 500000,
				"hold":      0,
			},
		},
	}

	for _, v := range rows {
		// "distributor_vaccines/" + slug.Make(v["brand"].(string))
		// doc := client.Doc(fmt.Sprintf("distributor_vaccines/%s/%s", slug.Make(v["brand"].(string)), slug.Make(v["vaccinetype"].(string))))
		doc := client.Doc(fmt.Sprintf("distributor_vaccines/%s", slug.Make(v["brand"].(string)+"-"+v["vaccinetype"].(string))))
		_, err = doc.Create(ctx, map[string]interface{}{
			"sku":         v["sku"],
			"brand":       v["brand"],
			"vaccinetype": v["vaccinetype"],
			"status":      v["status"],
			"distributor": v["distributor"],
			"stock":       v["stock"],
			// default
			"created_at": time.Now(),
			"created_by": map[string]interface{}{
				"id":      "system",
				"name":    "system",
				"colname": "distributor_users",
			},
			"updated_at": nil,
			"updated_by": map[string]interface{}{
				"id":      nil,
				"name":    nil,
				"colname": "distributor_users",
			},
			"deleted_at": nil,
			"deleted_by": map[string]interface{}{
				"id":      nil,
				"name":    nil,
				"colname": "distributor_users",
			},
		})

		if err != nil {
			log.Fatalf("Failed adding: %v", err)
			return
		}
	}

	return
}
