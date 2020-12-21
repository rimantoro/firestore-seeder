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
	DataOpsPermits = []map[string]interface{}{
		// FASKES MANAGEMENT
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
			"label":  "Update / Change faskes",
			"object": "faskes",
			"action": "update",
		},
		{
			"label":  "Export faskes",
			"object": "faskes",
			"action": "export",
		},

		// DISTRIBUTOR MANAGEMENT
		{
			"label":  "View distributor",
			"object": "distributor",
			"action": "view",
		},
		{
			"label":  "Create distributor",
			"object": "distributor",
			"action": "create",
		},
		{
			"label":  "Update / Change distributor",
			"object": "distributor",
			"action": "update",
		},
		{
			"label":  "Export distributor",
			"object": "distributor",
			"action": "export",
		},

		// OPERATOR MANAGEMENT
		{
			"label":  "View operator",
			"object": "operator",
			"action": "view",
		},
		{
			"label":  "Create operator",
			"object": "operator",
			"action": "create",
		},
		{
			"label":  "Update / Change operator",
			"object": "operator",
			"action": "update",
		},
		{
			"label":  "Export operator",
			"object": "operator",
			"action": "export",
		},

		// ORDER  (PASIEN -> FASKES)
		{
			"label":  "View order pasien - faskes",
			"object": "order",
			"action": "view",
		},
		{
			"label":  "Export order pasien - faskes",
			"object": "order",
			"action": "export",
		},

		// ALOKASI (FASKES -> DISTRIBUTOR)
		{
			"label":  "View alokasi faskes - distributor",
			"object": "alokasi",
			"action": "view",
		},
		{
			"label":  "Export alokasi faskes - distributor",
			"object": "alokasi",
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

	DataOpsRoles = []map[string]interface{}{
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
			"permissions": []string{"faskes|view", "distributor|view", "operator|view", "order|view", "alokasi|view", "report|view"},
		},
	}

	DataOpsUsers = []map[string]interface{}{
		{
			"username":    "pengurus@operator.xyz",
			"email":       "pengurus@operator.xyz",
			"phonenumber": "+628001110001",
			"password":    util.GenHashedPass("password"),
			// "password_salt": "xxxxxxxxxxxx",
			"fullname": "Operator Superuser",
			"status":   1,
			"roles": []string{
				"pengurus",
			},
		},
		{
			"username":    "viewer@operator.xyz",
			"email":       "viewer@operator.xyz",
			"phonenumber": "+628001110003",
			"password":    util.GenHashedPass("password"),
			// "password_salt": "xxxxxxxxxxxx",
			"fullname": "Operator Viewer",
			"status":   1,
			"roles": []string{
				"view-only",
			},
		},
		{
			"username":    "inactive@operator.xyz",
			"email":       "inactive@operator.xyz",
			"phonenumber": "+628001110004",
			"password":    util.GenHashedPass("password"),
			// "password_salt": "xxxxxxxxxxxx",
			"fullname": "Operator Inactive",
			"status":   0,
			"roles": []string{
				"pengurus",
			},
		},
	}
)

func SeedOperatorData(client *firestore.Client, ctx context.Context) (err error) {
	for _, v := range DataOpsPermits {
		doc := client.Doc("ops_permissions/" + fmt.Sprintf("%s|%s", v["object"], v["action"]))
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

	for _, v := range DataOpsRoles {
		doc := client.Doc("ops_roles/" + slug.Make(v["label"].(string)))
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

	for _, v := range DataOpsUsers {
		doc := client.Doc("ops_users/" + v["username"].(string))
		_, err = doc.Create(ctx, map[string]interface{}{
			"username":      v["username"],
			"email":         v["email"],
			"phonenumber":   v["phonenumber"],
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
