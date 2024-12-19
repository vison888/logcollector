package log

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/vison888/go-vkit/logger"
	"github.com/vison888/logcollector/app"
	pb "github.com/vison888/logcollector/proto/log_collector"
)

var msgCh chan *pb.Log

func init() {
	msgCh = make(chan *pb.Log, 1024)
	go mainloop()
}

func log(log *pb.Log) error {
	fileName := log.MeetingId + ".log"
	filePath := app.Cfg.Log.Dir + fileName
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		logger.Error("can not create file", err)
		return err
	}
	defer file.Close()

	cur := time.UnixMilli(log.Timestamp)
	timeStr := cur.Format("[2006-01-02 15:04:05.000000]")
	bytes := []byte(fmt.Sprintf("%s MeetingId:%s UserId:%s Module:%s ", timeStr, log.MeetingId, log.UserId, log.Module))
	bytes = append(bytes, log.Info...)
	bytes = append(bytes, '\n')
	if _, err := file.Write(bytes); err != nil {
		logger.Error("write file fail", err)
		return err
	}
	return nil
}

func mainloop() {
	for {
		select {
		case msg := <-msgCh:
			log(msg)
		}
	}
}

func InitLog() error {
	// 判断文件夹是否存在
	path := app.Cfg.Log.Dir
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.Mkdir(path, os.ModePerm); err != nil {
			return err
		} else {
			return nil
		}
	} else {
		return err
	}
}

func Log(log *pb.Log) error {
	select {
	case msgCh <- log:
	default:
		logByte, _ := json.Marshal(log)
		logger.Errorf("queue block %s", string(logByte))
	}
	return nil
}

func Logs(logs []*pb.Log) error {
	for _, v := range logs {
		err := Log(v)
		if err != nil {
			return err
		}
	}
	return nil
}
