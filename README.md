# Finance Installment (Golang Module)
Calculator installment for loan simulation

## Features 
- Annuity Pricing
Calculation loan with pricing model annuity.



## Installation

Install module with go get

```bash
  go get "github.com/kurdilesmana/finance-installment"
```

## Usage/Examples

```golang
package main

import (
	"fmt"

	"github.com/kurdilesmana/finance-installment/installment"
)

func main() {
	Parameter := installment.SimulationParams{
		Pricing:       installment.Pricing.Flat,
		LoanAmount:    10000000,
		Duration:      12,
		EffectiveRate: 13,
	}
	items, err := installment.RunSimulation(Parameter)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	fmt.Printf("Items: %v", items)
}

```

