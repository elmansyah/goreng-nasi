package goreng_nasi

import "time"

func GorengNasi(name string) string {
	return "Goreng nasi rasa " + name
}

func JumlahNasi() {
	for i := 0; i < 1000; i++ {
		go GorengNasi("Pedas")
		time.Sleep(10 * time.Second)
	}
}
