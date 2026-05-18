package main
import "fmt"
func topla(sayi1 int,sayi2 int) int   {
return sayi1+sayi2 } 
func sunucukontrol(port int) (string,bool) {
	if port == 8080 {
		return "Port 8080 meşgul! APİ sunucusu zaten çalişiyor.", true}
		return "Port temiz,kullanilabilir" , false }
	 
func main(){
	toplam := topla(40,2)
	fmt.Printf("Toplam sonucu: %d\n",toplam)
   fmt.Println("------------------")

   mesaj, durum :=sunucukontrol(8080)
   fmt.Printf("sistem mesaji= %s\n",mesaj)
   fmt.Printf("port dolumu?: %t\n",durum )
 }  

