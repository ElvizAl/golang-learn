package log

import "github.com/sirupsen/logrus"

// Mendeklarasikan variabel global Logger yang akan digunakan untuk logging di seluruh aplikasi
var Logger *logrus.Logger

// Fungsi SetupLogger untuk mengonfigurasi logger menggunakan logrus
func SetupLogger() {
	// Membuat instance baru dari logrus Logger
	log := logrus.New()

	// Menetapkan format log dengan Timestamp penuh dan warna di konsol
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true, // Menampilkan timestamp lengkap (tanggal dan waktu)
		ForceColors:   true, // Memaksa tampilan warna di log jika mendukung terminal
	})

	// Mencetak pesan log dengan level Info yang menunjukkan bahwa logger telah diinisialisasi
	log.Info("Logged initiated using logrus!")

	// Menyimpan instance log yang telah dikonfigurasi ke dalam variabel global Logger
	Logger = log
}
