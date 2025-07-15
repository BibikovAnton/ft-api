package scors

import "github.com/BibikovAnton/finance-tracker-api/pkg/db"

type AccountsRepository struct {
	database *db.Db
}

func NewAccountsRepository(database *db.Db) *AccountsRepository {
	return &AccountsRepository{
		database: database,
	}
}

func (repo *AccountsRepository) Create(account *Account) (*Account, error) {
	result := repo.database.DB.Create(account)
	if result.Error != nil {
		return nil, result.Error
	}
	return account, nil
}

func (repo *AccountsRepository) Update(account *Account) (*Account, error) {
	result := repo.database.DB.Updates(account)
	if result.Error != nil {
		return nil, result.Error
	}

	return account, nil
}

func (repo *AccountsRepository) Get(id string) (*Account, error) {
	var account Account
	result := repo.database.DB.First(&account, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &account, nil
}

func (repo *AccountsRepository) GetAll(accounts *[]Account) (*[]Account, error) {

	result := repo.database.DB.Find(accounts)
	if result.Error != nil {
		return nil, result.Error
	}
	return accounts, nil
}

func (repo *AccountsRepository) Delete(id uint) error {
	result := repo.database.DB.Delete(&Account{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil

}
