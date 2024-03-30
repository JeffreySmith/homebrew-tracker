/*
BSD 3-Clause License

# Copyright (c) 2024, Jeffrey Smith

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
)

type Beer struct {
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	Rating     int       `json:"rating"`
	OG         float32   `json:"OG"`
	FG         float32   `json:"FG"`
	StartDate  time.Time `json:"start_date"`
	BottleDate time.Time `json:"bottle_date"`
}

func NewBeer(id int, name string, OG float32, FG float32, rating int) *Beer {
	return &Beer{
		Id:         id,
		Name:       name,
		OG:         OG,
		FG:         FG,
		Rating:     rating,
		StartDate:  time.Date(1950, 01, 01, 12, 0, 0, 0, time.Local),
		BottleDate: time.Date(1950, 01, 01, 12, 0, 0, 0, time.Local),
	}
}
func (b *Beer) Update(name string, OG float32, FG float32, rating int) {
	b.Name = name
	b.OG = OG
	b.FG = FG
	b.Rating = rating
}
func (b *Beer) GetABV() (float32, error) {
	if b.OG < 0 || b.FG < 0 {
		return -1, errors.New("Starting Gravity or Final Gravity is not valid")
	}
	return (b.OG - b.FG) * 131.25, nil
}
func (b *Beer) SetStart(year int, month time.Month, day int) {
	b.StartDate = time.Date(year, month, day, 12, 0, 0, 0, time.Local)
}
func (b *Beer) SetBottle(year int, month time.Month, day int) {
	b.BottleDate = time.Date(year, month, day, 12, 0, 0, 0, time.Local)
}

func (b *Beer) Age() (int, error) {
	now := time.Now()
	days := now.Sub(b.StartDate).Hours() / 24
	if days < 0 {
		return -1, errors.New("Your batch is from the future")
	}
	return int(days), nil
}

func (b *Beer) BottleAge() (int, error) {
	now := time.Now()
	days := now.Sub(b.BottleDate).Hours() / 24
	if days < 0 {
		return -1, errors.New("Your batch is from the future")
	}
	return int(days), nil
}

func PushBeer(m map[int]*Beer, beer *Beer) {
	m[beer.Id] = beer
}
func (b *Beer) Print() {
	var name string
	var rating string
	abv, e := b.GetABV()

	if b.Name == "" {
		name = "None"
	} else {
		name = b.Name
	}
	if b.Rating <= 0 {
		rating = "Unrated"
	} else {
		rating = fmt.Sprintf("%v/10", b.Rating)
	}
	if e != nil {
		fmt.Printf("Name: \"%v\", Rating: %v\n", name, rating)
	} else {
		fmt.Printf("Name: \"%v\", ABV: %.3v%%, Rating: %v\n", name, abv, rating)
	}
	if b.StartDate.Year() > 1950 {
		fmt.Printf("Start Date: %v\n", b.StartDate.Format("2006-01-02"))
	}
	if b.BottleDate.Year() > 1950 {
		fmt.Printf("Bottling Date: %v\n", b.BottleDate.Format("2006-01-02"))
	}

}
