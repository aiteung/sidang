package sidang

import (
	"database/sql"
	"fmt"
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

func GetPengujiSidang(db *mongo.Database, npm string) (sidang Sidang) {
	filter := bson.M{"npm": npm}
	sidang = atdb.GetOneDoc[Sidang](db, "sidang", filter)
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

func GetNPMfromHandphone(db *sql.DB, handphone string) (npm string) {
	q := "SELECT MhswId FROM simak_mst_mahasiswa WHERE handphone = %s"
	tsql := fmt.Sprintf(q, handphone)
	err := db.QueryRow(tsql).Scan(&npm)
	if err == sql.ErrNoRows {
		fmt.Printf("GetNPMfromHandphone, no user in table : simak_mst_mahasiswa")
	} else if err != nil {
		fmt.Printf("GetNPMfromHandphone: %v\n", err)
	}
	return
}

func GetTotalBimbinganfromNPM(db *sql.DB, npm string, tipebimbingan string) (jmlhpertemuan int) {
	q := "SELECT total_bimbingan FROM bimbingan_data WHERE npm = %s and tipe_bimbingan = '%s'"
	tsql := fmt.Sprintf(q, npm, tipebimbingan)
	err := db.QueryRow(tsql).Scan(&jmlhpertemuan)
	if err == sql.ErrNoRows {
		fmt.Printf("GetTotalBimbinganfromNPM, no user in table : simak_mst_mahasiswa")
	} else if err != nil {
		fmt.Printf("GetTotalBimbinganfromNPM: %v\n", err)
	}
	return
}
