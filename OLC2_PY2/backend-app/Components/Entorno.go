package Components

// Definición de la estructura Entorno
type Entorno struct {
	Padre         *Entorno
	TablaSimbolo  map[string]Simbolo
	TableFunction map[string]Function
	Ambito        string
	Retorno       bool
}

func NewEntorno(parent *Entorno, ambito string) *Entorno {
	return &Entorno{
		Padre:         parent,
		TablaSimbolo:  make(map[string]Simbolo),
		TableFunction: make(map[string]Function),
		Ambito:        ambito,
	}
}

// PARA SIMBOLOS

func (e *Entorno) EnvAddSimbolo(id string, value Simbolo) bool {
	_, ok := e.TablaSimbolo[id]
	if ok {
		return false
	} else {
		e.TablaSimbolo[id] = value
	}
	return true
}

func (e *Entorno) BuscarSimbolo(id string) Simbolo {
	for entorno := e; entorno != nil; entorno = entorno.Padre {
		value, ok := entorno.TablaSimbolo[id]
		if ok {
			return value
		}
	}
	return Simbolo{}
}

func (e *Entorno) ActualizarSimbolo(id string, value Simbolo) {
	for entorno := e; entorno != nil; entorno = entorno.Padre {
		_, ok := entorno.TablaSimbolo[id]
		if ok {
			entorno.TablaSimbolo[id] = value
			return
		}
	}
}

func SymbolVacio(s Simbolo) bool {
	if s.Value == nil && s.Type == "" && s.TypeValue == "" {
		return false
	}
	return true
}

func (e *Entorno) GetEntorno() string {
	return e.Ambito
}

// PARA FUNCIONES

func (e *Entorno) AddFuncion(id string, value Function) bool {
	_, ok := e.TableFunction[id]
	if !ok {
		e.TableFunction[id] = value
		return true
	}
	return false
}

func (e *Entorno) BuscarFuncion(id string) Function {
	for entorno := e; entorno != nil; entorno = entorno.Padre {
		value, ok := entorno.TableFunction[id]
		if ok {
			return value
		}
	}
	return Function{}
}

func FuncionVacio(s Function) bool {
	if s.ID == "" && s.Tipo == "" {
		return false
	}
	return true
}
