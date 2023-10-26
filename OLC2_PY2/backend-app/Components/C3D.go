package Components

import (
	"fmt"
	"strings"
)

type Traduction struct {
	Temp           int      // indice de tempprales
	Label          int      // indice de Etiquetas
	List_temp      []string // Lista de Temporales
	List_Process   []string // Lista de proceso dentro de la funcion main()
	IndexStack     int      // LLeva un indice de datos guardados en Stack
	Code_Final     []string // Genera la Traduccion final
	List_funciones []string // Esta es una Lista que traera todas las funciones
	IsMain         bool     // Para saber si es main o no
	IsTypeVec      int
}

func NewTraduction() *Traduction {
	return &Traduction{
		Temp:           0,
		Label:          0,
		List_temp:      make([]string, 0),
		IndexStack:     0,
		Code_Final:     make([]string, 0),
		List_funciones: []string{},
		IsMain:         true,
	}
}

func (t *Traduction) NewTemp() string {
	temporal := "t" + fmt.Sprintf("%v", t.Temp)
	t.Temp++
	t.List_temp = append(t.List_temp, temporal)
	return temporal
}

func (t *Traduction) GetIndex() int {
	newpos := t.IndexStack
	t.IndexStack++
	return newpos
}

// asignacion
func (t *Traduction) Igual(NewTemp string, l interface{}, r interface{}, op interface{}) {
	NewCode := ""
	if t.IsMain {
		if op == nil {
			NewCode = "\t\t" + NewTemp + " = " + fmt.Sprint(l) + ";" + "\n"
		} else {
			NewCode = "\t\t" + NewTemp + " = " + fmt.Sprint(l) + " " + fmt.Sprint(op) + " " + fmt.Sprint(r) + ";" + "\n"
		}
		t.List_Process = append(t.List_Process, NewCode)
	} else {
		if op == nil {
			NewCode = "\t\t" + NewTemp + " = " + fmt.Sprint(l) + ";" + "\n"
		} else {
			NewCode = "\t\t" + NewTemp + " = " + fmt.Sprint(l) + " " + fmt.Sprint(op) + " " + fmt.Sprint(r) + ";" + "\n"
		}
		t.List_funciones = append(t.List_funciones, NewCode)
	}
}

//---------------------------------- Etiquetas y Goto ----------------------------------

func (t *Traduction) NewLabel() string {
	label := "L" + fmt.Sprintf("%v", t.Label)
	t.Label++
	return label
}

func (t *Traduction) AddGoto(label string) {
	if t.IsMain {
		NewCode := "\t\t" + "goto " + label + ";" + "\n"
		t.List_Process = append(t.List_Process, NewCode)

	} else {
		NewCode := "\t\t" + "goto " + label + ";" + "\n"
		t.List_funciones = append(t.List_funciones, NewCode)
	}
}

func (t *Traduction) AddLabel(label string) {
	if t.IsMain {
		NewCode := "\t" + label + ":" + "\n"
		t.List_Process = append(t.List_Process, NewCode)
	} else {
		NewCode := "\t" + label + ":" + "\n"
		t.List_funciones = append(t.List_funciones, NewCode)
	}
}

func (t *Traduction) AddIf(izq string, der string, op string, label string) {
	if t.IsMain {
		NewCode := "\t\t" + "if (" + izq + " " + op + " " + der + ")  goto " + label + ";" + "\n"
		t.List_Process = append(t.List_Process, NewCode)
	} else {
		NewCode := "\t\t" + "if (" + izq + " " + op + " " + der + ")  goto " + label + ";" + "\n"
		t.List_funciones = append(t.List_funciones, NewCode)
	}
}

// Imprimir
func (t *Traduction) AddPrint(value string, tipo string) {
	if t.IsMain {
		NewCode := "\t\t" + "printf(\"%" + tipo + "\" , " + value + ");" + "\n"
		t.List_Process = append(t.List_Process, NewCode)
	} else {
		NewCode := "\t\t" + "printf(\"%" + tipo + "\" , " + value + ");" + "\n"
		t.List_funciones = append(t.List_funciones, NewCode)
	}
}

// Comentario
func (t *Traduction) AddComentario(comentario string) {
	if t.IsMain {
		NewCode := "\t" + comentario + "\n"
		t.List_Process = append(t.List_Process, NewCode)
	} else {
		NewCode := "\t" + comentario + "\n"
		t.List_funciones = append(t.List_funciones, NewCode)
	}
}

