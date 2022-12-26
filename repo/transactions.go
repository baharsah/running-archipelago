package repo

import (
	"baharsah/models"

	"gorm.io/gorm"
)

type TransactionRepo interface {
	GetTransactions() ([]models.Transactions, error)
	GetTransaction(id int) (models.Transactions, error)
	// UpdateTransaction() (models.Transactions, error)
	SetTransaction(models.Transactions) (models.Transactions, error)
}

func RepoTRX(db *gorm.DB) *repo {

	return &repo{db}

}

func (r *repo) GetTransactions() ([]models.Transactions, error) {
	var trxs []models.Transactions

	err := r.db.Debug().Preload("Trips").Preload("Trips.Country").Find(&trxs).Error

	return trxs, err

}

func (r *repo) GetTransaction(id int) (models.Transactions, error) {
	tx := models.Transactions{}
	err2 := r.db.Preload("Trips").Preload("Trips.Country").First(&tx, id).Error
	return tx, err2
}

func (r *repo) SetTransaction(trx models.Transactions) (models.Transactions, error) {

	err := r.db.Debug().Create(&trx).Error

	return trx, err

}
