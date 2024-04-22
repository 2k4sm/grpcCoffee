package services

import (
	"github.com/2k4sm/httpCoffee/src/main/dto"
	"github.com/2k4sm/httpCoffee/src/main/entities"
	"github.com/2k4sm/httpCoffee/src/main/repositories"
)

type PayServiceInterface interface {
	GetAll() []entities.Payment
	GetById(id int) entities.Payment
	GetAllByUserId(uid int) []entities.Payment
	GetAllByHouseId(hid int) []entities.Payment
	CreatePayment(payment dto.CreatePayment) entities.Payment
	DeletePayment(id int) (entities.Payment, error)
}

type payService struct {
	PayRepository repositories.PayRepositoryInterface
}

func NewPayService(payRepository repositories.PayRepositoryInterface) PayServiceInterface {
	return &payService{
		PayRepository: payRepository,
	}
}

func (ps *payService) GetAll() []entities.Payment {
	return ps.PayRepository.FindAll()
}

func (ps *payService) GetById(id int) entities.Payment {
	return ps.PayRepository.FindById(uint(id))
}

func (ps *payService) GetAllByUserId(uid int) []entities.Payment {
	return ps.PayRepository.FindAllByUserId(uint(uid))
}

func (ps *payService) GetAllByHouseId(hid int) []entities.Payment {
	return ps.PayRepository.FindAllByHouseId(uint(hid))
}

func (ps *payService) CreatePayment(payment dto.CreatePayment) entities.Payment {
	newpayment := dto.ParseToPaymentEntity(payment)

	return ps.PayRepository.SavePayment(newpayment)
}

func (ps *payService) DeletePayment(id int) (entities.Payment, error) {
	return ps.PayRepository.DeletePayment(uint(id))
}