// heap
func (t *Traduction) SetHeap(Pos string, Value string) {
	if t.IsMain {
		NewCode := "\t\t" + "heap[" + Pos + "] = " + Value + " ;" + "\n"
		t.List_Process = append(t.List_Process, NewCode)
	} else {
		NewCode := "\t\t" + "heap[" + Pos + "] = " + Value + " ;" + "\n"
		t.List_funciones = append(t.List_funciones, NewCode)
	}
}

func (t *Traduction) GetHeap(temp string, index string) {
	if t.IsMain {
		NewCode := "\t\t" + temp + " = " + "heap[" + index + "]" + ";" + "\n"
		t.List_Process = append(t.List_Process, NewCode)
	} else {
		NewCode := "\t\t" + temp + " = " + "heap[" + index + "]" + ";" + "\n"
		t.List_funciones = append(t.List_funciones, NewCode)
	}
}

// Stack
func (t *Traduction) SetStack(index string, temp string) {
	if t.IsMain {
		NewCode := "\t\t" + "stack[" + index + "] = " + temp + ";" + "\n"
		t.List_Process = append(t.List_Process, NewCode)
	} else {
		NewCode := "\t\t" + "stack[" + index + "] = " + temp + ";" + "\n"
		t.List_funciones = append(t.List_funciones, NewCode)
	}
}
func (t *Traduction) GetStack(temp string, index string) {
	if t.IsMain {
		NewCode := "\t\t" + temp + " = " + "stack[" + index + "]" + ";" + "\n"
		t.List_Process = append(t.List_Process, NewCode)
	} else {
		NewCode := "\t\t" + temp + " = " + "stack[" + index + "]" + ";" + "\n"
		t.List_funciones = append(t.List_funciones, NewCode)
	}
}

// Contador de Pointer Heap y Pointer Stack
func (t *Traduction) Contador(Actual string, Anterior string, op string, cuanto string) {
	if t.IsMain {
		NewCode := "\t\t" + Actual + " = " + Anterior + " " + op + " " + cuanto + ";" + "\n"
		t.List_Process = append(t.List_Process, NewCode)
	} else {
		NewCode := "\t\t" + Actual + " = " + Anterior + " " + op + " " + cuanto + ";" + "\n"
		t.List_funciones = append(t.List_funciones, NewCode)
	}
}

// -------------------------------- funciones --------------------------------

func (t *Traduction) callVoid(id string) {
	if t.IsMain {
		NewCode := "\t\t" + id + "();" + "\n"
		t.List_Process = append(t.List_Process, NewCode)
	} else {
		NewCode := "\t\t" + id + "();" + "\n"
		t.List_funciones = append(t.List_funciones, NewCode)
	}
}
func (t *Traduction) Return(tipo interface{}) {
	if t.IsMain {
		if tipo == nil {
			t.List_funciones = append(t.List_funciones, "\treturn; \n")
		} else {
			t.List_funciones = append(t.List_funciones, "\treturn "+fmt.Sprint(tipo)+"; \n")
		}
	} else {
		if tipo == nil {
			t.List_funciones = append(t.List_funciones, "\treturn; \n")
		} else {
			t.List_funciones = append(t.List_funciones, "\treturn "+fmt.Sprint(tipo)+"; \n")
		}
	}
}

// ------------------------------------ Funciones Nativas ------------------------------------
func (t *Traduction) PrimitiveString() string {
	NewTemp := t.NewTemp()
	NewTemp2 := t.NewTemp()
	NewTemp3 := t.NewTemp()
	Label := t.NewLabel()
	Label2 := t.NewLabel()

	PrintString := "\n"
	PrintString += "// Imprimir String Primitivo" + "\n"
	PrintString += "void ImprimirString() {" + "\n"
	PrintString += "\t" + NewTemp + " = " + "P + 1" + ";" + "\n"
	PrintString += "\t" + NewTemp2 + " = " + "stack[ (int) " + NewTemp + "]" + ";" + "\n"
	PrintString += "\t" + Label + ":" + "\n"
	PrintString += "\t" + NewTemp3 + " = " + "heap[(int) " + NewTemp2 + "]" + ";" + "\n"
	PrintString += "\t" + "if (" + NewTemp3 + " == -1) goto " + Label2 + ";" + "\n"
	PrintString += "\t" + "printf(\"%c\" , (char) " + NewTemp3 + ");" + "\n"
	PrintString += "\t" + NewTemp2 + " = " + NewTemp2 + " + 1" + ";" + "\n"
	PrintString += "\t" + "goto " + Label + ";" + "\n"
	PrintString += "\t" + Label2 + ":" + "\n"
	PrintString += "\t" + "return;" + "\n"
	PrintString += "}" + "\n"
	PrintString +=
		"\n"

	return PrintString
}

