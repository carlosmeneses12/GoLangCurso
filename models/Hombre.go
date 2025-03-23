package models

type Hombre struct {
	edad       int
	altura     float32
	peso       float32
	respirando bool
	pensando   bool
	comimendo  bool
	vivo       bool
}

func (h *Hombre) Respirar()    { h.respirando = true }
func (h *Hombre) Comer()       { h.comimendo = true }
func (h *Hombre) Pensar()      { h.pensando = true }
func (h *Hombre) Estavivo()    { h.vivo = true }
func (h *Hombre) Sexo() string { return "hombre" }
