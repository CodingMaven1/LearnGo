package main

import "fmt"

type Company struct{
	name string
	description string
}

type ParentCompany struct {
	Company
	companies int
}

func main() {
	google := Company{
		name: "Google",
		description: "The biggest search engine",
	}

	fmt.Println(google)

	alphabet := ParentCompany{
		Company: Company{
			name: "Alphabet",
			description: "We do everything",
		},
		companies: 25,
	}

	fmt.Println(alphabet)
}
