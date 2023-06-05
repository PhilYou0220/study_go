package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

//var lg *zap.Logger //全局logger

func Init(){
	//1.encoder 编码器、格式器
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "time" //更改 {"ts":1684843048.9437547}
	encoderCfg.EncodeTime =zapcore.ISO8601TimeEncoder //设置日期显示格式
	encoder :=zapcore.NewJSONEncoder(encoderCfg) //zap.NewProductionEncoderConfig()一些json
	//2、WriteSyncer（写到哪里）
	file,_:=os.OpenFile("./test.log",os.O_CREATE | os.O_APPEND | os.O_WRONLY,0644)
	//传入文件句柄输出到文件里
	fileWS := zapcore.AddSync(file)
	consoleWS := zapcore.AddSync(os.Stdout)
	//利用core生成logger
	core := zapcore.NewCore(encoder,zapcore.NewMultiWriteSyncer(fileWS,consoleWS),zapcore.InfoLevel)

	lg := zap.New(core,zap.AddCaller())

	//替换zap全局的logger
	zap.ReplaceGlobals(lg)
	zap.L().Info("success")
}
