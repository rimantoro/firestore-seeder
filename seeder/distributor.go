package seeder

import (
	"context"
	"log"
	"myseeder/firestore/util"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/gosimple/slug"
)

var (
	DocDistRoles  []*firestore.DocumentRef
	DocDistPermit []*firestore.DocumentRef
)

func SeedDistributorPermissions(client *firestore.Client, ctx context.Context) (err error) {
	rows := []map[string]interface{}{
		{
			"id":     "alokasi|view",
			"label":  "Alokasi - View",
			"object": "alokasi",
			"action": "view",
		},
		{
			"id":     "alokasi|write",
			"label":  "Alokasi - Write",
			"object": "alokasi",
			"action": "write",
		},
		{
			"id":     "alokasi|export",
			"label":  "Alokasi - Export",
			"object": "alokasi",
			"action": "export",
		},
		{
			"id":     "vaksin|view",
			"label":  "Vaksin - View",
			"object": "vaksin",
			"action": "view",
		},
		{
			"id":     "vaksin|write",
			"label":  "Vaksin - Write",
			"object": "vaksin",
			"action": "write",
		},
		{
			"id":     "vaksin|export",
			"label":  "Vaksin - Export",
			"object": "vaksin",
			"action": "export",
		},
	}

	for _, v := range rows {
		doc := client.Doc("distributor_permissions/" + v["id"].(string))
		_, err = doc.Create(ctx, map[string]interface{}{
			"label":  v["label"],
			"object": v["object"],
			"action": v["action"],
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

		DocDistRoles = append(DocDistRoles, doc)

		if err != nil {
			log.Fatalf("Failed adding: %v", err)
			return
		}
	}

	return
}

func SeedDistributorRoles(client *firestore.Client, ctx context.Context) (err error) {
	rows := []map[string]interface{}{
		{
			"label":       "Administrator",
			"description": "Superuser role",
			"status":      1,
			"permissions": []string{"*|*"},
		},
		{
			"label":       "Ops",
			"description": "Operational role",
			"status":      1,
			"permissions": []string{"alokasi|view", "alokasi|write", "alokasi|export", "vaksin|view", "vaksin|write", "vaksin|export"},
		},
		{
			"label":       "View Only",
			"description": "Viewer only role",
			"status":      1,
			"permissions": []string{"alokasi|view", "vaksin|view"},
		},
	}

	for _, v := range rows {
		doc := client.Doc("distributor_roles/" + slug.Make(v["label"].(string)))
		_, err = doc.Create(ctx, map[string]interface{}{
			"label":       rows[0]["label"],
			"description": rows[0]["description"],
			"status":      rows[0]["status"],
			"permissions": rows[0]["permissions"],
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

		DocDistPermit = append(DocDistPermit, doc)
	}

	return
}

func SeedDistributorDistributors(client *firestore.Client, ctx context.Context) (err error) {
	rows := []map[string]interface{}{
		{
			"company_name":  "Distributor A",
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
			"company_name":  "Distributor B",
			"address":       "Jl Perjuangan Bl DH/88",
			"kelurahan":     "",
			"kecamatan":     "",
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

	for _, v := range rows {
		doc := client.Doc("distributors/" + slug.Make(v["company_name"].(string)))
		_, err = doc.Create(ctx, map[string]interface{}{
			"company_name":  v["company_name"],
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

func SeedDistributorUser(client *firestore.Client, ctx context.Context) (err error) {

	// get specified role (by implicit ID) query (collection_name/{id})
	docRoleAdmin := client.Doc("distributor_roles/administrator")
	docRoleOPS := client.Doc("distributor_roles/ops")
	docRoleViewer := client.Doc("distributor_roles/view_only")

	rows := []map[string]interface{}{
		{
			"username":    "admin@company.xyz",
			"email":       "admin@company.xyz",
			"phonenumber": "+628001110001",
			"distributor": map[string]interface{}{
				"id":           "distributor-a",
				"company_name": "Distributor A",
			},
			"password": util.GenHashedPass("password"),
			// "password_salt": "xxxxxxxxxxxx",
			"fullname": "User Dummy Administrator",
			"status":   1,
			"roles": []string{
				docRoleAdmin.ID,
			},
		},
		{
			"username":    "ops@company.xyz",
			"email":       "ops@company.xyz",
			"phonenumber": "+628001110002",
			"distributor": map[string]interface{}{
				"id":           "distributor-a",
				"company_name": "Distributor A",
			},
			"password": util.GenHashedPass("password"),
			// "password_salt": "xxxxxxxxxxxx",
			"fullname": "User Dummy for Distributor Ops",
			"status":   1,
			"roles": []string{
				docRoleOPS.ID,
			},
		},
		{
			"username":    "view@company.xyz",
			"email":       "view@company.xyz",
			"phonenumber": "+628001110003",
			"distributor": map[string]interface{}{
				"id":           "distributor-a",
				"company_name": "Distributor A",
			},
			"password": util.GenHashedPass("password"),
			// "password_salt": "xxxxxxxxxxxx",
			"fullname": "Dummy for view only user",
			"status":   1,
			"roles": []string{
				docRoleViewer.ID,
			},
		},
		{
			"username":    "inactive@company.xyz",
			"email":       "inactive@company.xyz",
			"phonenumber": "+628001110004",
			"distributor": map[string]interface{}{
				"id":           "distributor-a",
				"company_name": "Distributor A",
			},
			"password": util.GenHashedPass("password"),
			// "password_salt": "xxxxxxxxxxxx",
			"fullname": "Dummy for inactive user",
			"status":   0,
			"roles": []string{
				docRoleAdmin.ID,
			},
		},
	}

	for _, v := range rows {
		doc := client.Doc("distributor_users/" + v["username"].(string))
		_, err = doc.Create(ctx, map[string]interface{}{
			"username":      v["username"],
			"email":         v["email"],
			"phonenumber":   v["phonenumber"],
			"distributor":   v["distributor"],
			"password":      v["password"],
			"password_salt": v["password_salt"],
			"fullname":      v["fullname"],
			"status":        v["status"],
			"roles":         v["roles"],
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

		// DocDistPermit = append(DocDistPermit, doc)
	}

	return
}
