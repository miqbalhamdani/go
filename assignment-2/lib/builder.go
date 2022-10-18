package lib

import (
	"golang-web-service/assignment-2/models"
	"golang-web-service/assignment-2/transports"
)

func BuildResponseData(orderData models.Orders, itemData []models.Items) transports.Response {
	var itemRespData transports.ItemResponse
	responseData := transports.Response{
		OrderID:      orderData.OrderID,
		CustomerName: orderData.CustomerName,
	}

	for _, v := range itemData {

		itemRespData.ItemID = v.ItemID
		itemRespData.ItemCode = v.ItemCode
		itemRespData.Description = v.Description
		itemRespData.Quantity = v.Quantity

		responseData.CustomerItems = append(responseData.CustomerItems, itemRespData)
	}

	return responseData
}
