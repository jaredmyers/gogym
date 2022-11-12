//log handling

package logging

import (
		"log"
		"os"
)

var (
		logPath string = "logging.log"
)

type aggregateLogger struct {
		infoLogger *log.Logger
		warnLogger *log.Logger
		errorLogger *log.Logger
}

func (l *aggregateLogger) info(v ...interface{}){
		l.infoLogger.Println(v...)

}

func (l *aggregateLogger) warn(v ...interface{}){
		l.warnLogger.Println(v...)

}

func (l *aggregateLogger) error(v ...interface{}){
		l.errorLogger.Println(v...)

}

func main() {

		flags := log.LstdFlags | log.Lshortfile

		al := aggregateLogger{
				infoLogger: log.New(os.Stdout, "INFO:", flags),
				warnLogger: log.New(os.Stdout, "WARN:", flags),
				errorLogger: log.New(os.Stdout, "ERR:", flags),
		}

		al.info("info log") 
		al.warn("warn log") 
		al.error("error log") 

}


