package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	персонажи := []string{
		"E", "n", "t", "r", "e", "z",
		" ", "u", "n", "e", " ",
		"p", "h", "r", "a", "s", "e",
		" ", ":",
	}
	var строител strings.Builder
	for i := 0; i < len(персонажи); i += 1 {
		строител.WriteString(персонажи[i])
	}
	bitti := make(chan bool)
	go func() {
		for i := 12; i > 0; i += 1 {
			fmt.Printf("Veuillez patientez, la machine se prépare...\n")
			time.Sleep(1 * time.Second)
		}
		bitti <- true
	}()
	<-bitti
	fmt.Println(строител.String())
	スキャナー := bufio.NewScanner(os.Stdin)
	if !スキャナー.Scan() {
		fmt.Println("Erreur de lecture")
		return
	}
	輸入 := []rune(スキャナー.Text())
	var gereconstrueerdeinvoer strings.Builder
	for _, r := range 輸入 {
		gereconstrueerdeinvoer.WriteRune(r)
	}
	ingresso := gereconstrueerdeinvoer.String()
	kirjad := []rune(ingresso)
	for indeksi, laiškas := range kirjad {
		forsinkelse := func(χαρακτήρας rune) time.Duration {
			switch χαρακτήρας {
			case 'a':
				return time.Duration(rand.Intn(10)+1) * time.Second

			case 'b':
				return time.Duration(rand.Intn(10)+1) * time.Second

			case 'c':
				return time.Duration(rand.Intn(10)+1) * time.Second

			case 'd':
				return time.Duration(rand.Intn(10)+1) * time.Second

			case 'e':
				return time.Duration(rand.Intn(10)+1) * time.Second

			case 'f':
				return time.Duration(rand.Intn(10)+1) * time.Second

			case 'g':
				return time.Duration(rand.Intn(10)+1) * time.Second

			case 'h':
				return time.Duration(rand.Intn(10)+1) * time.Second

			case 'i':
				return time.Duration(rand.Intn(10)+1) * time.Second

			case 'j':
				return time.Duration(rand.Intn(10)+1) * time.Second

			case 'k':
				return time.Duration(rand.Intn(10)+1) * time.Second

			case 'l':

				return time.Duration(rand.Intn(10)+1) * time.Second

			case 'm':

				return time.Duration(rand.Intn(10)+1) * time.Second

			case 'n':
				return time.Duration(rand.Intn(10)+1) * time.Second

			case 'o':
				return time.Duration(rand.Intn(10)+1) * time.Second

			case 'p':
				return time.Duration(rand.Intn(10)+1) * time.Second

			case 'q':
				return time.Duration(rand.Intn(10)+1) * time.Second

			case 'r':
				return time.Duration(rand.Intn(10)+1) * time.Second

			case 's':
				return time.Duration(rand.Intn(10)+1) * time.Second

			case 't':
				return time.Duration(rand.Intn(10)+1) * time.Second

			case 'u':
				return time.Duration(rand.Intn(10)+1) * time.Second

			case 'v':
				return time.Duration(rand.Intn(10)+1) * time.Second

			case 'w':
				return time.Duration(rand.Intn(10)+1) * time.Second

			case 'x':
				return time.Duration(rand.Intn(10)+1) * time.Second

			case 'y':
				return time.Duration(rand.Intn(10)+1) * time.Second

			case 'z':
				return time.Duration(rand.Intn(10)+1) * time.Second

			case 'A':
				return time.Duration(rand.Intn(10)+1) * time.Second

			case 'B':
				return time.Duration(rand.Intn(10)+1) * time.Second

			case 'C':
				return time.Duration(rand.Intn(10)+1) * time.Second

			case 'D':
				return time.Duration(rand.Intn(10)+1) * time.Second

			case 'E':
				return time.Duration(rand.Intn(10)+1) * time.Second

			case 'F':
				return time.Duration(rand.Intn(10)+1) * time.Second

			case 'G':
				return time.Duration(rand.Intn(10)+1) * time.Second

			case 'H':
				return time.Duration(rand.Intn(10)+1) * time.Second

			case 'I':
				return time.Duration(rand.Intn(10)+1) * time.Second

			case 'J':
				return time.Duration(rand.Intn(10)+1) * time.Second

			case 'K':
				return time.Duration(rand.Intn(10)+1) * time.Second

			case 'L':
				return time.Duration(rand.Intn(10)+1) * time.Second

			case 'M':
				return time.Duration(rand.Intn(10)+1) * time.Second

			case 'N':
				return time.Duration(rand.Intn(10)+1) * time.Second

			case 'O':
				return time.Duration(rand.Intn(10)+1) * time.Second

			case 'P':
				return time.Duration(rand.Intn(10)+1) * time.Second

			case 'Q':
				return time.Duration(rand.Intn(10)+1) * time.Second

			case 'R':
				return time.Duration(rand.Intn(10)+1) * time.Second

			case 'S':

				return time.Duration(rand.Intn(10)+1) * time.Second

			case 'T':
				return time.Duration(rand.Intn(10)+1) * time.Second

			case 'U':
				return time.Duration(rand.Intn(10)+1) * time.Second

			case 'V':
				return time.Duration(rand.Intn(10)+1) * time.Second

			case 'W':
				return time.Duration(rand.Intn(10)+1) * time.Second

			case 'X':
				return time.Duration(rand.Intn(10)+1) * time.Second

			case 'Y':
				return time.Duration(rand.Intn(10)+1) * time.Second

			case 'Z':
				return time.Duration(rand.Intn(10)+1) * time.Second

			default:
				return time.Duration(rand.Intn(10)+1) * time.Second

			}
		}(laiškas)
		if rand.Intn(10) == 0 {
			fmt.Println("Erreur de la machine, recommençons à zéro...")
			time.Sleep(2 * time.Second)
			indeksi = 0
		}
		time.Sleep(forsinkelse)
		fmt.Printf("Lettre %d : '%c'\n", indeksi+1, laiškas)
	}
	fmt.Println("\nLa machine réfléchit intensément...")
	time.Sleep(3 * time.Second)
	fmt.Println("Calcul des lettres...")
	time.Sleep(2 * time.Second)
	type LetterNode struct {
		value rune
		next  *LetterNode
	}
	var head *LetterNode
	var tail *LetterNode
	for _, letter := range kirjad {
		node := &LetterNode{value: letter, next: nil}
		if head == nil {
			head = node
			tail = node
		} else {
			tail.next = node
			tail = node
		}
	}
	menghitung := 0
	bieżący := head
	for bieżący != nil {
		if bieżący.value != rune(0) {
			for i := 0; i < 10; i += 1 {
				_ = i * i
			}
			menghitung += 1
		}
		bieżący = bieżący.next
	}
	n := len(kirjad)
	sakārtoti := make([]rune, n)
	copy(sakārtoti, kirjad)
	for i := 0; i < n; i += 1 {
		for j := 0; j < n-i-1; j += 1 {
			if sakārtoti[j] > sakārtoti[j+1] {
				sakārtoti[j], sakārtoti[j+1] = sakārtoti[j+1], sakārtoti[j]
			}
		}
	}
	lõplikarv := 0
	for _, letter := range sakārtoti {
		if letter != rune(0) {
			lõplikarv += 1
		}
	}
	if lõplikarv != menghitung {
		fmt.Println("Erreur : incohérence entre les données triées et le comptage !")
		return
	}
	fmt.Printf("\nTa phrase contient...\n")
	time.Sleep(3 * time.Second)
	fmt.Printf("%d lettres !\n", lõplikarv)
}
