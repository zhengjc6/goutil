package logxfc

import (
	"bufio"
	"os"

	"github.com/zeromicro/go-zero/core/logx"
)

type MultiWriter struct {
	writer        logx.Writer
	consoleWriter logx.Writer
}

func NewMultiWriter(writer logx.Writer) (logx.Writer, error) {
	return &MultiWriter{
		writer:        writer,
		consoleWriter: logx.NewWriter(bufio.NewWriter(os.Stdout)),
	}, nil
}

func (w *MultiWriter) Alert(v interface{}) {
	w.consoleWriter.Alert(v)
	w.writer.Alert(v)
}

func (w *MultiWriter) Close() error {
	w.consoleWriter.Close()
	return w.writer.Close()
}

func (w *MultiWriter) Debug(v interface{}, fields ...logx.LogField) {
	w.consoleWriter.Debug(v, fields...)
	w.writer.Debug(v, fields...)
}

func (w *MultiWriter) Error(v interface{}, fields ...logx.LogField) {
	w.consoleWriter.Error(v, fields...)
	w.writer.Error(v, fields...)
}

func (w *MultiWriter) Info(v interface{}, fields ...logx.LogField) {
	w.consoleWriter.Info(v, fields...)
	w.writer.Info(v, fields...)
}

func (w *MultiWriter) Severe(v interface{}) {
	w.consoleWriter.Severe(v)
	w.writer.Severe(v)
}

func (w *MultiWriter) Slow(v interface{}, fields ...logx.LogField) {
	w.consoleWriter.Slow(v, fields...)
	w.writer.Slow(v, fields...)
}

func (w *MultiWriter) Stack(v interface{}) {
	w.consoleWriter.Stack(v)
	w.writer.Stack(v)
}

func (w *MultiWriter) Stat(v interface{}, fields ...logx.LogField) {
	w.consoleWriter.Stat(v, fields...)
	w.writer.Stat(v, fields...)
}
