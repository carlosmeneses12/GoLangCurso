package models

type Mujer struct {
	Hombre
	Esmadre bool
}

func (m *Mujer) Respirar()    { m.respirando = true }
func (m *Mujer) Comer()       { m.comimendo = true }
func (m *Mujer) Pensar()      { m.pensando = true }
func (m *Mujer) Sexo() string { return "Mujer" }
