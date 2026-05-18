package main

import "fmt"

func main() {
	// 1. Kısa Tanımlama (:=) - En çok bunu kullanacağız
	sunucuAdi := "Istanbul-Main-Auth" // Go bunun bir string (metin) olduğunu anlar
	port := 8080                      // Go bunun bir int (tam sayı) olduğunu anlar
	aktifMi := true                   // Go bunun bir bool (boolean) olduğunu anlar

	// 2. Değer atamadan tanımlama (Sıfır değerleri görmek için)
	var istekSayisi int
	var hataMesaji string

	// Ekrana formatlı yazdırma (Printf)
	fmt.Println("=== SUNUCU PROFİLİ ===")
	fmt.Printf("Sunucu Adı: %s\n", sunucuAdi) // %s -> string için
	fmt.Printf("Çalıştığı Port: %d\n", port)  // %d -> digit/int için
	fmt.Printf("Aktiflik Durumu: %t\n", aktifMi) // %t -> toggle/bool için

	fmt.Println("\n=== İLK DEĞERİ VERİLMEYENLER (Zero Values) ===")
	fmt.Printf("Gelen İstek Sayısı: %d\n", istekSayisi) // Çıktı: 0
	fmt.Printf("Son Hata Mesajı: '%s'\n", hataMesaji)    // Çıktı: '' (Boş metin)
}
