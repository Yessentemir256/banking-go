package stats

import (
	"fmt"
	"github.com/Yessentemir256/banking-go/pkg/banking-go/types"
)

func ExampleAvg() {
	fmt.Println(Avg([]types.Payment{
		{
			Amount: 1_000_00,
		},
	}))
	fmt.Println(Avg([]types.Payment{
		{
			Amount: 1_000_00,
		},
		{
			Amount: 2_000_00,
		},
	}))
	fmt.Println(Avg([]types.Payment{
		{
			Amount: 1_000_00,
		},
		{
			Amount: 2_000_00,
		},
		{
			Amount: 3_000_00,
		},
	}))

	// Output:
	// 100000
	// 150000
	// 200000
}

func ExampleTotal() {
	fmt.Println(Total([]types.Payment{
		{
			Amount: 1_000_00,
		},
	}))
	fmt.Println(Total([]types.Payment{
		{
			Amount: 1_000_00,
		},
		{
			Amount: 2_000_00,
		},
	}))
	fmt.Println(Total([]types.Payment{
		{
			Amount: 1_000_00,
		},
		{
			Amount: -500_00,
		},
	}))

	// Output:
	// 100000
	// 300000
	// 100000
}
