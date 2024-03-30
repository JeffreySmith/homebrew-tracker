/*
BSD 3-Clause License

Copyright (c) 2024, Jeffrey Smith

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

1. Redistributions of source code must retain the above copyright notice, this
   list of conditions and the following disclaimer.

2. Redistributions in binary form must reproduce the above copyright notice,
   this list of conditions and the following disclaimer in the documentation
   and/or other materials provided with the distribution.

3. Neither the name of the copyright holder nor the names of its
   contributors may be used to endorse or promote products derived from
   this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/

package main

import (
	"errors"
	"fmt"

	"time"
	//"net/http"
)

type Beer struct {
	id         int
	name       string
	rating     int
	OG         float32
	FG         float32
	startDate  time.Time
	bottleDate time.Time
}

func NewBeer(id int, name string, OG float32, FG float32, rating int) *Beer {
	return &Beer{
		id:         id,
		name:       name,
		OG:         OG,
		FG:         FG,
		rating:     rating,
		startDate:  time.Date(1950, 01, 01, 12, 0, 0, 0, time.Local),
		bottleDate: time.Date(1950, 01, 01, 12, 0, 0, 0, time.Local),
	}
}
func (b *Beer) Update(name string, OG float32, FG float32, rating int) {
	b.name = name
	b.OG = OG
	b.FG = FG
	b.rating = rating
}
func (b *Beer) GetABV() (float32, error) {
	if b.OG < 0 || b.FG < 0 {
		return -1, errors.New("Starting Gravity or Final Gravity is not valid")
	}
	return (b.OG - b.FG) * 131.25, nil
}
func (b *Beer) setStart(year int, month time.Month, day int) {
	b.startDate = time.Date(year, month, day, 12, 0, 0, 0, time.Local)
}
func (b *Beer) setBottle(year int, month time.Month, day int) {
	b.bottleDate = time.Date(year, month, day, 12, 0, 0, 0, time.Local)
}

func (b *Beer) Age() (int, error) {
	now := time.Now()
	days := now.Sub(b.startDate).Hours() / 24
	if days < 0 {
		return -1, errors.New("Your batch is from the future")
	}
	return int(days), nil
}

func (b *Beer) BottleAge() (int,error) {
	now := time.Now()
	days := now.Sub(b.bottleDate).Hours() / 24
	if days < 0 {
		return -1, errors.New("Your batch is from the future")
	}
	return int(days), nil
}

func PushBeer(m map[int]*Beer, beer *Beer) {
	m[beer.id] = beer
}
func (b *Beer) Print() {
	var name string
	var rating string
	abv, e := b.GetABV()

	if b.name == "" {
		name = "None"
	} else {
		name = b.name
	}
	if b.rating <= 0 {
		rating = "Unrated"
	} else {
		rating = fmt.Sprintf("%v/10", b.rating)
	}
	if e != nil {
		fmt.Printf("Name: \"%v\", Rating: %v\n", name, rating)
	} else {
		fmt.Printf("Name: \"%v\", ABV: %.3v%%, Rating: %v\n", name, abv, rating)
	}
	if b.startDate.Year() > 1950 {
		fmt.Printf("Start Date: %v\n", b.startDate.Format("2006-01-02"))
	}
	if b.bottleDate.Year() > 1950 {
		fmt.Printf("Start Date: %v\n", b.bottleDate.Format("2006-01-02"))
	}

}

func main() {

	beer_map := make(map[int]*Beer)
	year := time.Now().Year()
	month := time.Now().Month()
	day := time.Now().Day()

	fmt.Printf("The current date is: %d-%d-%d\n", day, month, year)

	my_beer := NewBeer(1, "Furious Violin", 1.055, -1, 6)

	PushBeer(beer_map, my_beer)
	beer_map[1].Print()
	beer_map[1].Update("Furious Violin", 1.055, 1.015, 7)
	beer_map[1].setStart(2023, 04, 22)
	beer_map[1].setBottle(2024, 03, 29)
	beer_map[1].Print()
	age, err := beer_map[1].Age()
	if err == nil {
		plural := "days"
		if age == 1{
			plural = "day"
		}
		fmt.Printf("Your batch is %v %v old\n", age,plural)
	}
	bottle_age, err := beer_map[1].BottleAge()
	if err == nil {
		plural := "days"
		if bottle_age == 1{
			plural = "day"
		}
		fmt.Printf("Your batch has been in the bottle for %v %v\n",bottle_age,plural)
	}
}
