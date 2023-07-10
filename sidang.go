package sidang

import (
	"reflect"

	"github.com/aiteung/atdb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetSkt(db *mongo.Database, nomor string) (usercred Skt) {
	filter := bson.M{"nomor": nomor}
	usercred = atdb.GetOneDoc[Skt](db, "skt", filter)
	return
}

func IsValidSKTNumber(db *mongo.Database, nomor string) bool {
	hasil := GetSkt(db, nomor)
	if reflect.DeepEqual(hasil, Skt{}) {
		return false
	} else {
		return true
	}
}
