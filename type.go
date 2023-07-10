package sidang

type Skt struct {
	Nomor string   `json:"nomor,omitempty" bson:"nomor,omitempty"`
	Judul []string `json:"judul,omitempty" bson:"judul,omitempty"`
}

type Sidang struct {
	NPM        string `json:"npm,omitempty" bson:"npm,omitempty"`
	Nama       string `json:"nama,omitempty" bson:"nama,omitempty"`
	Penguji    string `json:"penguji,omitempty" bson:"penguji,omitempty"`
	Pembimbing string `json:"pembimbing,omitempty" bson:"pembimbing,omitempty"`
}
