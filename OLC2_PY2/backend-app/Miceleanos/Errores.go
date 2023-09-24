package Miceleanos

type ErrorAnalizador struct {
	Fila        int
	Columna     int
	Intruccion  string
	Descripcion string
	Tipo        string
}

func (a *ErrorAnalizador) NewError(file int, Col int, Ins string, des string, tipo string) {
	a.Fila = file
	a.Columna = Col
	a.Intruccion = Ins
	a.Descripcion = des
	a.Tipo = tipo
}
