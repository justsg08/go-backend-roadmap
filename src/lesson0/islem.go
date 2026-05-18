package main

import "fmt"

func main() {
	// Değiştirilemeyen sabit bir KDV oranı tanımlıyoruz
	const kdvOrani = 0.20

	urunAdi := "Premium Go Kursu"
	hamFiyat := 150.00 // float64 (ondalıklı sayı)

	kdvTutari := hamFiyat * kdvOrani
	toplamFiyat := hamFiyat + kdvTutari

	fmt.Println("------ FATURA DETAYI ------")
	fmt.Printf("Ürün: %s\n", urunAdi)
	fmt.Printf("Net Fiyat: %.2f TL\n", hamFiyat) // %.2f -> Virgülden sonra 2 basamak göster
	fmt.Printf("KDV Tutarı (%%20): %.2f TL\n", kdvTutari)
	fmt.Printf("Müşteriden Tahsil Edilen: %.2f TL\n", toplamFiyat)
	fmt.Println("---------------------------")
}
