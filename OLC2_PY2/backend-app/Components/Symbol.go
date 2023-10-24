package Components

type Simbolo struct {
	Value     interface{}
	Type      string
	TypeValue string
	EsVector  bool
	Stack     int
}

func NewSimbolo(valor interface{}, tipo string, tipo_variable string, vec bool) Simbolo {
	return Simbolo{
		Value:     valor,
		Type:      tipo,
		TypeValue: tipo_variable,
		EsVector:  vec,
		Stack:     0,
	}
}

type GlobalSymbol struct {
	TypeSymbol string
	Type       string
	Env        string
	Line       int
	Col        int
}

func NewSimbGlobal(t string, ty string, env string, l int, c int) GlobalSymbol {
	return GlobalSymbol{
		TypeSymbol: t,
		Type:       ty,
		Env:        env,
		Line:       l,
		Col:        c,
	}
}

func (e *Simbolo) SetStack(d int) {
	e.Stack = d
}

func (e *Simbolo) GetStack() int {
	return e.Stack
}
