package common

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"
	"net/http"
	"strconv"
)

func PrometheusBoot(port int)  {
   http.Handle("/metrics",promhttp.Handler())
   //start web service
   go func() {
	   if err := http.ListenAndServe("0.0.0.0:"+strconv.Itoa(port), nil);err!=nil{
	   	   log.Fatal("start failed please try can later")
	   }
	   log.Info("prometheus start ,the port is:"+strconv.Itoa(port)) //int  to string
   }()
}
//取件88558554