package user

import (
	"bitPay/logger"
	"github.com/gin-gonic/gin"
)

func CreateAgency(c *gin.Context) {
	logger.Info("Enter to CreateAgency controller successfully")


	//agency := users.Agency{}
	//if err := c.ShouldBindJSON(&agency); err != nil {
	//	logger.Error("error when trying to bind json", err)
	//	restError := responses.NewBadRequestError("Invalid json structure","please make sure to send correct json body", http.StatusBadRequest)
	//	c.JSON(http.StatusBadRequest, restError )
	//	return
	//}
	//
	//body, _ := json.Marshal(agency)
	//request := requests.Request{}
	//request.Init(parent.Id, c.Request.Host+c.Request.URL.Path, string(body))
	//if saveErr := request.Save(); saveErr != nil {
	//	saveErr.SetRequestId(request.Token)
	//	c.JSON(saveErr.Status(), saveErr)
	//	return
	//}
	//
	//agent, serviceErr := services.UsersService.CreateAgency(agency)
	//if serviceErr != nil {
	//	serviceErr.SetRequestId(request.Token)
	//	c.JSON(serviceErr.Status(), serviceErr)
	//	return
	//}
	//
	//ok := responses.NewRequestSuccessOk("Agency Created successfully", "", agent)
	//ok.SetRequestId(request.Token)
	//c.JSON(http.StatusOK,  ok)
	logger.Info("Close from CreateAgency controller successfully")
}

