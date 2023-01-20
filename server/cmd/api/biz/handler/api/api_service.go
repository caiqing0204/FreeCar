// Code generated by hertz generator.

package api

import (
	"context"
	"time"

	"github.com/CyanAsterisk/FreeCar/server/cmd/api/biz/errno"
	"github.com/CyanAsterisk/FreeCar/server/cmd/api/biz/model/server/cmd/api"
	"github.com/CyanAsterisk/FreeCar/server/cmd/api/global"
	"github.com/CyanAsterisk/FreeCar/server/cmd/api/kitex_gen/auth"
	"github.com/CyanAsterisk/FreeCar/server/cmd/api/kitex_gen/car"
	"github.com/CyanAsterisk/FreeCar/server/cmd/api/kitex_gen/profile"
	"github.com/CyanAsterisk/FreeCar/server/cmd/api/kitex_gen/trip"
	models "github.com/CyanAsterisk/FreeCar/server/cmd/api/model"
	"github.com/CyanAsterisk/FreeCar/shared/consts"
	"github.com/CyanAsterisk/FreeCar/shared/middleware"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/golang-jwt/jwt"
)

// Login .
// @router /auth/login [POST]
func Login(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.LoginRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		errno.SendResponse(c, errno.BindAndValidateFail, nil)
		return
	}
	// rpc to get accountID
	resp, err := global.AuthClient.Login(ctx, &auth.LoginRequest{Code: req.Code})
	if err != nil {
		errno.SendResponse(c, errno.RequestServerFail, nil)
		return
	}
	// create a JWT
	j := middleware.NewJWT()
	claims := models.CustomClaims{
		ID: resp.AccountID,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),
			ExpiresAt: time.Now().Unix() + consts.ThirtyDays,
			Issuer:    consts.JWTIssuer,
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		errno.SendResponse(c, errno.GenerateTokenFail, nil)
		return
	}
	// return token
	errno.SendResponse(c, errno.Success, api.LoginResponse{
		Token:     token,
		ExpiredAt: time.Now().Unix() + consts.ThirtyDays,
	})
}

// CreateCar .
// @router /car [POST]
func CreateCar(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.CreateCarRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		errno.SendResponse(c, errno.BindAndValidateFail, nil)
		return
	}
	aid, flag := c.Get(consts.AccountID)
	if !flag {
		errno.SendResponse(c, errno.ParamErr, nil)
		return
	}

	resp, err := global.CarClient.CreateCar(ctx, &car.CreateCarRequest{
		AccountId: aid.(int64),
		PlateNum:  req.PlateNum,
	})
	if err != nil {
		errno.SendResponse(c, errno.RequestServerFail, nil)
		return
	}
	errno.SendResponse(c, errno.Success, resp)
}

// GetCar .
// @router /car [GET]
func GetCar(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.GetCarRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		errno.SendResponse(c, errno.BindAndValidateFail, nil)
		return
	}
	aid, flag := c.Get(consts.AccountID)
	if !flag {
		errno.SendResponse(c, errno.ParamErr, nil)
		return
	}

	resp, err := global.CarClient.GetCar(ctx, &car.GetCarRequest{
		AccountId: aid.(int64),
		Id:        req.Id,
	})
	if err != nil {
		errno.SendResponse(c, errno.RequestServerFail, nil)
		return
	}
	errno.SendResponse(c, errno.Success, resp)
}

// GetCars .
// @router /cars [GET]
func GetCars(ctx context.Context, c *app.RequestContext) {
	aid, flag := c.Get(consts.AccountID)
	if !flag {
		errno.SendResponse(c, errno.AuthorizeFail, nil)
		return
	}
	resp, err := global.CarClient.GetCars(ctx, &car.GetCarsRequest{AccountId: aid.(int64)})
	if err != nil {
		errno.SendResponse(c, errno.RequestServerFail, nil)
		return
	}
	errno.SendResponse(c, errno.Success, resp)
}

// GetProfile .
// @router /profile [GET]
func GetProfile(ctx context.Context, c *app.RequestContext) {
	var err error

	aid, flag := c.Get(consts.AccountID)
	if !flag {
		errno.SendResponse(c, errno.ParamErr, nil)
		return
	}

	resp, err := global.ProfileClient.GetProfile(ctx, &profile.GetProfileRequest{AccountId: aid.(int64)})
	if err != nil {
		errno.SendResponse(c, errno.RequestServerFail, nil)
		return
	}

	errno.SendResponse(c, errno.Success, resp)
}

// SubmitProfile .
// @router /profile [POST]
func SubmitProfile(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.SubmitProfileRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		errno.SendResponse(c, errno.BindAndValidateFail, nil)
		return
	}
	aid, flag := c.Get(consts.AccountID)
	if !flag {
		errno.SendResponse(c, errno.ParamErr, nil)
		return
	}

	resp, err := global.ProfileClient.SubmitProfile(ctx, &profile.SubmitProfileRequest{
		AccountId: aid.(int64),
		Identity: &profile.Identity{
			LicNumber:       req.Identity.LicNumber,
			Name:            req.Identity.Name,
			Gender:          profile.Gender(req.Identity.Gender),
			BirthDateMillis: req.Identity.BirthDateMillis,
		},
	})
	if err != nil {
		errno.SendResponse(c, errno.RequestServerFail, nil)
		return
	}

	errno.SendResponse(c, errno.Success, resp)
}

// ClearProfile .
// @router /profile [DELETE]
func ClearProfile(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.ClearProfileRequest
	err = c.BindAndValidate(&req)
	aid, flag := c.Get(consts.AccountID)
	if !flag {
		errno.SendResponse(c, errno.ParamErr, nil)
		return
	}

	resp, err := global.ProfileClient.ClearProfile(ctx, &profile.ClearProfileRequest{AccountId: aid.(int64)})
	if err != nil {
		errno.SendResponse(c, errno.RequestServerFail, nil)
		return
	}

	errno.SendResponse(c, errno.Success, resp)
}

