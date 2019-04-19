package proto

type DepartureRef struct {
	DepartureUUID    string `protobuf:"bytes,2,opt,name=departureUUID,proto3" json:"departureUUID,omitempty" sql:"DepartureUUID" pk:"DepartureUUID"`
	Title            string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty" sql:"Title"`
	PublishedAt      string `protobuf:"bytes,7,opt,name=publishedAt,proto3" json:"publishedAt,omitempty" sql:"PublishedAt"`
	CreatedAt        string `protobuf:"bytes,8,opt,name=createdAt,proto3" json:"createdAt,omitempty" sql:"created_at"`
	DepartureStateID int32  `protobuf:"varint,9,opt,name=departureStateID,proto3" json:"departureStateID,omitempty" sql:"-"`
	ModifiedAt       string `protobuf:"bytes,11,opt,name=modifiedAt,proto3" json:"modifiedAt,omitempty" sql:"-"`
	Data             string `protobuf:"bytes,12,opt,name=data,proto3" json:"data,omitempty" sql:"Data"`
	Currency         string `protobuf:"bytes,13,opt,name=currency,proto3" json:"currency,omitempty" sql:"Currency"`
	UserUUID         string `protobuf:"bytes,14,opt,name=userUUID,proto3" json:"userUUID,omitempty" sql:"UserUUID"`
	DisplayImgSrc    string `protobuf:"bytes,15,opt,name=displayImgSrc,proto3" json:"displayImgSrc,omitempty" sql:"display_img_src"`
	DisplayImgUrl    string `protobuf:"bytes,27,opt,name=displayImgUrl,proto3" json:"displayImgUrl,omitempty"`
}
