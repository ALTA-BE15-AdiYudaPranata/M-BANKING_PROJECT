package entities

type TopUp struct {
	Id      int
	IdUser  int
	IdOther int
	Value   int
}

type Transfer struct {
	Id     int
	IdUser int
	Value  int
}
