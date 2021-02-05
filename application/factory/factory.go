package factory

import (
	"github.com/eliashenriqueferreira/codepix/application/usecase"
	"github.com/eliashenriqueferreira/codepix/infrastructure/repository"
	"github.com/jinzhu/gorm"
	//"github.com/renatoosaka/codepix/application/usecase"
	//"github.com/renatoosaka/codepix/infrastructure/repository"
)

func TransactionUseCaseFactory(database *gorm.DB) usecase.TransactionUseCase {
	pixRepository := repository.PixKeyRepositoryDB{DB: database}
	transactionRepository := repository.TransactionRepositoryDB{DB: database}

	transactionUseCase := usecase.TransactionUseCase{
		TransactionRepository: &transactionRepository,
		PixRepository:         pixRepository,
	}

	return transactionUseCase
}
