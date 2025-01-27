package handler

import (
	"github.com/labstack/echo/v4"
)

func (h *httpHandler) initRoutes(e *echo.Echo) {
	e.GET("/health", h.HealthCheck)

	v1 := e.Group("/v1")
	v1.POST("/invite", h.Invite)
	v1.POST("/invite/accept", h.AcceptInvite)
	v1.POST("/invite/decline", h.DeclineInvite)
	v1.GET("/joining", h.JoiningInquiry)
	v1.GET("/members/:businessName", h.MemberInquiry)
	v1.GET("/joined/:username", h.UserJoinedInquiry)
}
