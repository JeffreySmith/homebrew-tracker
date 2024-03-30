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
	"fmt"
	"time"
	//"net/http"
)

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
	beer_map[1].SetStart(2023, 04, 22)
	beer_map[1].SetBottle(2023, 9, 29)
	beer_map[1].Print()
	age, err := beer_map[1].Age()
	if err == nil {
		plural := "days"
		if age == 1 {
			plural = "day"
		}
		fmt.Printf("Your batch is %v %v old\n", age, plural)
	}
	bottle_age, err := beer_map[1].BottleAge()
	if err == nil {
		plural := "days"
		if bottle_age == 1 {
			plural = "day"
		}
		date := beer_map[1].BottleDate.Format("2006-01-02")
		fmt.Printf("Your batch was bottled on %v and has been in the bottle for %v %v\n", date, bottle_age, plural)
	}
}
