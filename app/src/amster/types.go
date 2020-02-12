package amster

//People is the structure datatype for the package
type People struct {
	Age               string `json:"age"`
	ParentsOrChildren string `json:"parentsOrChildrenAboard"`
	SiblingsOrSpouse  string `json:"siblingsOrSpousesAboard"`
	Fare              string `json:"fare"`
	Sex               string `json:"sex"`
	UUID              string `json:"uuid"`
	Survived          string `json:"survived"`
	PassengerClass    string `json:"passengerClass"`
	Name              string `json:"name"`
}
