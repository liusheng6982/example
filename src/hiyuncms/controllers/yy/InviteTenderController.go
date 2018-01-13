package yy

import (
	"github.com/gin-gonic/gin"
	"hiyuncms/controllers"
	"net/http"
	"hiyuncms/models"
	"hiyuncms/models/yy"
	"hiyuncms/controllers/frontend"
	"strconv"
	"log"
	"fmt"
	"net/url"
	"encoding/json"
	"io/ioutil"
)

func InviteTenderListShow(c *gin.Context){
	c.HTML(http.StatusOK, "invitetenderlist.html", gin.H{
		"bodyCss":"no-skin",
		"mainMenu" :"招标项目管理",
		"sessionInfo":controllers.GetSessionUser(c),
	})
}

func InviteTenderDetail(c *gin.Context){
	projectId, _ := c.GetQuery("id")
	log.Printf("projectId=%s\n", projectId)
	projectIdInt64, _ := strconv.ParseInt(projectId, 10, 64)
	projectDetail := yy.GetInviteTenderById( projectIdInt64 )
	sessionInfo := frontend.GetSessionInfo(c)
	if sessionInfo.UserName == "" {
		c.Redirect(http.StatusFound, "/userlogin")
		return
	}
	if sessionInfo.VipLevel == 0  {
		c.Redirect(http.StatusFound, "/novip")
		return
	}
	if sessionInfo.VipExpired == 1{
		c.Redirect(http.StatusFound, "/vipexpired")
		return
	}
	log.Printf("%+v", projectDetail)
	c.HTML(http.StatusOK, "projectdetail.html", gin.H{
		"projectDetail" : projectDetail,
		"path":"",
		"sessionInfo":frontend.GetSessionInfo(c),
	})
}

func InviteTenderList(c *gin.Context){
	page := models.PageRequest{}
	c.Bind( &page )
	responsePage := yy.GetAllInviteTenderByPage(&page)
	c.JSON(http.StatusOK, responsePage)
}

func InviteTenderEdit(c * gin.Context){
	tender := yy.YyPorject{}
	c.Bind( &tender)
	tender.ProjectType = 1
	openTenderTime := c.PostForm("InviteOpenTenderTime")
	tender.InviteOpenTenderTime.UnmarshalText( []byte(openTenderTime))
	submitTenderEndTime := c.PostForm("InviteSubmitTenderEndTime")
	tender.InviteSubmitTenderEndTime.UnmarshalText([]byte(submitTenderEndTime))
	oper, _ := c.GetPostForm("oper")
	if "edit" == oper {
		id, _:= c.GetPostForm("id")
		tender.Id, _= strconv.ParseInt(id, 10, 64)
		_, err := models.DbMaster.ID(tender.Id).Update(&tender)
		if err != nil {
			log.Printf("更新Purchase报错:%s\n",models.GetErrorInfo(err))
		}
	}else if "add" == oper {
		tender.CreateTime = models.Date{}
		_, err := models.DbMaster.Insert( &tender)
		if err != nil {
			log.Printf("新增Purchase报错:%s\n",models.GetErrorInfo(err))
			c.String(http.StatusInternalServerError, "%s", "fail")
			return
		}
		c.String(http.StatusOK, "%s", "success")
	} else if "del" == oper{
		id, _:= c.GetPostForm("id")
		tender.Id, _= strconv.ParseInt(id, 10, 64)
		_, err := models.DbMaster.Delete(&tender)
		if err != nil {
			log.Printf("删除Org报错:%s\n",models.GetErrorInfo(err))
			c.String(http.StatusInternalServerError, "%s", "fail")
			return
		}
		c.String(http.StatusOK, "%s", "success")
	}
}



