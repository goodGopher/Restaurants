package service

import (
	"log"

	"github.com/goodGopher/Restaurants"
	"github.com/goodGopher/Restaurants/pkg/repository"
)

type BookingService struct {
	bk repository.Booking
}

func NewBookingService(bk repository.Booking) *BookingService {
	return &BookingService{bk: bk}
}

func (b *BookingService) CreateBookingService(bookingReq Restaurants.BookingRequest) ([]int, error) {

	var BkList Restaurants.BookingList
	var resultIds []int
	TablesForBookingIds, err := b.bk.CheckTablesIdsRepo(bookingReq.RestName, bookingReq.PeopNum)
	if err != nil {
		log.Fatalf("error CheckTablesIdsRepo: %s", err.Error())
		return nil, err
	}

	BkList.UserId, err = b.bk.CreateUserForBookingRepo(bookingReq.UserName, bookingReq.UserPhone)
	if err != nil {
		log.Fatalf("error CreateUserForBookingRepo: %s", err.Error())
		return nil, err
	}

	for TableId, TableNum := range TablesForBookingIds {
		BkList.TableId = TableId
		BkList.TableNum = TableNum
		tId, err := b.bk.CreateBookingRepo(BkList)
		if err != nil {
			log.Fatalf("error CreateUserForBookingRepo: %s", err.Error())
			return nil, err
		}
		resultIds = append(resultIds, tId)
	}
	return resultIds, nil
}
