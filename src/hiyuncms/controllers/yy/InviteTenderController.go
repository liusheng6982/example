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
		isSaleBidServicefee string
		preBidServicePayStatus string
		signUpStartTime string
		industryName string
		BidStartTime string `json:"bidStartTime"`
		PreDocDownloadStartTime string `json:"preDocDownloadStartTime"`
		isSignUp string `json:"is_sign_up"`
		preTechnicalOpenBidStartTime string
		projectAreaName string
		preisSaleDocfee string
		bidServicePayStatusMemo string
		agenciesId string
		bidServicePayStatus string
		preBidServiceFeeOrderNo string
		preSignUpEndTime string
		preStageId string
		bidStatus string
		preSignUpStartTime string
		docSaleEndTime string
		bidServiceFeeOrderNo string
		docDownloadStartTime string
		bidDocFeeOrderNo string
		bidDocPayStatusMemo string
		agentCode string
		isPack string
		stageType string
		palceAddress string
		limitPrice string
		useStatus string
		docSaleStartTime string
		preBidStartTime string
		createTime string
		buyersid string
		purchaserCode string
		tenderMethod string
		packInfoList string
		stageId string
		qualificationMethod string
		preIsSaleBidServicefee string
		isTwoBidOpening string
		tenderNo string
		preBidDocPayStatus string
		tenderType string
		archiveStatus string
		preBidServicePayStatusMemo string
		preDocSaleEndTime string
		openBidStartTime string
		preOpenBidStartTime string
		prebidStatus string
		budgetMoney string
		isSaleDocfee string
		docDownloadEndTime string
		tenderId string
		purchaserName string
		preBidDocPayStatusMemo string
		purcategoryNames string
		preIsHaveBidDoc string
		preIsSignUp string
		isRemoteOpening string
		preBidDocFeeOrderNo string
		tenderNoNumber string
		agentName string
		BuyersName string `json:"buyersName"`
		BidEndTime string `json:"bidEndTime"`
		IsHaveBidDoc string `json:"isHaveBidDoc"`
		BidDocPayStatus string `json:"bidDocPayStatus"`
		TenderName string `json:"tenderName"`
		PreDocSaleStartTime string `json:"preDocSaleStartTime"`
		bidBond string
		signUpEndTime string
		openBidUnPriceStartTime string
		preBidEndTime string
		preDocDownloadEndTime string
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