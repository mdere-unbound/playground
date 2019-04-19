package main

import (
	"fmt"

	"github.com/mdere-unbound/playground/handler"
	proto "github.com/mdere-unbound/playground/proto/departure"
)

func main() {
	ref := proto.DepartureRef{
		DepartureUUID:    "123",
		Title:            "Mike's Travel",
		PublishedAt:      "04-09-2019",
		CreatedAt:        "04-09-2019",
		DepartureStateID: 123,
		ModifiedAt:       "04-09-2019",
		Data:             "Some data",
		Currency:         "USD",
		UserUUID:         "123",
		DisplayImgSrc:    "image-src",
		DisplayImgUrl:    "image-url",
	}
	fmt.Println(handler.GetDeparture(ref))
}
