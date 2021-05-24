package services

import "bitPay/utills/responses"

var (
	StragetiesService	strategieServiceInterface = &strategiesService{}
)

type strategiesService struct {}

type strategieServiceInterface interface {

}

func (s *strategiesService) Martingale() *responses.Response {
	return nil
}