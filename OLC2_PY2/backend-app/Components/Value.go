package Components

type Value struct {
	Temp   interface{}
	Valor  interface{}
	Label1 interface{}
	Label2 interface{}
}

func NewValue(temp interface{}, label1 interface{}, label2 interface{}, valor interface{}) Value {
	return Value{
		Temp:   temp,
		Label1: label1,
		Label2: label2,
		Valor:  valor,
	}
}
