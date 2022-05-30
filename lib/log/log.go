package log

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/xiaolaji422/golink/config"
	"github.com/xiaolaji422/golink/lib/file"
)

var (
	log_dir  = "./"
	log_file = "golink.log"
)

func init() {
	// 设置默认文件夹和文件
	defalut_file := config.Conf.Get("log.defalut_file")
	defalut_dir := config.Conf.Get("log.defalut_dir")
	log_file = fmt.Sprintf("%s", defalut_file)
	log_dir = fmt.Sprintf("%s", defalut_dir)
}

// 获取实例
func Instance(filename ...string) *logger {
	var file = log_file
	if len(filename) > 0 {
		file = filename[0]
	}
	return &logger{
		log_dir:  log_dir,
		log_file: file,
	}
}

// 日志文件
type logger struct {
	log_dir  string
	log_file string
	ctx      context.Context
}

// 指定文件
func (l *logger) File(file string) *logger {
	l.log_file = file
	return l
}

// 上下文
func (l *logger) Ctx(ctx context.Context) *logger {
	l.ctx = ctx
	return l
}

// 错误日志
func (l *logger) Error(data ...interface{}) error {
	return l.Write("Error", data...)
}

// 成功日志
func (l *logger) Success(data ...interface{}) error {
	return l.Write("Success", data...)
}

// 普通级别日志
func (l *logger) Info(data ...interface{}) error {
	return l.Write("Info", data...)
}

// 写入日志
func (l *logger) Write(level string, data ...interface{}) error {
	var (
		path = l.log_dir + l.log_file
	)
	fmt.Println(path, "Write")
	if err := file.CreateDir(path, 0666); err != nil {
		return err
	}

	content := fmt.Sprintf("【%s】", level)
	var data1 = make([]interface{}, 0)
	data1 = append(data1, content)
	data1 = append(data1, data...)
	// 组织时间
	// if len(data) > 0 {
	// 	for _, v := range data {
	// 		content += " | " + str.String(v)
	// 	}
	// }

	// 写入文件
	f, err := os.OpenFile("log.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		return err
	}
	defer func() {
		f.Close()
	}()

	multiWriter := io.MultiWriter(os.Stdout, f)
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)

	log.SetOutput(multiWriter)
	log.Println(data1...)
	return nil
}
