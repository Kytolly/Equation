package launch

import(
    "back/pck/cal"
	"fmt"
	"net/http"
    "log"
    "back/pck/websocket"
)

//处理`/Equation`页面发来的数据
func dataHandler(sock *websocket.Conn){
    var(
        Mat cal.Equation //方程的变量个数和系数矩阵
        err error //检查方程是否传送到后端
        answer cal.Vec
    )
    for{
        err = websocket.Message.Receive(sock, &expression)
        if err !=nil{
            fmt.Printf(" 数据通信失败：%s.\n", err.Error())
			//panic(err.Error())
            break
        }
        answer=cal.Guass(Mat)
        
    }
}


//8080端口启动服务器
func Server(){
    http.Handle("/", http.FileServer(http.Dir("..")))
    //http.FileServer()方法返回的是fileHandler实例，
    //fileHandler结构体实现了Handler接口中的ServerHTTP()方法。
	//http.Dir()方法会返回http.Dir类型用于将字符串路径转换为文件系统
    http.Handle("/Equation", websocket.Handler(dataHandler))
    fmt.Println("服务器即将开启，访问地址 http://localhost:8080")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        //打印错误并且退出应用程序
        log.Fatal("http.ListenAndServe() error: ", err.Error())
    }
}



// func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
//     switch req.URL.Path {
//     case "/list":
//     for item, price := range db {
//     fmt.Fprintf(w, "%s: %s\n", item, price)
//     }
//     case "/price":
//     item := req.URL.Query().Get("item")
//     price, ok := db[item]
//     if !ok {
//     w.WriteHeader(http.StatusNotFound) // 404
//     fmt.Fprintf(w, "no such item: %q\n", item)
//     return
//     }
//     fmt.Fprintf(w, "%s\n", price)
//     default:
//     w.WriteHeader(http.StatusNotFound) // 404
//     fmt.Fprintf(w, "no such page: %s\n", req.URL)
//     }
// }
