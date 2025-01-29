package data

import "github.com/brianvoe/gofakeit/v6"

type Instrument struct {
	Name string `fake{musical instrument}`
	Responsavel string `fake{lorenipsumsentence:10}`
	Situacao string `fake{lorenipsumsentence:10}`

}

func generateMusicalInstrument() []string {
	var i Instrument
	gofakeit.Struct(&i)

	iroot := []string{}
	iroot = append(iroot, i.Name)
	iroot = append(iroot, i.Responsavel)
	iroot = append(iroot, i.Situacao)
	return iroot
}

func MusicalInstrumentList(length int) [][]string {
	var instruments [][]string

	for i := 0; i < length; i++ {
		oneInstrument := generateMusicalInstrument()
		instruments = append(instruments, oneInstrument)
	}

	return instruments
}