package sidang

type Skt struct {
	Nomor string   `json:"nomor,omitempty" bson:"nomor,omitempty"`
	Judul []string `json:"judul,omitempty" bson:"judul,omitempty"`
}

type Mhsdt struct {
	MhswId    string
	Nama      string
	Handphone string
}