func (t *Traduction) PrimiteConcatString() string {
	NewTemp := t.NewTemp()
	NewTemp2 := t.NewTemp()
	NewTemp3 := t.NewTemp()
	Label := t.NewLabel()
	Label2 := t.NewLabel()

	PrintString := "\n"
	PrintString += "// Concatenar String Primitivo" + "\n"
	PrintString += "void CopiarHeap() {" + "\n"
	PrintString += "\t" + NewTemp + " = " + "P + 1" + ";" + "\n"
	PrintString += "\t" + NewTemp2 + " = " + "stack[ (int) " + NewTemp + "]" + ";" + "\n"
	PrintString += "\t" + Label + ":" + "\n"
	PrintString += "\t" + NewTemp3 + " = " + "heap[(int) " + NewTemp2 + "]" + ";" + "\n"
	PrintString += "\t" + "if (" + NewTemp3 + " == -1) goto " + Label2 + ";" + "\n"
	PrintString += "\t" + "heap[ (int) H ] = " + NewTemp3 + ";" + "\n"
	PrintString += "\t" + "H = H + 1" + ";" + "\n"
	PrintString += "\t" + NewTemp2 + " = " + NewTemp2 + " + 1" + ";" + "\n"
	PrintString += "\t" + "goto " + Label + ";" + "\n"
	PrintString += "\t" + Label2 + ":" + "\n"
	PrintString += "\t" + "return;" + "\n"
	PrintString += "}" + "\n"
	PrintString += "\n"

	return PrintString
}

func (t *Traduction) VectorCount() string {

	NewTemp := t.NewTemp()
	NewTemp2 := t.NewTemp()
	NewTemp3 := t.NewTemp()
	NewTemp4 := t.NewTemp()

	NewLabelSI := t.NewLabel()
	NewLabelNO := t.NewLabel()

	Print_String := "\n"
	Print_String += "// Contar Vector" + "\n"
	Print_String += "void VectorCount() {" + "\n"
	Print_String += "\t" + NewTemp + " = " + "P + 1" + ";" + "\n"
	Print_String += "\t" + NewTemp2 + " = " + "stack[ (int) " + NewTemp + "]" + ";" + "\n"
	Print_String += "\t" + NewTemp3 + " = " + "0" + ";" + "\n"
	Print_String += "\t" + NewLabelNO + ":" + "\n"
	Print_String += "\t" + NewTemp4 + " = " + "heap[(int) " + NewTemp2 + "]" + ";" + "\n"
	Print_String += "\t" + "if (" + NewTemp4 + " == -1) goto " + NewLabelSI + ";" + "\n"
	Print_String += "\t" + NewTemp3 + " = " + NewTemp3 + " + 1" + ";" + "\n"
	Print_String += "\t" + NewTemp2 + " = " + NewTemp2 + " + 1" + ";" + "\n"
	Print_String += "\t" + "goto " + NewLabelNO + ";" + "\n"
	Print_String += "\t" + NewLabelSI + ":" + "\n"
	Print_String += "\t" + "stack[(int) P] = " + NewTemp3 + ";" + "\n"
	Print_String += "\t" + "return;" + "\n"
	Print_String += "}" + "\n"
	Print_String += "\n"

	return Print_String

}

// ------------------------------------------------ Codigo Final de la traduccion ----------------------------------------------

func (t *Traduction) GenerateTraduction() []string {
	Encabezado := "#include <stdio.h> \n\n"
	Encabezado += "float stack[100000];\n"
	Encabezado += "float heap[100000];\n"
	Encabezado += "float P;\n"
	Encabezado += "float H;\n"
	t.Code_Final = append(t.Code_Final, Encabezado)
	if len(t.List_temp) != 0 {
		Temp := "float " + strings.Join(t.List_temp, ",") + ";" + "\n"
		t.Code_Final = append(t.Code_Final, Temp)
	}
	t.Code_Final = append(t.Code_Final, t.List_funciones...)
	t.Code_Final = append(t.Code_Final, "int main () {  \n")
	t.Code_Final = append(t.Code_Final, t.List_Process...)
	t.Code_Final = append(t.Code_Final, "\treturn 0; \n")
	t.Code_Final = append(t.Code_Final, "} \n")
	return t.Code_Final
}

func (t *Traduction) Br() {
	t.List_Process = append(t.List_Process, "\n")
}
