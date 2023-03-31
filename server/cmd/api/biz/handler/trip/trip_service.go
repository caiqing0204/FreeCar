// Code generated by hertz generator.

package trip

import (
	"context"
	"net/http"

	htrip "github.com/CyanAsterisk/FreeCar/server/cmd/api/biz/model/trip"
	"github.com/CyanAsterisk/FreeCar/server/cmd/api/config"
	"github.com/CyanAsterisk/FreeCar/server/cmd/api/pkg"
	"github.com/CyanAsterisk/FreeCar/server/shared/consts"
	"github.com/CyanAsterisk/FreeCar/server/shared/errno"
	kbase "github.com/CyanAsterisk/FreeCar/server/shared/kitex_gen/base"
	ktrip "github.com/CyanAsterisk/FreeCar/server/shared/kitex_gen/trip"
	"github.com/CyanAsterisk/FreeCar/server/shared/tools"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

// DeleteTrip .
// @router /trip/admin/trip [DELETE]
func DeleteTrip(ctx context.Context, c *app.RequestContext) {
	var err error
	var req htrip.DeleteTripRequest
	resp := new(ktrip.DeleteTripResponse)

	if err = c.BindAndValidate(&req); err != nil {
		resp.BaseResp = tools.BuildBaseResp(errno.ParamsErr)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp, err = config.GlobalTripClient.DeleteTrip(ctx, &ktrip.DeleteTripRequest{Id: req.ID})
	if err != nil {
		hlog.Error("rpc trip service err", err)
		resp.BaseResp = tools.BuildBaseResp(errno.ServiceErr)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	c.JSON(http.StatusOK, resp)
}

// GetAllTrips .
// @router /trip/admin/all [GET]
func GetAllTrips(ctx context.Context, c *app.RequestContext) {
	var err error
	var req htrip.GetAllTripsRequest
	resp := new(ktrip.GetAllTripsResponse)

	if err = c.BindAndValidate(&req); err != nil {
		resp.BaseResp = tools.BuildBaseResp(errno.ParamsErr)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp, err = config.GlobalTripClient.GetAllTrips(ctx, &ktrip.GetAllTripsRequest{})
	if err != nil {
		hlog.Error("rpc trip service err", err)
		resp.BaseResp = tools.BuildBaseResp(errno.ServiceErr)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	c.JSON(http.StatusOK, resp)
}

// GetSomeTrips .
// @router /trip/admin/some [GET]
func GetSomeTrips(ctx context.Context, c *app.RequestContext) {
	var err error
	var req htrip.GetSomeTripsRequest
	resp := new(ktrip.GetSomeTripsResponse)

	if err = c.BindAndValidate(&req); err != nil {
		resp.BaseResp = tools.BuildBaseResp(errno.ParamsErr)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp, err = config.GlobalTripClient.GetSomeTrips(ctx, &ktrip.GetSomeTripsRequest{})
	if err != nil {
		hlog.Error("rpc trip service err", err)
		resp.BaseResp = tools.BuildBaseResp(errno.ServiceErr)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	c.JSON(http.StatusOK, resp)
}

// CreateTrip .
// @router /trip/mini/trip [POST]
func CreateTrip(ctx context.Context, c *app.RequestContext) {
	var err error
	var req htrip.CreateTripRequest
	resp := new(ktrip.CreateTripResponse)

	if err = c.BindAndValidate(&req); err != nil {
		resp.BaseResp = tools.BuildBaseResp(errno.ParamsErr)
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	aid, flag := c.Get(consts.AccountID)
	if !flag {
		resp.BaseResp = tools.BuildBaseResp(errno.AuthorizeFail)
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	resp, err = config.GlobalTripClient.CreateTrip(ctx, &ktrip.CreateTripRequest{
		Start:     pkg.ConvertTripLocation(req.Start),
		CarId:     req.CarID,
		AvatarUrl: req.AvatarURL,
		AccountId: aid.(int64),
	})
	if err != nil {
		hlog.Error("rpc trip service err", err)
		resp.BaseResp = tools.BuildBaseResp(errno.ServiceErr)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	c.JSON(http.StatusOK, resp)
}

// GetTrip .
// @router /trip/mini/trip [GET]
func GetTrip(ctx context.Context, c *app.RequestContext) {
	var err error
	var req htrip.GetTripRequest
	resp := new(ktrip.GetTripResponse)

	if err = c.BindAndValidate(&req); err != nil {
		resp.BaseResp = tools.BuildBaseResp(errno.ParamsErr)
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	aid, flag := c.Get(consts.AccountID)
	if !flag {
		resp.BaseResp = tools.BuildBaseResp(errno.AuthorizeFail)
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	resp, err = config.GlobalTripClient.GetTrip(ctx, &ktrip.GetTripRequest{
		Id:        req.ID,
		AccountId: aid.(int64),
	})
	if err != nil {
		hlog.Error("rpc trip service err", err)
		resp.BaseResp = tools.BuildBaseResp(errno.ServiceErr)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	c.JSON(http.StatusOK, resp)
}

// GetTrips .
// @router /trip/mini/trips [GET]
func GetTrips(ctx context.Context, c *app.RequestContext) {
	var err error
	var req htrip.GetTripsRequest
	resp := new(ktrip.GetTripsResponse)

	if err = c.BindAndValidate(&req); err != nil {
		resp.BaseResp = tools.BuildBaseResp(errno.ParamsErr)
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	aid, flag := c.Get(consts.AccountID)
	if !flag {
		resp.BaseResp = tools.BuildBaseResp(errno.AuthorizeFail)
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	resp, err = config.GlobalTripClient.GetTrips(ctx, &ktrip.GetTripsRequest{
		Status:    kbase.TripStatus(req.Status),
		AccountId: aid.(int64),
	})
	if err != nil {
		hlog.Error("rpc trip service err", err)
		resp.BaseResp = tools.BuildBaseResp(errno.ServiceErr)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateTrip .
// @router /trip [PUT]
func UpdateTrip(ctx context.Context, c *app.RequestContext) {
	var err error
	var req htrip.UpdateTripRequest
	resp := new(ktrip.UpdateTripResponse)

	if err = c.BindAndValidate(&req); err != nil {
		resp.BaseResp = tools.BuildBaseResp(errno.ParamsErr)
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	aid, flag := c.Get(consts.AccountID)
	if !flag {
		resp.BaseResp = tools.BuildBaseResp(errno.AuthorizeFail)
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	resp, err = config.GlobalTripClient.UpdateTrip(ctx, &ktrip.UpdateTripRequest{
		Id:        req.ID,
		Current:   (*kbase.Location)(req.Current),
		EndTrip:   req.EndTrip,
		AccountId: aid.(int64),
	})
	if err != nil {
		hlog.Error("rpc trip service err", err)
		resp.BaseResp = tools.BuildBaseResp(errno.ServiceErr)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	c.JSON(http.StatusOK, resp)
}
