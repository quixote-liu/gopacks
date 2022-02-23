package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	u, err := uuid.Parse("cb49b381-9012-40cb-b8ee-80c19a4801b5")
	if err != nil {
		panic(err)
	}
	fmt.Println("u.String", u.String())

	u2, err := uuid.Parse("0c4e939acacf4376bdcd1129f1a054ad")
	if err != nil {
		panic(err)
	}
	fmt.Println("u2.String", u2.String())

	u3, err := uuid.Parse("12343242513")
	if err != nil {
		panic(err)
	}
	fmt.Println("u3.String", u3.String())
}
