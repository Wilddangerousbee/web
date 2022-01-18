package main

import (
	"fmt"
	"strconv"
)

func getIdTargetPerfectCombinaton(s float64) string {
	postData("perfect_combination", "square", fmt.Sprintf("%f", s))
	idPerfectCombinaton := getData("perfect_combination", "id")

	return idPerfectCombinaton[len(idPerfectCombinaton)-1]
}

func getSolarPanelsFromDataBase(id string) (solarPanels []SolatPanel) {

	var square = getData("solar_panel s join perfect_combination_solar_panel p on s.id = p.solar_panel_id where p.perfect_combination_id = "+id, "square")
	var utilityCoefficient = getData("solar_panel s join perfect_combination_solar_panel p on s.id = p.solar_panel_id where p.perfect_combination_id = "+id, "utility_coefficient")
	var name = getData("solar_panel s join perfect_combination_solar_panel p on s.id = p.solar_panel_id where p.perfect_combination_id = "+id, "name")
	var quantity = getData("solar_panel s join perfect_combination_solar_panel p on s.id = p.solar_panel_id where p.perfect_combination_id = "+id, "quantity")
	var idS = getData("solar_panel s join perfect_combination_solar_panel p on s.id = p.solar_panel_id where p.perfect_combination_id =  "+id, "id")

	for i := 0; i < len(idS); i++ {
		var solarPanel SolatPanel
		var floatUtilityCoefficient, _ = strconv.ParseFloat(utilityCoefficient[i], 64)
		var floatSquare, _ = strconv.ParseFloat(square[i], 64)
		var intQuantity, _ = strconv.ParseInt(quantity[i], 10, 64)
		var id = idS[i]

		solarPanel.Name = name[i]
		solarPanel.Quantity = intQuantity
		solarPanel.Square = floatSquare
		solarPanel.UtilityCoefficient = floatUtilityCoefficient
		solarPanel.id = id

		solarPanels = append(solarPanels, solarPanel)
	}
	return
}

func changeRecord(s float64, id string) (flUpdate bool) {
	flUpdate = false
	createdDateSolarPanels := getData("solar_panel", "created_date")
	createdDatePerfectCombinaton := getData("perfect_combination where id = "+id, "created_date")

	for _, createdDateSolarPanel := range createdDateSolarPanels {
		if createdDateSolarPanel > createdDatePerfectCombinaton[0] {
			flUpdate = true
		}
	}
	return
}

func deletePerfectCombination(id string) {
	delete("perfect_combination_solar_panel", "perfect_combination_id = "+id)
	delete("perfect_combination", "id = "+id)
}

func updatePerfectCombination(s float64) (solarPanels []SolatPanel) {
	squarePerfectCombinaton := getData("perfect_combination", "square")
	idPerfectCombinaton := getData("perfect_combination", "id")

	for i := 0; i < len(squarePerfectCombinaton); i++ {
		var floatSquarePerfectCombinaton, _ = strconv.ParseFloat(squarePerfectCombinaton[i], 64)
		if floatSquarePerfectCombinaton == s {
			flChangeRecord := changeRecord(s, idPerfectCombinaton[i])
			if flChangeRecord {
				deletePerfectCombination(idPerfectCombinaton[i])
			}
			fmt.Printf("вызов функции обработки существующей записи" + idPerfectCombinaton[i])
			if !flChangeRecord {
				solarPanels = getSolarPanelsFromDataBase(idPerfectCombinaton[i])
				return
			}
		}
	}

	solarPanels = findOutTheOptimalModel(s)
	idTargetPerfectCombinaton := getIdTargetPerfectCombinaton(s)

	for _, solarPanel := range solarPanels {
		postData("perfect_combination_solar_panel", "solar_panel_id", "perfect_combination_id", solarPanel.id, idTargetPerfectCombinaton)
	}
	return
}
