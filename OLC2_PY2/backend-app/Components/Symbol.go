package Components

type Simbolo struct {
	Value     interface{}
	Type      string
	TypeValue string
	EsVector  bool
}

func NewSimbolo(valor interface{}, tipo string, tipo_variable string, vec bool) Simbolo {
	return Simbolo{
		Value:     valor,
		Type:      tipo,
		TypeValue: tipo_variable,
		EsVector:  vec,
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
