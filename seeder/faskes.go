package seeder

import (
	"context"
	"fmt"
	"log"
	"myseeder/firestore/util"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/gosimple/slug"
)

var (

	// ################################# FASKES #######################################

	DataFaskesPermits = []map[string]interface{}{
		// STATISTIK
		{
			"label":  "View statistik",
			"object": "statistik",
			"action": "view",
		},
		{
			"label":  "Export statistik",
			"object": "statistik",
			"action": "export",
		},

		// APPOINTMENT
		{
			"label":  "View appointment",
			"object": "appointment",
			"action": "view",
		},
		{
			"label":  "Create appointment",
			"object": "appointment",
			"action": "create",
		},
		{
			"label":  "Update / Change appointment",
			"object": "appointment",
			"action": "update",
		},
		{
			"label":  "Export appointment",
			"object": "appointment",
			"action": "export",
		},

		// ALOKASI (order faskes -> distributor)
		{
			"label":  "View alokasi",
			"object": "alokasi",
			"action": "view",
		},
		{
			"label":  "Create alokasi",
			"object": "alokasi",
			"action": "create",
		},
		{
			"label":  "Update / Change alokasi",
			"object": "alokasi",
			"action": "update",
		},
		{
			"label":  "Export alokasi",
			"object": "alokasi",
			"action": "export",
		},

		// ORDER (pasien -> faskes)
		{
			"label":  "View order",
			"object": "order",
			"action": "view",
		},
		{
			"label":  "Create order",
			"object": "order",
			"action": "create",
		},
		{
			"label":  "Update / Change order",
			"object": "order",
			"action": "update",
		},
		{
			"label":  "Export order",
			"object": "order",
			"action": "export",
		},

		// PAKET MGT
		{
			"label":  "View paket",
			"object": "paket",
			"action": "view",
		},
		{
			"label":  "Create paket",
			"object": "paket",
			"action": "create",
		},
		{
			"label":  "Update paket",
			"object": "paket",
			"action": "update",
		},
		{
			"label":  "Export paket",
			"object": "paket",
			"action": "export",
		},

		// REPORT
		{
			"label":  "View report",
			"object": "report",
			"action": "view",
		},
		{
			"label":  "Export report",
			"object": "report",
			"action": "export",
		},
	}

	DataFaskesRoles = []map[string]interface{}{
		{
			"label":       "Pengurus",
			"description": "Superuser role",
			"status":      1,
			"permissions": []string{"*|*"},
		},
		{
			"label":       "Administratif",
			"description": "Operational/administratif role",
			"status":      1,
			"permissions": []string{"statistik|view", "alokasi|view", "appointment|view", "appointment|update", "order|view", "paket|view", "report|view"},
		},
		{
			"label":       "View Only",
			"description": "Viewer only role",
			"status":      1,
			"permissions": []string{"statistik|view", "alokasi|view", "appointment|view", "order|view", "paket|view", "report|view"},
		},
	}

	DataFaskesType = []map[string]interface{}{
		{"label": "Rumah Sakit Tipe A"},
		{"label": "Rumah Sakit Tipe B"},
		{"label": "Rumah Sakit Tipe C"},
		{"label": "Rumah Sakit Tipe D"},
		{"label": "Klinik Pratama"},
		{"label": "Klinik Utama"},
		{"label": "Dokter Praktek Perorangan"},
	}

	RumahSakitTipeA = slug.Make(DataFaskesType[0]["label"].(string))
	RumahSakitTipeB = slug.Make(DataFaskesType[1]["label"].(string))
	RumahSakitTipeC = slug.Make(DataFaskesType[2]["label"].(string))
	RumahSakitTipeD = slug.Make(DataFaskesType[3]["label"].(string))
	KlinikPratama   = slug.Make(DataFaskesType[4]["label"].(string))
	KlinikUtama     = slug.Make(DataFaskesType[5]["label"].(string))
	PrakterDokter   = slug.Make(DataFaskesType[6]["label"].(string))

	DataFaskes = []map[string]interface{}{
		{
			"id":       "3Rev3Lo46Rrg",
			"label":    "Rumah Sakit Tipe A - Bandung",
			"medftype": RumahSakitTipeA,
			"distributor": map[string]interface{}{
				"id":           "distributor-a",
				"company_name": "Distributor A",
			},
			"address":       "Jl Otto Iskandardinata Psr Baru Bl D Lt D-2/5,Kebon Jeruk",
			"kelurahan":     "Kebon Jeruk",
			"kecamatan":     "Andir",
			"city":          "Bandung",
			"country":       "Indonesia",
			"postalcode":    "",
			"location":      []float32{-7.025253, 107.519760},
			"status":        1,
			"website":       "http://alterra.id",
			"contact_email": []string{"div1@email.xyz", "div2@email.xyz"},
			"phonenumber":   []string{"+62211230001", "+6281100001"},
			"fax_no":        []string{"+62211230002", "+62211230003"},
			"whatsapp_no":   []string{"+6281100002", "+6281100002"},
		},
		{
			"id":       "mc2YQDA4Tq1o",
			"label":    "Rumah Sakit Tipe B - Bandung",
			"medftype": RumahSakitTipeB,
			"distributor": map[string]interface{}{
				"id":           "distributor-a",
				"company_name": "Distributor A",
			},
			"address":       "Jl Hegarmanah Tgh 43",
			"kelurahan":     "Hegarmanah",
			"kecamatan":     "Hergamanah",
			"city":          "Bandung",
			"country":       "Indonesia",
			"postalcode":    "",
			"location":      []float32{-7.025253, 107.519760},
			"status":        1,
			"website":       "http://alterra.id",
			"contact_email": []string{"div1@email.xyz", "div2@email.xyz"},
			"phonenumber":   []string{"+62211230001", "+6281100001"},
			"fax_no":        []string{"+62211230002", "+62211230003"},
			"whatsapp_no":   []string{"+6281100002", "+6281100002"},
		},
		{
			"id":       "S1ITwgvo9fwg",
			"label":    "Rumah Sakit Tipe A - Jakarta",
			"medftype": RumahSakitTipeA,
			"distributor": map[string]interface{}{
				"id":           "distributor-b",
				"company_name": "Distributor B",
			},
			"address":       "Jl MH Thamrin Kav 28-30 Plaza Indonesia Lt 1/140",
			"kelurahan":     "Menteng",
			"kecamatan":     "Menteng",
			"city":          "Jakarta",
			"country":       "Indonesia",
			"postalcode":    "",
			"location":      []float32{-7.025253, 107.519760},
			"status":        1,
			"website":       "http://alterra.id",
			"contact_email": []string{"div1@email.xyz", "div2@email.xyz"},
			"phonenumber":   []string{"+62211230001", "+6281100001"},
			"fax_no":        []string{"+62211230002", "+62211230003"},
			"whatsapp_no":   []string{"+6281100002", "+6281100002"},
		},
		{
			"id":       "ZaN2mCX86myZ",
			"label":    "Klinik Pratama - Jakarta",
			"medftype": KlinikPratama,
			"distributor": map[string]interface{}{
				"id":           "distributor-b",
				"company_name": "Distributor B",
			},
			"address":       "Jl Pegangsaan Tmr 1-A Wisma Financial Manulife Indonesia Lt 1-8",
			"kelurahan":     "Pegangsaan",
			"kecamatan":     "Pegangsaan",
			"city":          "Jakarta",
			"country":       "Indonesia",
			"postalcode":    "",
			"location":      []float32{-7.025253, 107.519760},
			"status":        1,
			"website":       "http://alterra.id",
			"contact_email": []string{"div1@email.xyz", "div2@email.xyz"},
			"phonenumber":   []string{"+62211230001", "+6281100001"},
			"fax_no":        []string{"+62211230002", "+62211230003"},
			"whatsapp_no":   []string{"+6281100002", "+6281100002"},
		},
	}

	DataFaskesUsers = []map[string]interface{}{
		{
			"username":    "pengurus@rs-a-bandung.xyz",
			"email":       "pengurus@rs-a-bandung.xyz",
			"phonenumber": "+628001110001",
			"medfac": map[string]interface{}{
				"id":           DataFaskes[0]["id"].(string),
				"company_name": DataFaskes[0]["label"].(string),
			},
			"password": util.GenHashedPass("password"),
			// "password_salt": "xxxxxxxxxxxx",
			"fullname": "Pengurus di RS A Bandung",
			"status":   1,
			"avatar":   "https://images.pexels.com/photos/36469/woman-person-flowers-wreaths.jpg?auto=compress&cs=tinysrgb&crop=faces&fit=crop&h=100&w=100",
			"roles": []string{
				"pengurus",
			},
		},
		{
			"username":    "suster@rs-a-bandung.xyz",
			"email":       "suster@rs-a-bandung.xyz",
			"phonenumber": "+628001110003",
			"distributor": map[string]interface{}{
				"id":           DataFaskes[0]["id"].(string),
				"company_name": DataFaskes[0]["label"].(string),
			},
			"password": util.GenHashedPass("password"),
			// "password_salt": "xxxxxxxxxxxx",
			"fullname": "Suster/Administratif di RS A Bandung",
			"status":   1,
			"avatar":   "https://uifaces.co/our-content/donated/N8kxcjRw.jpg",
			"roles": []string{
				"administratif",
			},
		},
		{
			"username":    "viewonly@rs-a-bandung.xyz",
			"email":       "viewonly@rs-a-bandung.xyz",
			"phonenumber": "+628001110004",
			"distributor": map[string]interface{}{
				"id":           DataFaskes[0]["id"].(string),
				"company_name": DataFaskes[0]["label"].(string),
			},
			"password": util.GenHashedPass("password"),
			// "password_salt": "xxxxxxxxxxxx",
			"fullname": "Viewonly di RS A Bandung",
			"status":   1,
			"avatar":   "https://uifaces.co/our-content/donated/wteXSjwk.jpg",
			"roles": []string{
				"view-only",
			},
		},
		{
			"username":    "tidakaktif@rs-a-bandung.xyz",
			"email":       "tidakaktif@rs-a-bandung.xyz",
			"phonenumber": "+628001110005",
			"distributor": map[string]interface{}{
				"id":           DataFaskes[0]["id"].(string),
				"company_name": DataFaskes[0]["label"].(string),
			},
			"password": util.GenHashedPass("password"),
			// "password_salt": "xxxxxxxxxxxx",
			"fullname": "Tidak Aktif di RS A Bandung",
			"status":   0,
			"avatar":   "https://images.pexels.com/photos/412840/pexels-photo-412840.jpeg?crop=faces&fit=crop&h=200&w=200&auto=compress&cs=tinysrgb",
			"roles": []string{
				"pengurus",
			},
		},

		{
			"username":    "pengurus@rs-a-jakarta.xyz",
			"email":       "pengurus@rs-a-jakarta.xyz",
			"phonenumber": "+628001110001",
			"medfac": map[string]interface{}{
				"id":           DataFaskes[2]["id"].(string),
				"company_name": DataFaskes[2]["label"].(string),
			},
			"password": util.GenHashedPass("password"),
			// "password_salt": "xxxxxxxxxxxx",
			"fullname": "Pengurus di RS A Jakarta",
			"status":   1,
			"avatar":   "https://uifaces.co/our-content/donated/noplz47r59v1uxvyg8ku.png",
			"roles": []string{
				"pengurus",
			},
		},
		{
			"username":    "suster@rs-a-jakarta.xyz",
			"email":       "suster@rs-a-jakarta.xyz",
			"phonenumber": "+628001110003",
			"distributor": map[string]interface{}{
				"id":           DataFaskes[2]["id"].(string),
				"company_name": DataFaskes[2]["label"].(string),
			},
			"password": util.GenHashedPass("password"),
			// "password_salt": "xxxxxxxxxxxx",
			"fullname": "Suster/Administratif di RS A Jakarta",
			"status":   1,
			"avatar":   "https://images.pexels.com/photos/412840/pexels-photo-412840.jpeg?crop=faces&fit=crop&h=200&w=200&auto=compress&cs=tinysrgb",
			"roles": []string{
				"administratif",
			},
		},
		{
			"username":    "viewonly@rs-a-jakarta.xyz",
			"email":       "viewonly@rs-a-jakarta.xyz",
			"phonenumber": "+628001110004",
			"distributor": map[string]interface{}{
				"id":           DataFaskes[2]["id"].(string),
				"company_name": DataFaskes[2]["label"].(string),
			},
			"password": util.GenHashedPass("password"),
			// "password_salt": "xxxxxxxxxxxx",
			"fullname": "Viewonly di RS A Jakarta",
			"status":   1,
			"avatar":   "https://uifaces.co/our-content/donated/wteXSjwk.jpg",
			"roles": []string{
				"view-only",
			},
		},
		{
			"username":    "tidakaktif@rs-a-jakarta.xyz",
			"email":       "tidakaktif@rs-a-jakarta.xyz",
			"phonenumber": "+628001110005",
			"distributor": map[string]interface{}{
				"id":           DataFaskes[2]["id"].(string),
				"company_name": DataFaskes[2]["label"].(string),
			},
			"password": util.GenHashedPass("password"),
			// "password_salt": "xxxxxxxxxxxx",
			"fullname": "Tidak Aktif di RS A Jakarta",
			"status":   0,
			"avatar":   "https://images.pexels.com/photos/36469/woman-person-flowers-wreaths.jpg?auto=compress&cs=tinysrgb&crop=faces&fit=crop&h=100&w=100",
			"roles": []string{
				"pengurus",
			},
		},
	}
)

