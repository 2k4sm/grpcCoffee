package repositories

import (
	"github.com/2k4sm/httpCoffee/src/main/entities"
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/gorm"
)

type PayRepositoryInterface interface {
	FindAll() []entities.Payment
	FindById(id uint) entities.Payment
	FindAllByUserId(uid uint) []entities.Payment
	FindAllByHouseId(hid uint) []entities.Payment
	SavePayment(payment entities.Payment) entities.Payment
	DeletePayment(id uint) (entities.Payment, error)
}

type payRepository struct {
	Db *gorm.DB
}

func NewPayRepository(db *gorm.DB) PayRepositoryInterface {
	return &payRepository{
		Db: db,
	}
}

func (p *payRepository) FindAll() []entities.Payment {
	payments := make([]entities.Payment, 0)

	p.Db.Preload("Items").Find(&payments)

	if len(payments) == 0 {
		log.Info("No Payment Found")
	}

	return payments
}

func (p *payRepository) FindById(id uint) entities.Payment {
	payment := entities.Payment{}

	p.Db.Preload("Items").First(&payment, id)

	if payment.ID == 0 {
		log.Infof("No payment found with id: %s", id)
	}

	return payment

}

func (p *payRepository) FindAllByUserId(uid uint) []entities.Payment {
	payments := []entities.Payment{}

	p.Db.Preload("Items").Find(&payments, "UserID = ?", uid)

	if len(payments) == 0 {
		log.Infof("No Payment Found for user: %s", uid)
	}

	return payments
}

func (p *payRepository) FindAllByHouseId(hid uint) []entities.Payment {
	payments := []entities.Payment{}

	p.Db.Preload("Items").Find(&payments, "CoffeeHouseID = ?", hid)

	if len(payments) == 0 {
		log.Infof("No Payment Found for house: %s", hid)
	}

	return payments
}

func (p *payRepository) SavePayment(payment entities.Payment) entities.Payment {
	err := p.Db.Save(&payment)

	if err != nil {
		log.Infof("Error saving payment: %s", err.Error)
	}

	return payment
}

func (p *payRepository) DeletePayment(id uint) (entities.Payment, error) {
	paymentToDel := entities.Payment{}

	p.Db.Preload("Items").First(&paymentToDel, id)

	if paymentToDel.ID == 0 {
		return paymentToDel, gorm.ErrRecordNotFound
	}

	p.Db.Delete(&paymentToDel)

	return paymentToDel, nil
}
