package guild

// DO NOT EDIT!
// This code is generated with http://github.com/hexdigest/gowrap tool
// using ../../../../tools/templates/logger template

//go:generate gowrap gen -p go.zenithar.org/spotigraph/internal/services -i Guild -t ../../../../tools/templates/logger -o service.logger.gen.go

import (
	"context"

	"go.uber.org/zap"
	"go.zenithar.org/pkg/log"
	"go.zenithar.org/spotigraph/internal/services"
	"go.zenithar.org/spotigraph/pkg/protocol/v1/spotigraph"
)

// GuildLogger implements services.Guild that is instrumented with logger
type GuildLogger struct {
	_log  log.Logger
	_base services.Guild
}

// NewGuildLogger instruments an implementation of the services.Guild with simple logging
func NewGuildLogger(base services.Guild, logger log.LoggerFactory) GuildLogger {
	return GuildLogger{
		_base: base,
		_log:  logger.With(zap.String("decorator", "logger")),
	}
}

// Create implements services.Guild
func (_d GuildLogger) Create(ctx context.Context, req *spotigraph.GuildCreateReq) (sp1 *spotigraph.SingleGuildRes, err error) {
	_d._log.Debug("GuildLogger: calling Create", zap.String("method", "Create"))
	defer func() {
		if err != nil {
			_d._log.Error("GuildLogger: method Create returned an error", zap.String("method", "Create"), zap.Error(err), zap.Any("result", map[string]interface{}{
				"sp1": sp1,
				"err": err}))
		} else {
			_d._log.Error("GuildLogger: method Create finished", zap.String("method", "Create"), zap.Any("result", map[string]interface{}{
				"sp1": sp1,
				"err": err}))
		}
	}()
	return _d._base.Create(ctx, req)
}

// Delete implements services.Guild
func (_d GuildLogger) Delete(ctx context.Context, req *spotigraph.GuildGetReq) (ep1 *spotigraph.EmptyRes, err error) {
	_d._log.Debug("GuildLogger: calling Delete", zap.String("method", "Delete"))
	defer func() {
		if err != nil {
			_d._log.Error("GuildLogger: method Delete returned an error", zap.String("method", "Delete"), zap.Error(err), zap.Any("result", map[string]interface{}{
				"ep1": ep1,
				"err": err}))
		} else {
			_d._log.Error("GuildLogger: method Delete finished", zap.String("method", "Delete"), zap.Any("result", map[string]interface{}{
				"ep1": ep1,
				"err": err}))
		}
	}()
	return _d._base.Delete(ctx, req)
}

// Get implements services.Guild
func (_d GuildLogger) Get(ctx context.Context, req *spotigraph.GuildGetReq) (sp1 *spotigraph.SingleGuildRes, err error) {
	_d._log.Debug("GuildLogger: calling Get", zap.String("method", "Get"))
	defer func() {
		if err != nil {
			_d._log.Error("GuildLogger: method Get returned an error", zap.String("method", "Get"), zap.Error(err), zap.Any("result", map[string]interface{}{
				"sp1": sp1,
				"err": err}))
		} else {
			_d._log.Error("GuildLogger: method Get finished", zap.String("method", "Get"), zap.Any("result", map[string]interface{}{
				"sp1": sp1,
				"err": err}))
		}
	}()
	return _d._base.Get(ctx, req)
}

// Update implements services.Guild
func (_d GuildLogger) Update(ctx context.Context, req *spotigraph.GuildUpdateReq) (sp1 *spotigraph.SingleGuildRes, err error) {
	_d._log.Debug("GuildLogger: calling Update", zap.String("method", "Update"))
	defer func() {
		if err != nil {
			_d._log.Error("GuildLogger: method Update returned an error", zap.String("method", "Update"), zap.Error(err), zap.Any("result", map[string]interface{}{
				"sp1": sp1,
				"err": err}))
		} else {
			_d._log.Error("GuildLogger: method Update finished", zap.String("method", "Update"), zap.Any("result", map[string]interface{}{
				"sp1": sp1,
				"err": err}))
		}
	}()
	return _d._base.Update(ctx, req)
}
