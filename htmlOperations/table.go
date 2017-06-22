package htmlOperations

import "github.com/willcruse/ComputingCoursework/dbOperations"

func GenerateNewTable(uID, operation int, setName string) string {
	rowStrings := dbOperations.GetSets(uID)
	htmlOutput := "<table><tr><th>Sets</th></tr>"

	for i := 0; i < len(rowStrings)-1; i++ {
		htmlOutput += "<tr><td>" + rowStrings[i] + "</td></tr>"
	}
	htmlOutput += "</table>"
	return htmlOutput
}