func PushInviteTenderProject( c * gin.Context ){
	type ProjectInfo struct{
		IsSaleBidServicefee string `json:"isSaleBidServicefee"`
		PreBidServicePayStatus string `json:"preBidServicePayStatus"`
		SignUpStartTime string `json:"signUpStartTime"`
		IndustryName string `json:"industryName"`
		BidStartTime string `json:"bidStartTime"`
		PreDocDownloadStartTime string `json:"preDocDownloadStartTime"`
		IsSignUp string `json:"isSignUp"`
		PreTechnicalOpenBidStartTime string `json:"preTechnicalOpenBidStartTime"`
		ProjectAreaName string `json:"projectAreaName"`
		PreisSaleDocfee string `json:"preisSaleDocfee"`
		BidServicePayStatusMemo string `json:"bidServicePayStatusMemo"`
		AgenciesId string `json:"agenciesId"`
		BidServicePayStatus string `json:"bidServicePayStatus"`
		PreBidServiceFeeOrderNo string `json:"preBidServiceFeeOrderNo"`
		PreSignUpEndTime string `json:"preSignUpEndTime"`
		PreStageId string `json:"preStageId"`
		BidStatus string `json:"bidStatus"`
		PreSignUpStartTime string `json:"preSignUpStartTime"`
		DocSaleEndTime string `json:"docSaleEndTime"`
		BidServiceFeeOrderNo string `json:"bidServiceFeeOrderNo"`
		DocDownloadStartTime string `json:"docDownloadStartTime"`
		BidDocFeeOrderNo string `json:"bidDocFeeOrderNo"`
		BidDocPayStatusMemo string `json:"bidDocPayStatusMemo"`
		AgentCode string `json:"agentCode"`
		IsPack string `json:"isPack"`
		StageType string `json:"stageType"`
		PalceAddress string `json:"palceAddress"`
		LimitPrice string `json:"limitPrice"`
		UseStatus string `json:"useStatus"`
		DocSaleStartTime string `json:"docSaleStartTime"`
		PreBidStartTime string `json:"preBidStartTime"`
		CreateTime string 	`json:"createTime"`
		Buyersid string	`json:"buyersid"`
		PurchaserCode string	`json:"purchaserCode"`
		TenderMethod string	`json:"tenderMethod"`
		PackInfoList string	`json:"packInfoList"`
		StageId string	`json:"stageId"`
		QualificationMethod string	`json:"qualificationMethod"`
		PreIsSaleBidServicefee string	`json:"preIsSaleBidServicefee"`
		IsTwoBidOpening string	`json:"isTwoBidOpening"`
		TenderNo string	`json:"tenderNo"`
		PreBidDocPayStatus string	`json:"preBidDocPayStatus"`
		TenderType string	`json:"tenderType"`
		ArchiveStatus string	`json:"archiveStatus"`
		PreBidServicePayStatusMemo string	`json:"preBidServicePayStatusMemo"`
		PreDocSaleEndTime string	`json:"preDocSaleEndTime"`
		OpenBidStartTime string	`json:"openBidStartTime"`
		PreOpenBidStartTime string	`json:"preOpenBidStartTime"`
		PrebidStatus string	`json:"prebidStatus"`
		BudgetMoney string	`json:"budgetMoney"`
		IsSaleDocfee string	`json:"isSaleDocfee"`
		DocDownloadEndTime string	`json:"docDownloadEndTime"`
		TenderId string	`json:"tenderId"`
		PurchaserName string	`json:"purchaserName"`
		PreBidDocPayStatusMemo string	`json:"preBidDocPayStatusMemo"`
		PurcategoryNames string	`json:"purcategoryNames"`
		PreIsHaveBidDoc string	`json:"preIsHaveBidDoc"`
		PreIsSignUp string	`json:"preIsSignUp"`
		IsRemoteOpening string	`json:"isRemoteOpening"`
		PreBidDocFeeOrderNo string	`json:"preBidDocFeeOrderNo"`
		TenderNoNumber string	`json:"tenderNoNumber"`
		AgentName string	`json:"agentName"`
		BuyersName string `json:"buyersName"`
		BidEndTime string `json:"bidEndTime"`
		IsHaveBidDoc string `json:"isHaveBidDoc"`
		BidDocPayStatus string `json:"bidDocPayStatus"`
		TenderName string `json:"tenderName"`
		PreDocSaleStartTime string `json:"preDocSaleStartTime"`
		BidBond string  `json:"bidBond"`
		SignUpEndTime string `json:"signUpEndTime"`
		OpenBidUnPriceStartTime string `json:"openBidUnPriceStartTime"`
		PreBidEndTime string `json:"preBidEndTime"`
		PreDocDownloadEndTime string `json:"preDocDownloadEndTime"`
	}
	type Data struct {
		TenderProjectInfo  ProjectInfo `json:"tenderProjectInfo"`
	}
	type ProjecgNo struct {
		ProjectNo string `json:"projectNo"`
	}
	project := ProjecgNo{}

	err := c.BindJSON( &project)
	if err != nil {
		log.Printf("获取项目时绑定参数出错%s\n",models.GetErrorInfo(err))
	}

	fmt.Printf("proejctNo=%s", project.ProjectNo)


	{//项目同步
		data := make(url.Values)
		data["tenderNo"] = []string{fmt.Sprintf("%d",project.ProjectNo)}
		data["userName"] = []string{"daili"}
		data["password"] = []string{"MTIzNDU2"}
		data["tenderNoNumber"] = []string{"9f26ce0f5ce44146b42340ea31331fcf"}

		res, err := http.PostForm("http://219.239.33.98:8080/yyg/tenderProjectInfoHS.do?getProjectInfoByCode", data)
		log.Printf("!!!!!!!!!!!!!!!!!!%s", err)
		if err == nil {
			data := Data{}
			//projectInfo := ProjectInfo{}
			//data.TenderProjectInfo = projectInfo
			//err1 := Bind(res, &data)
			body, err1 := ioutil.ReadAll(res.Body)
			json.Unmarshal(body, &data)
			log.Printf("data=%v\n", data)
			if err1 != nil {
				log.Printf("err1=%s\n", err1)
			}
		} else {
			log.Printf("err=%s\n", err)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"success":true,
		"msg":fmt.Sprintf("调用成功,projectNo=%s", project.ProjectNo),
	})
}