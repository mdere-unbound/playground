package departureModel

type Input struct {
	Ref DepartureRef
}

func DoSomeDBUpdate(input Input) DepartureRef {
	return input.Ref
}
