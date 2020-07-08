package main

func newCard() string {
	return "Five of Diamonds"
}

func main() {

	cards := newDeck()
	cards.suffle()
	cards.print()

	///////////////dasar/////////////////////
	// var mycard string = "tesing"
	// nycard2 := "testing2"
	// cards := []string{"Ace of Diamonds", newCard()}
	// cards := deck{"Ace of Diamonds", newCard()}
	//cards = append(cards, "Six of Spades")
	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }
	// cards.print()
	// fmt.Println(cards[2:])
	///////////////partition/////////////////
	// cards := newDeck()
	// hand, remainingDeck := deal(cards, 5)
	// hand.print()
	// remainingDeck.print()
	//////////string to []byte///////////////
	// greeting := "Hi there!"
	// fmt.Println([]byte(greeting))
	/////////input- output in drive//////////
	// cards := newDeck()
	// // fmt.Println(cards.toString())
	// cards.saveToFile("my_cards")
	// cardsFile := newDeckFromFile("my_cards")
	// cardsFile.print()
	/////////////Simple random//////////////
	// test := []int{1, 2, 3, 4, 5}
	// fmt.Println(test)
	// for i := range test {

	// 	newPositon := rand.Intn(len(test) - 1)
	// 	fmt.Print("i = ", i, ", random = ", newPositon, " | test[", i, "] = ", test[i], ", test[", newPositon, "] = ", test[newPositon])
	// 	test[i], test[newPositon] = test[newPositon], test[i]
	// 	fmt.Println(" ========>", " test[", i, "] = ", test[i], ", test[", newPositon, "] = ", test[newPositon])
	// }
	// fmt.Println(test)
	/////////////////////////////////////////
	/////////////////////////////////////////
}
