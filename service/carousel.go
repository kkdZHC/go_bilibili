package service

import (
	"go_bilibili/e"
	"go_bilibili/model"
	"go_bilibili/serializer"
)

type Carousel struct {
	Img string `json:"img"`
	Url string `json:"url"`
}

func (service *Carousel) Carousel() serializer.Response {
	code := e.SUCCESS
	var carousels []model.Carousel
	model.DB.Model(&model.Carousel{}).Find(&carousels)
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildCarousels(carousels),
	}
}
