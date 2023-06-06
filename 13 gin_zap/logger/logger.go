package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

//var lg *zap.Logger //全局logger

func Init() error{
	//1.encoder 编码器、格式器
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "time" //更改 {"ts":1684843048.9437547}
	encoderCfg.EncodeTime =zapcore.ISO8601TimeEncoder //设置日期显示格式
	encoder :=zapcore.NewJSONEncoder(encoderCfg) //zap.NewProductionEncoderConfig()一些json
	//2、WriteSyncer（写到哪里）
	file,err:=os.OpenFile("./test.log",os.O_CREATE | os.O_APPEND | os.O_WRONLY,0644)
	if err !=nil{
		return err
	}
	//传入文件句柄输出到文件里
	fileWS := zapcore.AddSync(file)
	consoleWS := zapcore.AddSync(os.Stdout)
	//利用core生成logger
	core := zapcore.NewCore(encoder,zapcore.NewMultiWriteSyncer(fileWS,consoleWS),zapcore.InfoLevel)

	lg := zap.New(core,zap.AddCaller())

	//替换zap全局的logger zap.L() zap.L()时，实际上是获取了lg的快捷方式。
	zap.ReplaceGlobals(lg) //这里之后就可以使用 zap.l()了
	zap.L().Info("success")
	return  nil //没报错
}
