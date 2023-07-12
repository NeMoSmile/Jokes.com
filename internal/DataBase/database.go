package database

type PData struct {
	FirstPl  string
	SecondPl string
	ThirdPl  string
	MyTitle  string
	MyText1  string
	MyText2  string
}

func Check(email, pass string) int {
	// data match: 1
	// data does not match: 2
	// No data: 3
	return 1
}

func Append(email, pass, name string) {

}

func PageData(email string) PData {
	firstPl := "Hubabuba"
	secondPl := "p1hdu)jd"
	thirdPl := "kapemo77"
	firstPoints := "107"
	secondPoints := "71"
	thirdPoints := "44"
	name := "NeMoSmile"
	todayPoints := "12"
	todayPlace := "49"
	monthPoints := "46"
	monthPlace := "32"
	dat := PData{
		FirstPl:  firstPl + ": " + firstPoints,
		SecondPl: secondPl + ": " + secondPoints,
		ThirdPl:  thirdPl + ": " + thirdPoints,
		MyTitle:  "You: " + name,
		MyText1:  "Today: " + todayPoints + " #" + todayPlace,
		MyText2:  "This Month: " + monthPoints + " #" + monthPlace,
	}
	return dat
}
