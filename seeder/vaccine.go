package seeder

import (
	"context"
	"fmt"
	"log"
	"time"

	"cloud.google.com/go/firestore"
)

var (
	DataDistVaksin = []map[string]interface{}{
		{
			"id":          "mRwD8V7f3Fep",
			"sku":         "",
			"brand":       "Pfizer Inc.",
			"vaccinetype": "BNT162b2",
			"status":      1,
			"distributor": map[string]interface{}{
				"id":           DataDistributors[0]["id"],
				"company_name": DataDistributors[0]["company_name"],
			},
			"stock": map[string]interface{}{
				"available":     100000,
				"hold":          0,
				"allocated":     0,
				"on_allocation": 0,
			},
			"price": 85000,
		},
		{
			"id":          "X58fGNcATk0h",
			"sku":         "",
			"brand":       "Sinovac Biotech",
			"vaccinetype": "CoronaVac",
			"status":      1,
			"distributor": map[string]interface{}{
				"id":           DataDistributors[0]["id"],
				"company_name": DataDistributors[0]["company_name"],
			},
			"stock": map[string]interface{}{
				"available":     1500000,
				"hold":          0,
				"allocated":     0,
				"on_allocation": 0,
			},
			"price": 50000,
		},
		{
			"id":          "qm1tCv20G9QI",
			"sku":         "",
			"brand":       "Moderna",
			"vaccinetype": "mRNA-1273",
			"status":      1,
			"distributor": map[string]interface{}{
				"id":           DataDistributors[0]["id"],
				"company_name": DataDistributors[0]["company_name"],
			},
			"stock": map[string]interface{}{
				"available":     100000,
				"hold":          0,
				"allocated":     0,
				"on_allocation": 0,
			},
			"price": 90000,
		},

		// vaksin dari distributor B
		{
			"id":          "5e93aBzSFaNw",
			"sku":         "",
			"brand":       "Pfizer Inc.",
			"vaccinetype": "BNT162b2",
			"status":      1,
			"distributor": map[string]interface{}{
				"id":           DataDistributors[1]["id"],
				"company_name": DataDistributors[1]["company_name"],
			},
			"stock": map[string]interface{}{
				"available":     100000,
				"hold":          0,
				"allocated":     0,
				"on_allocation": 0,
			},
			"price": 83000,
		},
		{
			"id":          "j6bT5DQ2At8w",
			"sku":         "",
			"brand":       "Sinovac Biotech",
			"vaccinetype": "CoronaVac",
			"status":      1,
			"distributor": map[string]interface{}{
				"id":           DataDistributors[1]["id"],
				"company_name": DataDistributors[1]["company_name"],
			},
			"stock": map[string]interface{}{
				"available":     1500000,
				"hold":          0,
				"allocated":     0,
				"on_allocation": 0,
			},
			"price": 48000,
		},
		{
			"id":          "3jXgVo3Rgw88",
			"sku":         "",
			"brand":       "Moderna",
			"vaccinetype": "mRNA-1273",
			"status":      1,
			"distributor": map[string]interface{}{
				"id":           DataDistributors[1]["id"],
				"company_name": DataDistributors[1]["company_name"],
			},
			"stock": map[string]interface{}{
				"available":     100000,
				"hold":          0,
				"allocated":     0,
				"on_allocation": 0,
			},
			"price": 91000,
		},
	}

	dummyDate1, _ = time.Parse(time.RFC3339, "2020-12-01T19:00:00+07:00")
	dummyDate2, _ = time.Parse(time.RFC3339, "2020-12-03T10:00:00+07:00")

	DataDistPasokan = []map[string]interface{}{
		{
			"id":       "M2u4s14DMr2S",
			"batch_id": "2020-12-00001",
			"vaccine": map[string]interface{}{
				"id":          DataDistVaksin[0]["id"].(string),
				"sku":         DataDistVaksin[0]["sku"],
				"brand":       DataDistVaksin[0]["brand"],
				"vaccinetype": DataDistVaksin[0]["vaccinetype"],
			},
			"distributor": DataDistVaksin[0]["distributor"],
			"qty":         100000,
			"status":      1,
			"rcv_date":    dummyDate1,
			"act_date":    dummyDate2,
			"note":        "dummy data",
		},
		{
			"id":       "Ve8bs17P7H8k",
			"batch_id": "2020-12-00002",
			"vaccine": map[string]interface{}{
				"id":          DataDistVaksin[1]["id"].(string),
				"sku":         DataDistVaksin[1]["sku"],
				"brand":       DataDistVaksin[1]["brand"],
				"vaccinetype": DataDistVaksin[1]["vaccinetype"],
			},
			"distributor": DataDistVaksin[1]["distributor"],
			"qty":         1500000,
			"status":      1,
			"rcv_date":    dummyDate1,
			"act_date":    dummyDate2,
			"note":        "dummy data",
		},
		{
			"id":       "86I8FlrCJ5Yp",
			"batch_id": "2020-12-00003",
			"vaccine": map[string]interface{}{
				"id":          DataDistVaksin[2]["id"].(string),
				"sku":         DataDistVaksin[2]["sku"],
				"brand":       DataDistVaksin[2]["brand"],
				"vaccinetype": DataDistVaksin[2]["vaccinetype"],
			},
			"distributor": DataDistVaksin[2]["distributor"],
			"qty":         100000,
			"status":      1,
			"rcv_date":    dummyDate1,
			"act_date":    dummyDate2,
			"note":        "dummy data",
		},

		{
			"id":       "9QJAw0RjbTwJ",
			"batch_id": "2020-12-00004",
			"vaccine": map[string]interface{}{
				"id":          DataDistVaksin[3]["id"].(string),
				"sku":         DataDistVaksin[3]["sku"],
				"brand":       DataDistVaksin[3]["brand"],
				"vaccinetype": DataDistVaksin[3]["vaccinetype"],
			},
			"distributor": DataDistVaksin[3]["distributor"],
			"qty":         100000,
			"status":      1,
			"rcv_date":    dummyDate1,
			"act_date":    dummyDate2,
			"note":        "dummy data",
		},
		{
			"id":       "YP67f053igBT",
			"batch_id": "2020-12-00005",
			"vaccine": map[string]interface{}{
				"id":          DataDistVaksin[4]["id"].(string),
				"sku":         DataDistVaksin[4]["sku"],
				"brand":       DataDistVaksin[4]["brand"],
				"vaccinetype": DataDistVaksin[4]["vaccinetype"],
			},
			"distributor": DataDistVaksin[4]["distributor"],
			"qty":         1500000,
			"status":      1,
			"rcv_date":    dummyDate1,
			"act_date":    dummyDate2,
			"note":        "dummy data",
		},
		{
			"id":       "JA8HMqJfl8hd",
			"batch_id": "2020-12-00006",
			"vaccine": map[string]interface{}{
				"id":          DataDistVaksin[5]["id"].(string),
				"sku":         DataDistVaksin[5]["sku"],
				"brand":       DataDistVaksin[5]["brand"],
				"vaccinetype": DataDistVaksin[5]["vaccinetype"],
			},
			"distributor": DataDistVaksin[5]["distributor"],
			"qty":         100000,
			"status":      1,
			"rcv_date":    dummyDate1,
			"act_date":    dummyDate2,
			"note":        "dummy data",
		},
	}
)

func SeedDistributorVaccines(client *firestore.Client, ctx context.Context) (err error) {

	for _, v := range DataDistVaksin {
		doc := client.Doc(fmt.Sprintf("distributor_vaccines/%s", v["id"]))
		_, err = doc.Create(ctx, map[string]interface{}{
			"sku":         v["sku"],
			"brand":       v["brand"],
			"vaccinetype": v["vaccinetype"],
			"status":      v["status"],
			"distributor": v["distributor"],
			"stock":       v["stock"],
			"price":       v["price"],
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

	for _, v := range DataDistPasokan {
		doc := client.Doc(fmt.Sprintf("distributor_vaccine_stockin/%s", v["id"]))
		_, err = doc.Create(ctx, map[string]interface{}{
			"batch_id":    v["batch_id"],
			"vaccine":     v["vaccine"],
			"distributor": v["distributor"],
			"qty":         v["qty"],
			"status":      v["status"],
			"rcv_date":    v["rcv_date"],
			"act_date":    v["act_date"],
			"note":        v["note"],

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
