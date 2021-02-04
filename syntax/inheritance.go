package syntax

import (
	"fmt"
)

type Interface interface {
	Spark()
}

type Animal struct {
	eat   string
	spark string
}

type Dog struct {
	Animal
	legs int
}

type Cat struct {
	Animal
	legs int
}

type Labrador struct {
	Dog
}
func (dog Labrador) Spark() {
	fmt.Println("Labrador can spark:wang")
	spark(dog)
}

func (a Animal) Spark() {
	fmt.Println("Any animal can spark")
}

// 重写method 覆盖掉Animal

func (dog Dog) Spark() {
	fmt.Println("dog can spark:wangwangwang")
}
func (dog *Dog) spark() {
	fmt.Printf("==> %#v\n", dog)
	spark(dog)
}

func (cat Cat) Spark() {
	fmt.Println("cat can spark:miaomiaomiao")
}

func RunInheritance()  {
	animal := Animal{"food", "wowo"}
	animal.Spark()
	dog := Dog{Animal{"bone", "wangwangwang"}, 4}
	dog.Spark()
	cat := Cat{Animal{"fish", "miaomiaomiao"}, 4}
	cat.Spark()

	Labrador := Labrador{}
	Labrador.spark() // 实际调用的是Labrador.Dog.spark()
	Labrador.Dog.spark()
}

func spark(h Interface){
	h.Spark()
}