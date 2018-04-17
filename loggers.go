package gautils

import (
	"log"
	"os"
	"io"
)

//logger
var(
	Info	*log.Logger
	Warning *log.Logger
	Error	*log.Logger
	Debug	*log.Logger
)

func init(){
	Info = log.New(os.Stdout,"[Info] ",log.Ldate|log.Ltime)
	Warning = log.New(os.Stdout,"[Warning] ", log.Ldate|log.Ltime)
	Error = log.New(os.Stdout,"[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile)
	Debug = log.New(os.Stdout,"[Debug]",log.Ldate|log.Ltime|log.Lshortfile)
}

func InitError(errFile *os.File){
	if errFile == nil{
		return
	}else{
		Error = log.New(io.MultiWriter(os.Stdout,errFile),"[ERROR] ",log.Ldate|log.Ltime|log.Lshortfile)
	}
}


