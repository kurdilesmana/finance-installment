package installment

var (
	Duration int64
	Base     float64
	Rate     float64
)

var Pricing = struct {
	Annuity  string
	Sliding  string
	Rest     string
	Progress string
	Flat     string
}{
	Annuity:  "1",
	Sliding:  "2",
	Rest:     "4",
	Progress: "5",
	Flat:     "6",
}

var PricingDesc = map[string]string{
	"1": "Annuity",
	"2": "Sliding",
	"4": "Rest",
	"5": "Progress",
	"6": "Flat",
}
