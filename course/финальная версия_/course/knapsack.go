package main

import (
	"sort"
	"strconv"
)

type Item struct {
	Name   string
	Value  float64
	Weight float64
}

type SolatPanel struct {
	id                 string
	Square             float64
	UtilityCoefficient float64
	Quantity           int64
	Name               string
}

const conuntOfItem = 120

func findOutTheOptimalModel(s float64) []SolatPanel {
	var allSquares = getData("solar_panel", "square")
	var allUtilityCoefficient = getData("solar_panel", "utility_coefficient")
	var allNames = getData("solar_panel", "name")
	var allQuantity = getData("solar_panel", "quantity")
	var allId = getData("solar_panel", "id")

	//postData("solar_panel", "name", "efficiency", "voltage", "price", "square", "'ferst h'", "0.05", "342342", "423", "4324")
	//update("solar_panel", "2", "name", "voltage", "'todo'", "8912")
	//delete("solar_panel", "4")

	var dataSet DataSet
	var solarPanels = []SolatPanel{}
	var resultSolarPanel = []SolatPanel{}

	for i := 0; i < conuntOfItem; i++ {
		var solarPanel SolatPanel
		var utilityCoefficient, _ = strconv.ParseFloat(allUtilityCoefficient[i], 64)
		var square, _ = strconv.ParseFloat(allSquares[i], 64)
		var quantity, _ = strconv.ParseInt(allQuantity[i], 10, 64)
		var id = allId[i]

		solarPanel.Name = allNames[i]
		solarPanel.Quantity = quantity
		solarPanel.Square = square
		solarPanel.UtilityCoefficient = utilityCoefficient
		solarPanel.id = id

		solarPanels = append(solarPanels, solarPanel)
	}

	sort.Slice(solarPanels, func(i, j int) bool {
		return solarPanels[i].UtilityCoefficient > solarPanels[j].UtilityCoefficient
	})

	var i = 0
	var count = 0
	for count < conuntOfItem {
		for j := 0; j < int(solarPanels[i].Quantity); j++ {
			if count == conuntOfItem {
				break
			}

			dataSet.values = append(dataSet.values, solarPanels[j].UtilityCoefficient)
			dataSet.weights = append(dataSet.weights, solarPanels[j].Square)

			count++
		}
		i++
	}

	dataSet.limitedWeight = s

	var res = Backtracking(dataSet)

	for i := 0; i < conuntOfItem; i++ {
		if res.solution[i] == '1' {
			resultSolarPanel = append(resultSolarPanel, solarPanels[i])
		}
	}

	return resultSolarPanel
}
