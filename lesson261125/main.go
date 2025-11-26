package main

type MyType struct {
	Name string
}

func (MyType) Val() {

}

func (*MyType) Ptr() {

}

type V interface {
	Val()
}

type PV interface {
	Ptr()
}

func main() {
	var _ V = &MyType{}
	var _ V = MyType{}
	var _ PV = &MyType{}
	// var _ PV = MyType{} // Ошибка компиляции
	/*
		Интерфейс V может быть реализован как указателем на MyType, так и самим MyType,
		поскольку метод Val() имеет значение получателя.
		Интерфейс PV может быть реализован только указателем на MyType,
		поскольку метод Ptr() имеет указатель в качестве получателя.
		_ {
			type PV
			value ptr *MyType {
				ptr *MyType
				method Ptr ptr
				value MyType
				method Val MyType
			}
		}
	*/
}
