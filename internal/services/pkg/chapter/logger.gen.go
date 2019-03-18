package chapter

// DO NOT EDIT!
// This code is generated with http://github.com/hexdigest/gowrap tool
// using ../../../../tools/templates/logger template

//go:generate gowrap gen -p go.zenithar.org/spotigraph/internal/services -i Chapter -t ../../../../tools/templates/logger -o chapter.logger.go

import (
	"context"

	"go.uber.org/zap"
	"go.zenithar.org/pkg/log"
	"go.zenithar.org/spotigraph/internal/services"
	"go.zenithar.org/spotigraph/pkg/protocol/v1/spotigraph"
)

// ChapterLogger implements services.Chapter that is instrumented with logger
type ChapterLogger struct {
	_log  log.Logger
	_base services.Chapter
}

// NewChapterLogger instruments an implementation of the services.Chapter with simple logging
func NewChapterLogger(base services.Chapter, logger log.LoggerFactory) ChapterLogger {
	return ChapterLogger{
		_base: base,
		_log:  logger.With(zap.String("decorator", "logger")),
	}
}

// Create implements services.Chapter
func (_d ChapterLogger) Create(ctx context.Context, req *spotigraph.ChapterCreateReq) (sp1 *spotigraph.SingleChapterRes, err error) {
	_d._log.Debug("ChapterLogger: calling Create", zap.String("method", "Create"))
	defer func() {
		if err != nil {
			_d._log.Error("ChapterLogger: method Create returned an error", zap.String("method", "Create"), zap.Error(err), zap.Any("result", map[string]interface{}{
				"sp1": sp1,
				"err": err}))
		} else {
			_d._log.Error("ChapterLogger: method Create finished", zap.String("method", "Create"), zap.Any("result", map[string]interface{}{
				"sp1": sp1,
				"err": err}))
		}
	}()
	return _d._base.Create(ctx, req)
}

// Delete implements services.Chapter
func (_d ChapterLogger) Delete(ctx context.Context, req *spotigraph.ChapterGetReq) (ep1 *spotigraph.EmptyRes, err error) {
	_d._log.Debug("ChapterLogger: calling Delete", zap.String("method", "Delete"))
	defer func() {
		if err != nil {
			_d._log.Error("ChapterLogger: method Delete returned an error", zap.String("method", "Delete"), zap.Error(err), zap.Any("result", map[string]interface{}{
				"ep1": ep1,
				"err": err}))
		} else {
			_d._log.Error("ChapterLogger: method Delete finished", zap.String("method", "Delete"), zap.Any("result", map[string]interface{}{
				"ep1": ep1,
				"err": err}))
		}
	}()
	return _d._base.Delete(ctx, req)
}

// Get implements services.Chapter
func (_d ChapterLogger) Get(ctx context.Context, req *spotigraph.ChapterGetReq) (sp1 *spotigraph.SingleChapterRes, err error) {
	_d._log.Debug("ChapterLogger: calling Get", zap.String("method", "Get"))
	defer func() {
		if err != nil {
			_d._log.Error("ChapterLogger: method Get returned an error", zap.String("method", "Get"), zap.Error(err), zap.Any("result", map[string]interface{}{
				"sp1": sp1,
				"err": err}))
		} else {
			_d._log.Error("ChapterLogger: method Get finished", zap.String("method", "Get"), zap.Any("result", map[string]interface{}{
				"sp1": sp1,
				"err": err}))
		}
	}()
	return _d._base.Get(ctx, req)
}

// Update implements services.Chapter
func (_d ChapterLogger) Update(ctx context.Context, req *spotigraph.ChapterUpdateReq) (sp1 *spotigraph.SingleChapterRes, err error) {
	_d._log.Debug("ChapterLogger: calling Update", zap.String("method", "Update"))
	defer func() {
		if err != nil {
			_d._log.Error("ChapterLogger: method Update returned an error", zap.String("method", "Update"), zap.Error(err), zap.Any("result", map[string]interface{}{
				"sp1": sp1,
				"err": err}))
		} else {
			_d._log.Error("ChapterLogger: method Update finished", zap.String("method", "Update"), zap.Any("result", map[string]interface{}{
				"sp1": sp1,
				"err": err}))
		}
	}()
	return _d._base.Update(ctx, req)
}
