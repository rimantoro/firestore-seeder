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
	DataDistPermits = []map[string]interface{}{
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

		// PASOKAN
		{
			"label":  "View pasokan",
			"object": "pasokan",
			"action": "view",
		},
		{
			"label":  "Create pasokan",
			"object": "pasokan",
			"action": "create",
		},
		{
			"label":  "Update / Change pasokan",
			"object": "pasokan",
			"action": "update",
		},
		{
			"label":  "Export pasokan",
			"object": "pasokan",
			"action": "export",
		},

		// ALOKASI
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

		// ORDER (FROM FASKES)
		{
			"label":  "View faskes order",
			"object": "faskes-order",
			"action": "view",
		},
		{
			"label":  "Export faskes order",
			"object": "faskes-order",
			"action": "export",
		},

		// FASKES MGT
		{
			"label":  "View faskes",
			"object": "faskes",
			"action": "view",
		},
		{
			"label":  "Create faskes",
			"object": "faskes",
			"action": "create",
		},
		{
			"label":  "Update faskes",
			"object": "faskes",
			"action": "update",
		},
		{
			"label":  "Export faskes",
			"object": "faskes",
			"action": "export",
		},

		// VAKSIN MGT
		{
			"label":  "View vaksin",
			"object": "vaksin",
			"action": "view",
		},
		{
			"label":  "Create vaksin",
			"object": "vaksin",
			"action": "create",
		},
		{
			"label":  "Update vaksin",
			"object": "vaksin",
			"action": "update",
		},
		{
			"label":  "Export vaksin",
			"object": "vaksin",
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

	DataDistRoles = []map[string]interface{}{
		{
			"label":       "Pengurus",
			"description": "Superuser role",
			"status":      1,
			"permissions": []string{"*|*"},
		},
		{
			"label":       "View Only",
			"description": "Viewer only role",
			"status":      1,
			"permissions": []string{"statistik|view", "alokasi|view", "vaksin|view", "pasokan|view", "faskes|view", "faskes-order|view", "report|view"},
		},
	}

	DataDistributors = []map[string]interface{}{
		{
			"id":            slug.Make("Distributor A"),
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
			"id":            slug.Make("Distributor B"),
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

	DataDistUsers = []map[string]interface{}{
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
			"avatar":   "https://images.generated.photos/PgakVwN6rVSolTj8D0052vofGLCCzpZ8wSF_LLcU5f4/rs:fit:512:512/Z3M6Ly9nZW5lcmF0/ZWQtcGhvdG9zL3Yy/XzA1OTYyNjguanBn.jpg",
			"roles": []string{
				"pengurus",
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
			"avatar":   "https://randomuser.me/api/portraits/men/55.jpg",
			"roles": []string{
				"view-only",
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
			"avatar":   "https://uifaces.co/our-content/donated/hvaUVob5.jpg",
			"roles": []string{
				"perngurus",
			},
		},
	}

	DocDistRoles  []*firestore.DocumentRef
	DocDistPermit []*firestore.DocumentRef
)

func SeedDistributorPermissions(client *firestore.Client, ctx context.Context) (err error) {
	for _, v := range DataDistPermits {
		doc := client.Doc("distributor_permissions/" + fmt.Sprintf("%s|%s", v["object"], v["action"]))
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
	for _, v := range DataDistRoles {
		doc := client.Doc("distributor_roles/" + slug.Make(v["label"].(string)))
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
	for _, v := range DataDistributors {
		doc := client.Doc("distributors/" + v["id"].(string))
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

	for _, v := range DataDistUsers {
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
			"avatar":        v["avatar"],
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
	}

	return
}
