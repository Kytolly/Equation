package launch

import(
	"fmt"
	"net/http"
    "log"
    "back/pck/websocket"
)
func myweb(w http.ResponseWriter,r *http.Request){
    
}

func Server(){
	http.HandleFunc("/", myweb)
    fmt.Println("服务器即将开启，访问地址 http://localhost:8080")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        fmt.Println("服务器开启错误: ", err)
    }
}