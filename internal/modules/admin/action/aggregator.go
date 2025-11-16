package action

// Aggregator собирает все процессы модуля в одно целое
type Aggregator struct {
	//CreateDoctor *create_doctor.Action
}

func NewAggregator() *Aggregator {
	return &Aggregator{
		//CreateDoctor: create_doctor.New(),
	}
}
