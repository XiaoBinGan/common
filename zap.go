package common

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)




var (
	logger *zap.SugaredLogger
)

func init()  {
	//log file name
	fileName :="micro.log"
	writeSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename: fileName, //file name
		MaxSize:  521,      //the file max size *MB
		//MaxAge:     0,		//the destroy time
		MaxBackups: 0,    //the max back up
		LocalTime:  true, //start local time
		Compress:   true, //is zip
	})
	//encode
	encoder := zap.NewProductionEncoderConfig()
	//time format
	encoder.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(
		//encoder
		zapcore.NewJSONEncoder(encoder),
		writeSyncer,
		zap.NewAtomicLevelAt(zap.DebugLevel))
	log := zap.New(
		core,
		zap.AddCaller(),
		zap.AddCallerSkip(1))//有时我们稍微封装了一下记录日志的方法，但是我们希望输出的文件名和行号是调用封装函数的位置。这时可以使用zap.AddCallerSkip(skip int)向上跳 1 层：
	logger= log.Sugar()
}
func Debug(args ...interface{})  {
	logger.Debug(args)
}
func Debugf(template string,args ...interface{}) {
	logger.Debugf(template, args)
}

func Info(args ...interface{})  {
     logger.Info(args...)
}
func Infof(template string,arg ...interface{})  {
	logger.Infof(template,arg...)
}
func Warn(args ...interface{})  {
	logger.Warn(args...)
}
func Warnf(template string,args ...interface{})  {
	logger.Warnf(template,args...)
}
func Error(args ...interface{})  {
	logger.Error(args...)
}
func Errorf(template string,args ...interface{})  {
	logger.Errorf(template,args)
}
func DPanic(args ...interface{})  {
	logger.DPanic(args...)
}
func DPanicf(template string,args ...interface{})  {
	logger.DPanicf(template,args...)
}
func Fatal(args ...interface{})  {
	logger.Fatal(args...)
}
func FatalF(tempalte string,args ...interface{})  {
	logger.Fatalf(tempalte,args...)
}