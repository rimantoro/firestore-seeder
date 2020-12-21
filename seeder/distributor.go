package seeder

import (
	"context"
	"fmt"
	"log"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/gosimple/slug"
)

var (
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
