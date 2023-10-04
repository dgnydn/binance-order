package store

import (
	"emperror.dev/errors"
	"github.com/adshao/go-binance/v2"
	"github.com/dgnydn/binance-order/internal/app/order/domain"
	"gorm.io/gorm"
)

type Store interface {
	CreateOrder(order domain.Order) (domain.Order, error)
	GetOrder(id int64) (domain.Order, error)
	GetAllOrders() ([]domain.Order, error)
	DeleteOrder(id int64) error
	UpdateOrder(orderId int64, order domain.Order) error
	GetAllOrdersBySymbol(symbol string) ([]domain.Order, error)
	GetAllOrdersByStatus(status string) ([]domain.Order, error)
	GetAllOrdersBySymbolAndStatus(symbol, status string) ([]domain.Order, error)
	GetAllOrdersBySide(side binance.SideType) ([]domain.Order, error)
}

type store struct {
	db *gorm.DB
}

type Option func(*store)

func WithGorm(db *gorm.DB) Option {
	return func(s *store) {
		s.db = db
	}
}

func New(opts ...Option) Store {
	s := &store{}
	for _, o := range opts {
		o(s)
	}
	return s
}

func (s *store) GetOrder(id int64) (domain.Order, error) {
	if id == 0 {
		return domain.Order{}, errors.New("Order ID is empty")
	}

	var order domain.Order

	if err := s.db.Where("id = ?", id).First(&order).Error; err != nil {
		return domain.Order{}, err
	}

	return order, nil
}

func (s *store) GetAllOrders() ([]domain.Order, error) {
	var orders []domain.Order
	if err := s.db.Find(&orders).Error; err != nil {
		return nil, err
	}

	return orders, nil
}

func (s *store) DeleteOrder(id int64) error {
	if id == 0 {
		return errors.New("Order ID is empty")
	}

	if err := s.db.Where("id = ?", id).Delete(&domain.Order{}).Error; err != nil {
		return err
	}

	return nil
}

func (s *store) UpdateOrder(orderId int64, order domain.Order) error {
	if orderId == 0 {
		return errors.New("Order ID is empty")
	}

	if !order.IsValid() {
		return errors.New("Order is not valid")
	}

	if err := s.db.Model(&domain.Order{}).Where("id = ?", orderId).Updates(order).Error; err != nil {
		return err
	}

	return nil
}

func (s *store) GetAllOrdersBySymbol(symbol string) ([]domain.Order, error) {
	if symbol == "" {
		return nil, errors.New("Symbol is empty")
	}

	var orders []domain.Order

	if err := s.db.Where("symbol = ?", symbol).Find(&orders).Error; err != nil {
		return nil, err
	}

	return orders, nil
}

func (s *store) GetAllOrdersByStatus(status string) ([]domain.Order, error) {
	if status == "" {
		return nil, errors.New("Status is empty")
	}

	var orders []domain.Order

	if err := s.db.Where("status = ?", status).Find(&orders).Error; err != nil {
		return nil, err
	}

	return orders, nil
}

func (s *store) GetAllOrdersBySymbolAndStatus(symbol, status string) ([]domain.Order, error) {
	if symbol == "" {
		return nil, errors.New("Symbol is empty")
	}
	if status == "" {
		return nil, errors.New("Status is empty")
	}

	var orders []domain.Order

	if err := s.db.Where("symbol = ? AND status = ?", symbol, status).Find(&orders).Error; err != nil {
		return nil, err
	}

	return orders, nil
}

func (s *store) GetAllOrdersBySide(side binance.SideType) ([]domain.Order, error) {
	if side == "" {
		return nil, errors.New("Side is empty")
	}

	var orders []domain.Order
	if err := s.db.Where("side = ?", side).Find(&orders).Error; err != nil {
		return nil, err
	}

	return orders, nil
}

func (s *store) CreateOrder(order domain.Order) (domain.Order, error) {
	if !order.IsValid() {
		return domain.Order{}, errors.New("Order is not valid")
	}

	if err := s.db.Create(&order).Error; err != nil {
		return domain.Order{}, err
	}

	return order, nil
}
