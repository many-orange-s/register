package logger

import (
	"bluebull/model"
	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"
)

var lg *zap.Logger

func Init(logconf *model.LogConfig) (err error) {
	//日志的两个参数获取
	encoder := getencoder()
	writersyncer := getsyncer(
		logconf.Filename,
		logconf.MaxSize,
		logconf.MaxAge,
		logconf.Maxbackups,
	)
	//获取配置中的level
	var l = new(zapcore.Level)
	err = l.UnmarshalText([]byte(viper.GetString("log.level")))
	if err != nil {

		return
	}
	//三个变量结合
	core := zapcore.NewCore(encoder, writersyncer, l)
	lg = zap.New(core, zap.AddCaller())
	//替换全局变量
	zap.ReplaceGlobals(lg)
	return
}

func getencoder() zapcore.Encoder {
	//获取设置日志的基础信息
	encoder := zap.NewProductionEncoderConfig()
	//改变日志的时间
	encoder.EncodeTime = zapcore.ISO8601TimeEncoder
	//改变日志的大写字母
	encoder.EncodeLevel = zapcore.CapitalLevelEncoder
	//josn 信息返回
	return zapcore.NewJSONEncoder(encoder)
}

func getsyncer(filename string, maxSize, maxAge, maxBackup int) zapcore.WriteSyncer {
	//导入lumberjack
	lumberjack := &lumberjack.Logger{
		Filename:   "./" + filename,
		MaxSize:    maxSize,
		MaxAge:     maxAge,
		MaxBackups: maxBackup,
	}
	return zapcore.AddSync(lumberjack)
}

func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		// Process request
		c.Next()

		cost := time.Since(start)
		zap.L().Info(path,
			zap.Int("statue", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("ip", c.ClientIP()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)

	}
}

func GinRecover(stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					zap.L().Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}

				if stack {
					zap.L().Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					zap.L().Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
