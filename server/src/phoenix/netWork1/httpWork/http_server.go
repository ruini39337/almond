package httpWork

import (
	"fmt"
	"github.com/siddontang/go-log/log"
	"net/http"
	"strings"
)

func sayHelloName (w http.ResponseWriter,r * http.Request){
	r.ParseForm() //解析参数，默认不会解析
	fmt.Println(r.Form)
	fmt.Println("path:  ",r.URL.Path)
	fmt.Println("scheme:  ",r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k,v := range r.Form {
		fmt.Println("key:  ",k)
		fmt.Println("val:  ",strings.Join(v," "))
	}
	fmt.Fprintf(w,"hello client!")
}

func Run (){
	http.HandleFunc("/",sayHelloName)
	err := http.ListenAndServe(":9090",nil)
	if err != nil {
		log.Fatal("ListenAndServer: ",err)
	}
	http.Serve()
}

//type srv struct {
//
//}


//
//func (c * srv)Serve(l net.Listener) error {
//	defer l.Close()
//	var tempDelay time.Duration //how long to sleep on accept failure
//
//	for {
//		rw, e := l.Accept()
//		if e != nil {
//			if ne , ok := e.(net.Error);ok && ne.Temporary() {
//					if tempDelay == 0 {
//						tempDelay = 5 *time.Millisecond
//						fmt.Println("tempDelay==0 ",tempDelay)
//					} else {
//						tempDelay *= 2
//					}
//					if max := 1 * time.Second; tempDelay > max {
//						tempDelay = max
//					}
//					log.Printf("http:Accept error:%v retrying in %v",e,tempDelay)
//					time.Sleep(tempDelay)
//					continue
//			}
//			return e
//		}
//		tempDelay = 0
//		//c,err := c.Serve(rw)
//		if err != nil {
//			continue
//		}
//		go
//	}
//
//}