// You can edit this code!
// Click here and start typing.
package main

import "fmt"
//BAŞARILI BİR UYGULAMA
func main() {
	midtermExam := 70
	finalExam := 30
	var midtermAndFinalCount int
	isMakeUpExam := false
	midtermAndFinalCount = Control(midtermExam, finalExam, isMakeUpExam)
	fmt.Println(midtermAndFinalCount)

}

//TODO
//golang dilinde eğer isimler büyük harf ile başlarsa 
//bu paketin dışında bu değerin kullanılacağı anlamıda gelir henüz bu konuyu işlemedik not olarak yazıyorum
func Control(midtermExam int, finalExam int, isMakeUpExam bool) int {

	if isMakeUpExam {
		makeUpExamCount := getMakeUpExam()
		return makeUpExam(midtermExam, makeUpExamCount)
	} else {
		midterm := (midtermExam * 40) / 100
		final := (finalExam * 60) / 100
		return midterm + final
	}

}
func makeUpExam(midTernExam int, makeUpExamCount int) int {
	midtermCount := (midTernExam * 40) / 100
	makeUpCount := (makeUpExamCount * 60) / 100
	return midtermCount + makeUpCount
}
func getMakeUpExam() int {
	return 62
}
