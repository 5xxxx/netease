/*
 *
 * log.go
 * netease
 *
 * Created by lintao on 2020/7/15 1:50 下午
 * Copyright © 2020-2020 LINTAO. All rights reserved.
 *
 */

package netease

type Logger interface {
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Fatal(args ...interface{})
	Infof(format string, args ...interface{})
	Info(args ...interface{})
	Warnf(format string, args ...interface{})
	Debugf(format string, args ...interface{})
	Debug(args ...interface{})
}
