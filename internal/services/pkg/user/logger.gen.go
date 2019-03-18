package user

// DO NOT EDIT!
// This code is generated with http://github.com/hexdigest/gowrap tool
// using ../../../../tools/templates/logger template

//go:generate gowrap gen -p go.zenithar.org/spotigraph/internal/services -i User -t ../../../../tools/templates/logger -o logger.gen.go

import (
	"context"

	"go.uber.org/zap"
	"go.zenithar.org/pkg/log"
	"go.zenithar.org/spotigraph/internal/services"
	"go.zenithar.org/spotigraph/pkg/protocol/v1/spotigraph"
)

// UserLogger implements services.User that is instrumented with logger
type UserLogger struct {
	_log  log.Logger
	_base services.User
}

// NewUserLogger instruments an implementation of the services.User with simple logging
func NewUserLogger(base services.User, logger log.Logger) UserLogger {
	return UserLogger{
		_base: base,
		_log:  logger.With(zap.String("decorator", "logger")),
	}
}

// Create implements services.User
func (_d UserLogger) Create(ctx context.Context, req *spotigraph.UserCreateReq) (sp1 *spotigraph.SingleUserRes, err error) {
	_d._log.Debug("UserLogger: calling Create", zap.String("method", "Create"))
	defer func() {
		if err != nil {
			_d._log.Error("UserLogger: method Create returned an error", zap.String("method", "Create"), zap.Error(err), zap.Any("result", map[string]interface{}{
				"sp1": sp1,
				"err": err}))
		} else {
			_d._log.Error("UserLogger: method Create finished", zap.String("method", "Create"), zap.Any("result", map[string]interface{}{
				"sp1": sp1,
				"err": err}))
		}
	}()
	return _d._base.Create(ctx, req)
}

// Delete implements services.User
func (_d UserLogger) Delete(ctx context.Context, req *spotigraph.UserGetReq) (ep1 *spotigraph.EmptyRes, err error) {
	_d._log.Debug("UserLogger: calling Delete", zap.String("method", "Delete"))
	defer func() {
		if err != nil {
			_d._log.Error("UserLogger: method Delete returned an error", zap.String("method", "Delete"), zap.Error(err), zap.Any("result", map[string]interface{}{
				"ep1": ep1,
				"err": err}))
		} else {
			_d._log.Error("UserLogger: method Delete finished", zap.String("method", "Delete"), zap.Any("result", map[string]interface{}{
				"ep1": ep1,
				"err": err}))
		}
	}()
	return _d._base.Delete(ctx, req)
}

// Get implements services.User
func (_d UserLogger) Get(ctx context.Context, req *spotigraph.UserGetReq) (sp1 *spotigraph.SingleUserRes, err error) {
	_d._log.Debug("UserLogger: calling Get", zap.String("method", "Get"))
	defer func() {
		if err != nil {
			_d._log.Error("UserLogger: method Get returned an error", zap.String("method", "Get"), zap.Error(err), zap.Any("result", map[string]interface{}{
				"sp1": sp1,
				"err": err}))
		} else {
			_d._log.Error("UserLogger: method Get finished", zap.String("method", "Get"), zap.Any("result", map[string]interface{}{
				"sp1": sp1,
				"err": err}))
		}
	}()
	return _d._base.Get(ctx, req)
}

// Update implements services.User
func (_d UserLogger) Update(ctx context.Context, req *spotigraph.UserUpdateReq) (sp1 *spotigraph.SingleUserRes, err error) {
	_d._log.Debug("UserLogger: calling Update", zap.String("method", "Update"))
	defer func() {
		if err != nil {
			_d._log.Error("UserLogger: method Update returned an error", zap.String("method", "Update"), zap.Error(err), zap.Any("result", map[string]interface{}{
				"sp1": sp1,
				"err": err}))
		} else {
			_d._log.Error("UserLogger: method Update finished", zap.String("method", "Update"), zap.Any("result", map[string]interface{}{
				"sp1": sp1,
				"err": err}))
		}
	}()
	return _d._base.Update(ctx, req)
}
