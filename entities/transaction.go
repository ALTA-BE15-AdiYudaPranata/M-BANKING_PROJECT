package entities

type Transfer struct {
	Id      int
	IdUser  int
	IdOther int
	Value   int
	Created string
}

type TopUp struct {
	Id      int
	IdUser  int
	Value   int
	Created string
}
