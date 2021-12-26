package twelve

import (
    "fmt"
)
const prefix = "On the %s day of Christmas my true love gave to me: %s"

type verse struct {
    day string
    present string
}

var verses = []verse {
    {"first",  "a Partridge in a Pear Tree."},
    {"second",  "two Turtle Doves, and "},
    {"third",   "three French Hens, "},
    {"fourth",   "four Calling Birds, "},
    {"fifth",   "five Gold Rings, "},
    {"sixth",   "six Geese-a-Laying, "},
    {"seventh", "seven Swans-a-Swimming, "},
    {"eighth",  "eight Maids-a-Milking, "},
    {"ninth",  "nine Ladies Dancing, "},
    {"tenth",   "ten Lords-a-Leaping, "},
    {"eleventh", "eleven Pipers Piping, "},
    {"twelfth",  "twelve Drummers Drumming, "},
}

func Verse(i int) string {
    if i < 1 || i > 12 {
        return "Invalid input"
    }
	i--
    day := verses[i].day
    presents := verses[i].present
    i--
    for i >= 0 {
        presents += verses[i].present
        i--
    }
	return fmt.Sprintf(prefix, day, presents)
}

func Song() string {
    song := Verse(1)
    for i := 2 ; i <= 12 ; i++ {
        song += "\n" + Verse(i)
    }
	return song
}

