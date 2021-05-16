package test1

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
)

const (
	url string = "https://www.xiaohongshu.com/user/profile/5ce4fcb7000000001003e169"
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

func HttpTest() {
	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatalln(err)
	}
	client := http.Client{Jar: jar}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.AddCookie(&http.Cookie{Name: "xhsTrackerId", Value: "153c9c7a-2075-4c77-cc05-3b8d27f0f1c5"})
	req.AddCookie(&http.Cookie{Name: "extra_exp_ids", Value: "gif_clt1,ques_exp2"})
	req.AddCookie(&http.Cookie{Name: "xhsuid", Value: "HuAYDHWVo43IkzXY"})
	req.AddCookie(&http.Cookie{Name: "timestamp2", Value: "20210330c92c402414e571bfcf126570"})
	req.AddCookie(&http.Cookie{Name: "timestamp2.sig", Value: "xK2LAExD_i_eJAZliDx0afxbhU-Wmx_bjIsqCo--G8M"})

	rsp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	bs, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	body := string(bs)
	log.Println(body)
}
