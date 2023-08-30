package interfaces

import(
	models "github.com/Subasri-V/application-new/netxd_customer_dal/models"
)
//import "Application-New/netxd_customer_dal/models"

type ICustomer interface {
	CreateCustomer(*models.CustomerDetails) (*models.DBResponse,error)
}
