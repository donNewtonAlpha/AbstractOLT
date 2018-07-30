package models

type Ont struct {
	Number int
	Svlan  int
	Cvlan  int
	Parent *Port
}
