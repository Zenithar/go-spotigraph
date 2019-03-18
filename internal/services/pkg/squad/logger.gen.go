package squad

// DO NOT EDIT!
// This code is generated with http://github.com/hexdigest/gowrap tool
// using ../../../../tools/templates/logger template

//go:generate gowrap gen -p go.zenithar.org/spotigraph/internal/services -i Squad -t ../../../../tools/templates/logger -o logger.gen.go

import (
	"context"

	"go.uber.org/zap"
	"go.zenithar.org/pkg/log"
	"go.zenithar.org/spotigraph/internal/services"
	"go.zenithar.org/spotigraph/pkg/protocol/v1/spotigraph"
)

// SquadLogger implements services.Squad that is instrumented with logger
type SquadLogger struct {
	_log  log.Logger
	_base services.Squad
}

// NewSquadLogger instruments an implementation of the services.Squad with simple logging
func NewSquadLogger(base services.Squad, logger log.LoggerFactory) SquadLogger {
	return SquadLogger{
		_base: base,
		_log:  logger.With(zap.String("decorator", "logger")),
	}
}

// Create implements services.Squad
func (_d SquadLogger) Create(ctx context.Context, req *spotigraph.SquadCreateReq) (sp1 *spotigraph.SingleSquadRes, err error) {
	_d._log.Debug("SquadLogger: calling Create", zap.String("method", "Create"))
	defer func() {
		if err != nil {
			_d._log.Error("SquadLogger: method Create returned an error", zap.String("method", "Create"), zap.Error(err), zap.Any("result", map[string]interface{}{
				"sp1": sp1,
				"err": err}))
		} else {
			_d._log.Error("SquadLogger: method Create finished", zap.String("method", "Create"), zap.Any("result", map[string]interface{}{
				"sp1": sp1,
				"err": err}))
		}
	}()
	return _d._base.Create(ctx, req)
}

// Delete implements services.Squad
func (_d SquadLogger) Delete(ctx context.Context, req *spotigraph.SquadGetReq) (ep1 *spotigraph.EmptyRes, err error) {
	_d._log.Debug("SquadLogger: calling Delete", zap.String("method", "Delete"))
	defer func() {
		if err != nil {
			_d._log.Error("SquadLogger: method Delete returned an error", zap.String("method", "Delete"), zap.Error(err), zap.Any("result", map[string]interface{}{
				"ep1": ep1,
				"err": err}))
		} else {
			_d._log.Error("SquadLogger: method Delete finished", zap.String("method", "Delete"), zap.Any("result", map[string]interface{}{
				"ep1": ep1,
				"err": err}))
		}
	}()
	return _d._base.Delete(ctx, req)
}

// Get implements services.Squad
func (_d SquadLogger) Get(ctx context.Context, req *spotigraph.SquadGetReq) (sp1 *spotigraph.SingleSquadRes, err error) {
	_d._log.Debug("SquadLogger: calling Get", zap.String("method", "Get"))
	defer func() {
		if err != nil {
			_d._log.Error("SquadLogger: method Get returned an error", zap.String("method", "Get"), zap.Error(err), zap.Any("result", map[string]interface{}{
				"sp1": sp1,
				"err": err}))
		} else {
			_d._log.Error("SquadLogger: method Get finished", zap.String("method", "Get"), zap.Any("result", map[string]interface{}{
				"sp1": sp1,
				"err": err}))
		}
	}()
	return _d._base.Get(ctx, req)
}

// Update implements services.Squad
func (_d SquadLogger) Update(ctx context.Context, req *spotigraph.SquadUpdateReq) (sp1 *spotigraph.SingleSquadRes, err error) {
	_d._log.Debug("SquadLogger: calling Update", zap.String("method", "Update"))
	defer func() {
		if err != nil {
			_d._log.Error("SquadLogger: method Update returned an error", zap.String("method", "Update"), zap.Error(err), zap.Any("result", map[string]interface{}{
				"sp1": sp1,
				"err": err}))
		} else {
			_d._log.Error("SquadLogger: method Update finished", zap.String("method", "Update"), zap.Any("result", map[string]interface{}{
				"sp1": sp1,
				"err": err}))
		}
	}()
	return _d._base.Update(ctx, req)
}
