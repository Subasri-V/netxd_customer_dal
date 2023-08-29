package interfaces

import "Application-New/netxd_customer_dal/models"

type ICustomer interface{
	CreateCustomer(*models.CustomerDetails)(string)
}