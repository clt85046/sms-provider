package http

import (
	"net/http"
	"github.com/open-falcon/sms-provider/config"
	"github.com/toolkits/web/param"
	"fmt"
    "github.com/toolkits/net/httplib"
    "time"
)

func configProcRoutes() {

	http.HandleFunc("/sms", func(w http.ResponseWriter, r *http.Request) {
		cfg := config.Config()
		token := param.String(r, "token", "")
		if cfg.Http.Token != token {
			http.Error(w, "no privilege", http.StatusForbidden)
			return
		}

		tos := param.MustString(r, "tos")
		content := param.MustString(r, "content")
		url := cfg.Sms.Addr
		s := httplib.Post(url).SetTimeout(5*time.Second, 30*time.Second)
    	s.Param("PhoneNumbers", tos)
    	s.Param("Content", content)
    	s.Param("ApiName",cfg.Sms.ApiName)
    	s.Param("Token",cfg.Sms.Token)
    	resp, err := s.String()
    	if err != nil {
    		http.Error(w, err.Error(), http.StatusInternalServerError)
        	fmt.Println("send sms fail, tos:%s, cotent:%s, error:%v", tos, content, err)
    	}else {
			http.Error(w, "success", http.StatusOK)
			fmt.Println("send sms:%v, resp:%v, url:%s", content, resp, url)
		}

	})

}
