package handler

import (
	departureModel "github.com/mdere-unbound/playground/model"
	proto "github.com/mdere-unbound/playground/proto/departure"
)

func GetDeparture(req proto.DepartureRef) proto.DepartureRef {
	input := departureModel.Input{
		Ref: departureModel.DepartureRef(req),
	}
	newRef := departureModel.DoSomeDBUpdate(input)
	newProtoRef := proto.DepartureRef(newRef)
	return newProtoRef
}