func SeedFaskesData(client *firestore.Client, ctx context.Context) (err error) {

	for _, v := range DataFaskes {
		doc := client.Doc("medfacs/" + fmt.Sprintf("%s", v["id"]))
		_, err = doc.Create(ctx, map[string]interface{}{
			"label":         v["label"],
			"medftype":      v["medftype"],
			"distributor":   v["distributor"],
			"address":       v["address"],
			"kelurahan":     v["kelurahan"],
			"kecamatan":     v["kecamatan"],
			"city":          v["city"],
			"country":       v["country"],
			"postalcode":    v["postalcode"],
			"location":      v["location"],
			"status":        v["status"],
			"website":       v["website"],
			"contact_email": v["contact_email"],
			"phonenumber":   v["phonenumber"],
			"fax_no":        v["fax_no"],
			"whatsapp_no":   v["whatsapp_no"],

			// default
			"created_at": time.Now(),
			"created_by": map[string]interface{}{
				"id":      "system",
				"name":    "system",
				"colname": "operator_users",
			},
			"updated_at": nil,
			"updated_by": map[string]interface{}{
				"id":      nil,
				"name":    nil,
				"colname": "operator_users",
			},
			"deleted_at": nil,
			"deleted_by": map[string]interface{}{
				"id":      nil,
				"name":    nil,
				"colname": "operator_users",
			},
		})

		if err != nil {
			log.Fatalf("Failed adding: %v", err)
			return
		}
	}

	for _, v := range DataFaskesPermits {
		doc := client.Doc("medfac_permissions/" + fmt.Sprintf("%s|%s", v["object"], v["action"]))
		_, err = doc.Create(ctx, map[string]interface{}{
			"label":  v["label"],
			"object": v["object"],
			"action": v["action"],

			// default
			"created_at": time.Now(),
			"created_by": map[string]interface{}{
				"id":      "system",
				"name":    "system",
				"colname": "operator_users",
			},
			"updated_at": nil,
			"updated_by": map[string]interface{}{
				"id":      nil,
				"name":    nil,
				"colname": "operator_users",
			},
			"deleted_at": nil,
			"deleted_by": map[string]interface{}{
				"id":      nil,
				"name":    nil,
				"colname": "operator_users",
			},
		})

		if err != nil {
			log.Fatalf("Failed adding: %v", err)
			return
		}
	}

	for _, v := range DataFaskesRoles {
		doc := client.Doc("medfac_roles/" + slug.Make(v["label"].(string)))
		_, err = doc.Create(ctx, map[string]interface{}{
			"label":       v["label"],
			"description": v["description"],
			"status":      v["status"],
			"permissions": v["permissions"],

			// default
			"created_at": time.Now(),
			"created_by": map[string]interface{}{
				"id":      "system",
				"name":    "system",
				"colname": "operator_users",
			},
			"updated_at": nil,
			"updated_by": map[string]interface{}{
				"id":      nil,
				"name":    nil,
				"colname": "operator_users",
			},
			"deleted_at": nil,
			"deleted_by": map[string]interface{}{
				"id":      nil,
				"name":    nil,
				"colname": "operator_users",
			},
		})

		if err != nil {
			log.Fatalf("Failed adding: %v", err)
			return
		}
	}

	for _, v := range DataFaskesUsers {
		doc := client.Doc("medfac_users/" + v["username"].(string))
		_, err = doc.Create(ctx, map[string]interface{}{
			"username":      v["username"],
			"email":         v["email"],
			"phonenumber":   v["phonenumber"],
			"medfac":        v["medfac"],
			"password":      v["password"],
			"password_salt": v["password_salt"],
			"fullname":      v["fullname"],
			"status":        v["status"],
			"avatar":        v["avatar"],
			"roles":         v["roles"],

			// default
			"created_at": time.Now(),
			"created_by": map[string]interface{}{
				"id":      "system",
				"name":    "system",
				"colname": "operator_users",
			},
			"updated_at": nil,
			"updated_by": map[string]interface{}{
				"id":      nil,
				"name":    nil,
				"colname": "operator_users",
			},
			"deleted_at": nil,
			"deleted_by": map[string]interface{}{
				"id":      nil,
				"name":    nil,
				"colname": "operator_users",
			},
		})

		if err != nil {
			log.Fatalf("Failed adding: %v", err)
			return
		}
	}

	return
}
