package tribe

// DO NOT EDIT!
// This code is generated with http://github.com/hexdigest/gowrap tool
// using ../../../../tools/templates/logger template

//go:generate gowrap gen -p go.zenithar.org/spotigraph/internal/services -i Tribe -t ../../../../tools/templates/logger -o logger.gen.go

import (
	"context"

	"go.uber.org/zap"
	"go.zenithar.org/pkg/log"
	"go.zenithar.org/spotigraph/internal/services"
	"go.zenithar.org/spotigraph/pkg/protocol/v1/spotigraph"
)

// TribeLogger implements services.Tribe that is instrumented with logger
type TribeLogger struct {
	_log  log.Logger
	_base services.Tribe
}

// NewTribeLogger instruments an implementation of the services.Tribe with simple logging
func NewTribeLogger(base services.Tribe, logger log.LoggerFactory) TribeLogger {
	return TribeLogger{
		_base: base,
		_log:  logger.With(zap.String("decorator", "logger")),
	}
}

// Create implements services.Tribe
func (_d TribeLogger) Create(ctx context.Context, req *spotigraph.TribeCreateReq) (sp1 *spotigraph.SingleTribeRes, err error) {
	_d._log.Debug("TribeLogger: calling Create", zap.String("method", "Create"))
	defer func() {
		if err != nil {
			_d._log.Error("TribeLogger: method Create returned an error", zap.String("method", "Create"), zap.Error(err), zap.Any("result", map[string]interface{}{
				"sp1": sp1,
				"err": err}))
		} else {
			_d._log.Error("TribeLogger: method Create finished", zap.String("method", "Create"), zap.Any("result", map[string]interface{}{
				"sp1": sp1,
				"err": err}))
		}
	}()
	return _d._base.Create(ctx, req)
}

// Delete implements services.Tribe
func (_d TribeLogger) Delete(ctx context.Context, req *spotigraph.TribeGetReq) (ep1 *spotigraph.EmptyRes, err error) {
	_d._log.Debug("TribeLogger: calling Delete", zap.String("method", "Delete"))
	defer func() {
		if err != nil {
			_d._log.Error("TribeLogger: method Delete returned an error", zap.String("method", "Delete"), zap.Error(err), zap.Any("result", map[string]interface{}{
				"ep1": ep1,
				"err": err}))
		} else {
			_d._log.Error("TribeLogger: method Delete finished", zap.String("method", "Delete"), zap.Any("result", map[string]interface{}{
				"ep1": ep1,
				"err": err}))
		}
	}()
	return _d._base.Delete(ctx, req)
}

// Get implements services.Tribe
func (_d TribeLogger) Get(ctx context.Context, req *spotigraph.TribeGetReq) (sp1 *spotigraph.SingleTribeRes, err error) {
	_d._log.Debug("TribeLogger: calling Get", zap.String("method", "Get"))
	defer func() {
		if err != nil {
			_d._log.Error("TribeLogger: method Get returned an error", zap.String("method", "Get"), zap.Error(err), zap.Any("result", map[string]interface{}{
				"sp1": sp1,
				"err": err}))
		} else {
			_d._log.Error("TribeLogger: method Get finished", zap.String("method", "Get"), zap.Any("result", map[string]interface{}{
				"sp1": sp1,
				"err": err}))
		}
	}()
	return _d._base.Get(ctx, req)
}

// Update implements services.Tribe
func (_d TribeLogger) Update(ctx context.Context, req *spotigraph.TribeUpdateReq) (sp1 *spotigraph.SingleTribeRes, err error) {
	_d._log.Debug("TribeLogger: calling Update", zap.String("method", "Update"))
	defer func() {
		if err != nil {
			_d._log.Error("TribeLogger: method Update returned an error", zap.String("method", "Update"), zap.Error(err), zap.Any("result", map[string]interface{}{
				"sp1": sp1,
				"err": err}))
		} else {
			_d._log.Error("TribeLogger: method Update finished", zap.String("method", "Update"), zap.Any("result", map[string]interface{}{
				"sp1": sp1,
				"err": err}))
		}
	}()
	return _d._base.Update(ctx, req)
}