// GetProfilePhoto .
// @router /profile/photo [GET]
func GetProfilePhoto(ctx context.Context, c *app.RequestContext) {
	var err error
	aid, flag := c.Get(consts.AccountID)
	if !flag {
		errno.SendResponse(c, errno.ParamErr, nil)
		return
	}

	resp, err := global.ProfileClient.GetProfilePhoto(ctx, &profile.GetProfilePhotoRequest{AccountId: aid.(int64)})
	if err != nil {
		errno.SendResponse(c, errno.RequestServerFail, nil)
		return
	}

	errno.SendResponse(c, errno.Success, resp)
}

// CreateProfilePhoto .
// @router /profile/photo [POST]
func CreateProfilePhoto(ctx context.Context, c *app.RequestContext) {
	var err error
	aid, flag := c.Get(consts.AccountID)
	if !flag {
		errno.SendResponse(c, errno.ParamErr, nil)
		return
	}

	resp, err := global.ProfileClient.CreateProfilePhoto(ctx, &profile.CreateProfilePhotoRequest{AccountId: aid.(int64)})
	if err != nil {
		errno.SendResponse(c, errno.RequestServerFail, nil)
		return
	}

	errno.SendResponse(c, errno.Success, resp)
}

// CompleteProfilePhoto .
// @router /profile/photo/complete [POST]
func CompleteProfilePhoto(ctx context.Context, c *app.RequestContext) {
	var err error
	aid, flag := c.Get(consts.AccountID)
	if !flag {
		errno.SendResponse(c, errno.ParamErr, nil)
		return
	}

	resp, err := global.ProfileClient.CompleteProfilePhoto(ctx, &profile.CompleteProfilePhotoRequest{AccountId: aid.(int64)})
	if err != nil {
		errno.SendResponse(c, errno.RequestServerFail, nil)
		return
	}

	errno.SendResponse(c, errno.Success, resp)
}

// ClearProfilePhoto .
// @router /profile/photo [DELETE]
func ClearProfilePhoto(ctx context.Context, c *app.RequestContext) {
	var err error
	aid, flag := c.Get(consts.AccountID)
	if !flag {
		errno.SendResponse(c, errno.ParamErr, nil)
		return
	}

	resp, err := global.ProfileClient.ClearProfilePhoto(ctx, &profile.ClearProfilePhotoRequest{AccountId: aid.(int64)})
	if err != nil {
		errno.SendResponse(c, errno.RequestServerFail, nil)
		return
	}

	errno.SendResponse(c, errno.Success, resp)
}

// CreateTrip .
// @router /trip [POST]
func CreateTrip(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.CreateTripRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		errno.SendResponse(c, errno.BindAndValidateFail, nil)
		return
	}
	aid, flag := c.Get(consts.AccountID)
	if !flag {
		errno.SendResponse(c, errno.ParamErr, nil)
		return
	}

	resp, err := global.TripClient.CreateTrip(ctx, &trip.CreateTripRequest{
		Start: &trip.Location{
			Latitude:  req.Start.Latitude,
			Longitude: req.Start.Longitude,
		},
		CarId:     req.CarId,
		AvatarUrl: req.AvatarUrl,
		AccountId: aid.(int64),
	})
	if err != nil {
		errno.SendResponse(c, errno.RequestServerFail, nil)
		return
	}

	errno.SendResponse(c, errno.Success, resp)
}

// GetTrip .
// @router /trip/:id [GET]
func GetTrip(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.GetTripRequest
	aid, flag := c.Get(consts.AccountID)
	if !flag {
		errno.SendResponse(c, errno.ParamErr, nil)
		return
	}
	req.Id = c.Param("id")

	resp, err := global.TripClient.GetTrip(ctx, &trip.GetTripRequest{
		Id:        req.Id,
		AccountId: aid.(int64),
	})
	if err != nil {
		errno.SendResponse(c, errno.RequestServerFail, nil)
		return
	}

	errno.SendResponse(c, errno.Success, resp)
}

// GetTrips .
// @router /trips [GET]
func GetTrips(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.GetTripsRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		errno.SendResponse(c, errno.BindAndValidateFail, nil)
		return
	}
	aid, flag := c.Get(consts.AccountID)
	if !flag {
		errno.SendResponse(c, errno.ParamErr, nil)
		return
	}

	resp, err := global.TripClient.GetTrips(ctx, &trip.GetTripsRequest{
		Status:    trip.TripStatus(req.Status),
		AccountId: aid.(int64),
	})
	if err != nil {
		errno.SendResponse(c, errno.RequestServerFail, nil)
		return
	}

	errno.SendResponse(c, errno.Success, resp)
}

// UpdateTrip .
// @router /trip/:id [PUT]
func UpdateTrip(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.UpdateTripRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		errno.SendResponse(c, errno.BindAndValidateFail, nil)
		return
	}
	aid, flag := c.Get(consts.AccountID)
	if !flag {
		errno.SendResponse(c, errno.ParamErr, nil)
		return
	}
	req.Id = c.Param(consts.ID)

	resp, err := global.TripClient.UpdateTrip(ctx, &trip.UpdateTripRequest{
		Id: req.Id,
		Current: &trip.Location{
			Latitude:  req.Current.Latitude,
			Longitude: req.Current.Longitude,
		},
		EndTrip:   req.EndTrip,
		AccountId: aid.(int64),
	})
	if err != nil {
		errno.SendResponse(c, errno.RequestServerFail, nil)
		return
	}

	errno.SendResponse(c, errno.Success, resp)
}
