package yy

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"hiyuncms/models/yy"
)

type Hospital struct{
	HospitalName string `json:"hospitalName"`
	HospitalId   int64  `json:"hospitalId"`
}

func GetHospital(c * gin.Context)  {
	hospitals := make( []* Hospital, 0)
	companies := yy.GetHospital()
	for _,company := range companies{
		hospitalTemp := Hospital{}
		hospitalTemp.HospitalId = company.Id
		hospitalTemp.HospitalName = company.CompanyName
		hospitals = append(hospitals, &hospitalTemp)
	}
	c.JSON(http.StatusOK, hospitals)
}