package main

import (
	"sort"
	"strconv"
)

func sumSliseFloat(slise []float64) (sum float64) {
	sum = 0.
	for i := 0; i < len(slise); i++ {
		sum += slise[i]
	}
	return
}

func findOutTheOptimalModel1(s float64) (float64, []Item) {
	const sliseSize = 10
	var allSquares = getData("solar_panel", "square")
	var allUtilityCoefficient = getData("solar_panel", "utility_coefficient")
	var allNames = getData("solar_panel", "name")
	//postData("solar_panel", "name", "efficiency", "voltage", "price", "square", "'ferst h'", "0.05", "342342", "423", "4324")
	//update("solar_panel", "2", "name", "voltage", "'todo'", "8912")
	//delete("solar_panel", "4")

	var data = []Item{}

	for i := 0; i < sliseSize; i++ {

		var item Item
		var utilityCoefficient, err = strconv.ParseFloat(allUtilityCoefficient[i], 64)
		var square, err1 = strconv.ParseFloat(allSquares[i], 64)

		item.Value = utilityCoefficient
		item.Weight = square
		item.Name = allNames[i]

		if err != nil {
			panic(err)
		}

		if err1 != nil {
			panic(err)
		}

		data = append(data, item)
	}

	sort.Slice(data, func(i, j int) bool {
		return data[i].Value > data[j].Value
	})

	prevSlise := [sliseSize]float64{}
	currentSlise := [sliseSize]float64{}

	for i := 0; i < sliseSize; i++ {
		var j = 0
		for s > sumSliseFloat(currentSlise[:j]) {

		}
	}

	return 5., data
}
