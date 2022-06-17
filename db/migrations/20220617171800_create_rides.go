// 20220617171800_create_rides

package migrations

import (
	"github.com/go-rel/rel"
)

// MigrateCreateRides definition
func MigrateCreateRides(schema *rel.Schema) {
	schema.CreateTable("rides", func(t *rel.Table) {
		t.ID("id")
		t.DateTime("created_at")
		t.DateTime("updated_at")
		t.Int("price")
		t.String("user_id")
		t.String("vehicle_id")
		t.Bool("finished")
	})
}

// RollbackCreateRides definition
func RollbackCreateRides(schema *rel.Schema) {
	schema.DropTable("rides")
}
