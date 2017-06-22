package htmlOperations

import "github.com/willcruse/ComputingCoursework/dbOperations"

func GenerateNewTable(uID int) string {
	rowStrings := dbOperations.GetSets(uID)
	htmlOutput := "<table><tr><th>Sets</th></tr>"

	for i := 0; i < len(rowStrings)-1; i++ {
		htmlOutput += "<tr><td>" + rowStrings[i] + "</td></tr>"
	}
	htmlOutput += "</table>"
	return htmlOutput
}

func GenerateNewTableTERMS(uID int, setName string) string {
	data := dbOperations.GetTerms(setName, uID)
	htmlOutput := "<table id='terms'><tr><th>Term</th><th>Definition</th></tr>"
	for i := 0; i < len(data)-1; i++ {
		for n := 0; n < 1; n++ {
			htmlOutput += "<tr><td>" + data[i][n] + "</td></tr>"
		}
	}
	return htmlOutput
}
