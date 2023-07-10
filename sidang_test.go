package sidang

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/aiteung/atdb"
)

var dbmongoinfo = atdb.DBInfo{
	DBString: os.Getenv("MONGOSTRING"),
	DBName:   "iteung",
}
var mongocon = atdb.MongoConnect(dbmongoinfo)

/* func TestInputData(m *testing.T) {
	judul1 := []string{"5 Tahap Membuat Dashboard Admin untuk Kemudahan Programmer dengan Reactjs dan Tailwindcss",
		"Analisis Hubungan Harga Bahan Bakar Terhadap Bahan Pangan Menggunakan Generalized Autoregressive Conditional Heteroscedasticity",
		"Analisis Sentimen Masyarakat Terhadap Kebijakan Polisi Tilang Manual di Indonesia",
		"Aplikasi Rekrutmen Karyawan dengan Alogaritme Artificial Neural Network (ANN) dan Framework Flask",
		"Cara Praktis Membuat Chatbot Whatsapp",
		"Deteksi Halaman Website Phising Menggunakan Algoritma Machine Learning Gradient Boosting Classifier",
		"Deteksi Spam Bot pada Komentar Youtube Menggunakan Artificial Neural Network(ANN)",
		"Deteksi Spam SMS Menggunakan Naive Bayes",
		"Helpdesk Ticketing System",
		"Hostpot Mikrotik API : Membuat Hotspot Mikrotik Menggunakan API",
		"Implementasi Algoritma C4.5 Untuk Memprediksi Kapasitas Penumpang Pesawat Menggunakan Python",
		"Klasifikasi Text Spam Menggunakan Metode Support Vector Machine dan Naive Bayes",
		"Kombinasi Golang, Mongodb dan Javascript Untuk Pengembangan Aplikasi Pengelolaan Koleksi Museum",
		"Komparasi Performa Model Terhadap Klasifikasi Sinyal MIT-BIH Arrhythmia Database",
		"Langkah Mudah Membuat Dashboard Rest API",
		"Membangun Aplikasi Prediksi Harga Token Kripto Menggunakan Python",
		"Membuat Analisis Komparatif Arima & Prophet pada Peramalan Penjualan",
		"Menerapkan Roblox di Metaverse",
		"Mouse Control Berbasis Eye Tracking Sebagai Alat Bantu Disabilitas",
		"Optimalisasi Perintah Suara sebagai Asisten Virtual",
		"Panduan Membuat Notifikasi Realtime Menggunakan Laravel Websocket dan React Js",
		"Panduan untuk Membuat Chatbot Cerdas (Sumber Elektronis) : Implementasi Openai di Telegram dan Discord",
		"Pembuatan Website Layanan Keuangan Dengan Metode Scrum",
		"Penerapan Metode Gamifikasi pada Rest Api Spring Boot",
		"Penerapan User Centered Design untuk Penentuan Sistem Contigency Plan",
		"Pengelompokan Kedisiplinan Pegawai Berdasarkan Absensi Menggunakan Algoritma K-Means",
		"Pengembangan Dashboard Laporan Bulanan untuk Monitoring Kinerja Perusahaan Menggunakan User Centered Design",
		"Pengembangan Smart Conveyor Dengan Arduino (Menggunakan GPS Tracking Berbasis Android)",
		"Penggunaan TAWK.TO untuk Implementasi Chatting Menggunakan Library Javascript Vanila JS",
		"Telegram Bot Wawancara Kerja Dengan Alogaritma Long Short Term Memory",
		"Toko Online Dengan Laravel dan Vue Js (Sumber Elektronis) : Implementasi Rest API Web Service",
		"Tutorial Membuat Game Obby",
		"Tutorial Membuat Game Puzzle Roblox",
		"Voice Cloning : Membuat Sendiri Suara Artifisial Metode Sequence To Sequence Speech Synthesis",
	}
	surat1 := Skt{
		Nomor: "3771/PUS.01.01/PDPBP",
		Judul: judul1,
	}

	atdb.InsertOneDoc(mongocon, "skt", surat1)
	judul2 := []string{"Peramalan Waktu Pengiriman Outbound Menggunakan Random Forest Regressor"}
	surat2 := Skt{
		Nomor: "4381/PUS.01.01/PDPBP",
		Judul: judul2,
	}
	atdb.InsertOneDoc(mongocon, "skt", surat2)

} */

func TestGetData(m *testing.T) {
	hasil := GetSkt(mongocon, "4381/PUS.01.01/PDPBP")
	fmt.Println(reflect.DeepEqual(hasil, Skt{}))
}

func TestIsValid(m *testing.T) {
	hasil := IsValidSKTNumber(mongocon, "4381/PUS.01.01/PDPBP")
	fmt.Println(hasil)
}
