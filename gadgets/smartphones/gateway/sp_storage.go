package gateway

import (
	"github.com/enajera/phone-review/gadgets/smartphones/models"
	"github.com/enajera/phone-review/internal/database"
	"github.com/enajera/phone-review/internal/logs"
)

type SmartphoneGateway interface {
	Add(cmd *models.CreateSmartphoneCMD) (*models.Smartphone, error)
	Delete()
	FindById()
	

} 

type SmartphoneStorage struct {
	*database.MySqlClient
}

func (s *SmartphoneStorage) Add(cmd *models.CreateSmartphoneCMD) (*models.Smartphone,error){
   tx, err := s.MySqlClient.Begin()

   if err != nil{
	logs.Log().Error("Error create transaction")
   }

   res, err := tx.Exec(`insert into smartphone(name,price,country_origin,os)
   values(?,?,?,?)`, cmd.Name, cmd.Price, cmd.CountryOrigin, cmd.Os)
   if err != nil{
	  logs.Log().Error("Cannot execute statement")
	  _ = tx.Rollback()
	  return nil,err
   }

   id, err := res.LastInsertId()
   if err != nil{
	logs.Log().Error("Cannot fetch last id")
     _ = tx.Rollback()
   }

   _ = tx.Commit()

   return &models.Smartphone{
	Id: id,
	Name: cmd.Name,
	Price: cmd.Price,
	CountryOrigin: cmd.CountryOrigin,
	OS: cmd.Os,
   }, nil

}