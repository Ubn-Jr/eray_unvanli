package main

import "fmt"

func main() {
	var kilolar [9]int = [9]int{500, 35, 59, 12, 107, 25, 100, 63, 74}
	asansöreBinenKisilerinKilosu(kilolar)
}

func asansöreBinenKisilerinKilosu(kilolar [9]int) {
	var maksKapasite int = 500
	var toplamAgirlik int = 0
	var asansoreAlinamiyan = 0
	var asansoreBinenKisiSayisi = 0
	//TODO
	// 9 yerine gerekli arrayın len() ile uzunluğunu kullanalım
	for x := 0; x < len(kilolar); x++ {

		var sonrakiAgirlik int = kilolar[x]

		//TODO
		// if else if else yapısıyla kullanım daha uygun olur burada
		//
		if maksKapasite == toplamAgirlik {
			fmt.Println(len(kilolar)-asansoreBinenKisiSayisi, "kişi asansöre alınamadı")
			break

		} else if (maksKapasite - toplamAgirlik) >= sonrakiAgirlik {
			toplamAgirlik += sonrakiAgirlik
			fmt.Println(toplamAgirlik)
			asansoreBinenKisiSayisi++
		} else {
			fmt.Println(kilolar[x], "Asansör yük kapasitesini geçtiği için alamadık")
			asansoreAlinamiyan++
		}

	}
	if asansoreAlinamiyan != 0 {
		fmt.Println(asansoreAlinamiyan, "kişi asansöre alınamadı")
	}
}
