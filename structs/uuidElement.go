package structs

import "kuzmin.com/structs2/utils"

type UUUIDElement struct {
	Id        string `json:"id"`
	CreatedAt string `json:"createdAt"`
}

func GenerateUUIDElement() UUUIDElement {
	return UUUIDElement{utils.GenerateUUID(), utils.GetCreatedAt()}
}

func (elemnt UUUIDElement) GetUUID() string {
	return elemnt.Id
}
