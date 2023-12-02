package installment

type SimulationParams struct {
	Duration      int64
	Pricing       string
	LoanAmount    float64
	EffectiveRate float64
}

type InstallmentResults struct {
	Installment float64
	Margin      float64
	Items       []InstallmentItems
}

type InstallmentItems struct {
	SequenceNumber     int32
	Oustanding         float64
	PrincipalAmount    float64
	ProfitAmount       float64
	PercentageRate     float64
	NominalInstallment float64
}
