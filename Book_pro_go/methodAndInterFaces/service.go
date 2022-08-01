package main

type Service struct {
	description string
	durationMonth int
	monthlyFee float64
}

func (s Service) getName() string {
	return s.description
}

func (s Service) getCost(recur bool) float64 {
	if (recur){
		return s.monthlyFee * float64(s.durationMonth)
	}
	return s.monthlyFee
}	