package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type deck []string

func newDeck() deck{
	cards := deck{}
	cardSuits :=[]string {"Spades","Hearts","Clubs","Diamonds"}
	cardValues :=[]string {"Aces","Two","Three","Four"}
	for _,suit := range cardSuits{
		for _,value := range cardValues{
			cards = append(cards,value +" " + "of" + " " +suit)
		}
	}
	return cards
}

func (d deck) print(){
	for i,card := range d {
		fmt.Println(i,card)
	}
}

func (d deck) toString() string {
	return strings.Join((d),",")
}

func (d deck) saveToFile(filename string) error {
	 return ioutil.WriteFile(filename,[]byte(d.toString()),0666)
}

// package main

// func main() {
// 	cards := newDeck()
// 	cards.saveToFile("my_Cards")
	
// }
