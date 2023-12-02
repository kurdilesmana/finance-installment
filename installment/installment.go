package installment

import (
	"fmt"
	"math"
)

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func RunSimulation(params SimulationParams) (results InstallmentResults, err error) {
	Duration = params.Duration
	Base = params.LoanAmount
	Rate = params.EffectiveRate

	switch params.Pricing {
	case Pricing.Annuity: // effective
		err = Annuity()
	default:
		remark := fmt.Sprintf("Pricing %s under development.", PricingDesc[params.Pricing])
		err = fmt.Errorf(remark)
	}

	results.Installment = Installment
	results.Margin = Margin
	results.Items = Items

	return
}

func Annuity() (err error) {
	prate := Rate / 100
	mvalue := Base * prate / 12
	mrate := prate / 12
	multipler := math.Pow(1+mrate, float64(Duration))

	// set installment & total margin
	Installment = roundFloat(mvalue/(1-(1/multipler)), 2)
	Margin = roundFloat(Installment*float64(Duration)-Base, 2)

	// get detail installment items
	effectiveOutstanding := Base
	for i := 1; i < int(Duration); i++ {
		item := InstallmentItems{}
		item.SequenceNumber = int32(i)
		item.NominalInstallment = Installment
		item.ProfitAmount = roundFloat(effectiveOutstanding*mrate, 2)
		item.PrincipalAmount = roundFloat(Installment-item.ProfitAmount, 2)
		item.Oustanding = roundFloat(effectiveOutstanding, 2)
		item.PercentageRate = prate
		Items = append(Items, item)

		effectiveOutstanding -= item.PrincipalAmount
	}

	// balancing last detail item
	item := InstallmentItems{}
	item.SequenceNumber = int32(Duration)
	item.NominalInstallment = Installment
	item.PrincipalAmount = roundFloat(effectiveOutstanding, 2)
	item.ProfitAmount = roundFloat(Installment-effectiveOutstanding, 2)
	item.Oustanding = roundFloat(effectiveOutstanding, 2)
	item.PercentageRate = prate
	Items = append(Items, item)

	return
}
