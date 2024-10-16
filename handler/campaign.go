package handler

import (
	"net/http"
	"startup-api/API/campaign"
	"startup-api/helper"
	"strconv"

	"github.com/gin-gonic/gin"
)

type campaignHandler struct {
	campaignService campaign.CampaignService
}

func NewCampaignHandler(campaignService campaign.CampaignService) *campaignHandler {
	return &campaignHandler{campaignService}
}

func (h *campaignHandler) GetCampaign(c *gin.Context) {
	usrID, _ := strconv.Atoi(c.Query("user_id"))

	campaign, err := h.campaignService.FindCampaign(usrID)
	if err != nil {
		res := helper.APIResponse("Failed to get campaign data", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	res := helper.APIResponse("List of campaigns", http.StatusOK, "success", campaign)
	c.JSON(http.StatusOK, res)
}
