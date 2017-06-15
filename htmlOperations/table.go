package htmlOperations

func GenerateNewTable(rowStrings [][]string) string {
	htmlOutput := "<table>"
	for i := 0; i < len(rowStrings)-1; i++ {
		htmlOutput += newRow(rowStrings[i])
	}
	htmlOutput += "</table>"
	return htmlOutput
}

func newRow(a []string) string {
	var k string
	k += "<tr>"
	for i := 0; i < len(a)-1; i++ {
		k += "<th>" + a[i] + "</th>"
	}
	k += "</tr>"
	return k
}
