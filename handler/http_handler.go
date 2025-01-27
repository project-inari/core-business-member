package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/project-inari/core-business-member/dto"
	"github.com/project-inari/core-business-member/pkg/request"
	"github.com/project-inari/core-business-member/pkg/response"
)

type httpHandler struct {
	d Dependencies
}

func newHTTPHandler(d Dependencies) *httpHandler {
	return &httpHandler{
		d: d,
	}
}

func (h *httpHandler) Invite(c echo.Context) error {
	ctx := c.Request().Context()
	wrapper := request.ContextWrapper(c)

	req := new(dto.InviteReq)
	if err := wrapper.Bind(req); err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("error - [Invite] bad request: %v", err), "")
	}

	res, err := h.d.Service.Invite(ctx, *req)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("error - [Invite] internal server error: %v", err), "")
	}

	return response.SuccessResponse(c, http.StatusOK, res)
}

func (h *httpHandler) AcceptInvite(c echo.Context) error {
	ctx := c.Request().Context()
	wrapper := request.ContextWrapper(c)

	req := new(dto.AcceptInviteReq)
	if err := wrapper.Bind(req); err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("error - [AcceptInvite] bad request: %v", err), "")
	}

	res, err := h.d.Service.AcceptInvite(ctx, *req)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("error - [AcceptInvite] internal server error: %v", err), "")
	}

	return response.SuccessResponse(c, http.StatusOK, res)
}

func (h *httpHandler) DeclineInvite(c echo.Context) error {
	ctx := c.Request().Context()
	wrapper := request.ContextWrapper(c)

	req := new(dto.DeclineInviteReq)
	if err := wrapper.Bind(req); err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("error - [DeclineInvite] bad request: %v", err), "")
	}

	res, err := h.d.Service.DeclineInvite(ctx, *req)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("error - [DeclineInvite] internal server error: %v", err), "")
	}

	return response.SuccessResponse(c, http.StatusOK, res)
}
