package launch

import(
    "back/pck/cal"
	"fmt"
	"net/http"
    "log"
    //"back/pck/websocket"
    "io/ioutil"
    "encoding/json"
)

func dataHandler(w http.ResponseWriter, r *http.Request) {
    var (
        Mat cal.Equation
        ans cal.Vec
        ans_exit error
        ans_str string
        ans_json []byte
    )

    switch r.Method {
        case "GET":
            if r.URL.Path == "/" {
                http.ServeFile(w, r, "../front/index.html")
            } else {
                http.NotFound(w, r)
            }
            fmt.Println("Received a GET request")

        case "POST":
            fmt.Println("Received a POST request")
            
            body, err := ioutil.ReadAll(r.Body)
            if err != nil {
                fmt.Println("Error reading request body:", err)
                http.Error(w, "Error reading request body", http.StatusInternalServerError)
                return
            }
            fmt.Println(string(body))

            err = json.Unmarshal(body, &Mat)
            if err != nil {
                http.Error(w, "Error parsing JSON data", http.StatusBadRequest)
                return
            }
            fmt.Println("Successfully parsed JSON data")

            //Mat.Init(Mat.Value_num)
            Mat.Print()
            ans,ans_exit= cal.Gauss(Mat)
            ans_str=cal.Output(ans,ans_exit)
            fmt.Println(ans_str)

            ans_json, err = json.Marshal(ans_str)
            if err != nil {
                http.Error(w, "Error encoding result to JSON", http.StatusInternalServerError)
                return
            }

            w.Header().Set("Content-Type", "application/json")
            w.Write(ans_json)

        default:
            http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
    }
}

// func dataHandler(w http.ResponseWriter, r *http.Request) {
//     var (
//         Mat cal.Equation
//         ans cal.Vec
//         ans_exit error
//         //ans_str cal.Answer
//         ans_str string
//         ans_json []byte
//     )
    
//     if r.URL.Path == "/" {
//         http.ServeFile(w, r, "../front/index.html")
//     } else {
//         http.NotFound(w, r)
//     }
//     fmt.Println("Received a request")


// 	if r.Method != "POST" {
// 		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
// 		return
// 	}
//     fmt.Println("Received a POST request")

// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
//         fmt.Println("Error reading request body:", err)
// 		http.Error(w, "Error reading request body", http.StatusInternalServerError)
// 		return
// 	}
//     fmt.Println(string(body))

// 	err = json.Unmarshal(body, &Mat)
// 	if err != nil {
// 		http.Error(w, "Error parsing JSON data", http.StatusBadRequest)
// 		return
// 	}
//     fmt.Println("Successfully parsed JSON data")

//     Mat.Init(Mat.Value_num)
//     Mat.Print()
//     ans,ans_exit= cal.Gauss(Mat)
//     ans_str=cal.Output(ans,ans_exit)
//     fmt.Println(ans_str)

// 	ans_json, err = json.Marshal(ans_str)
// 	if err != nil {
// 		http.Error(w, "Error encoding result to JSON", http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write(ans_json)
// }


//8080端口启动服务器
func Server(){
    fs := http.FileServer(http.Dir("../front"))
	http.Handle("/front/", http.StripPrefix("/front", fs))
    http.HandleFunc("/", dataHandler)
    fmt.Println("服务器即将开启，访问地址 http://localhost:8080")
	err:=http.ListenAndServe(":8080", nil)
    if err != nil {
        //打印错误并且退出应用程序
        log.Fatal("http.ListenAndServe() error: ", err.Error())
    }
}

// func Server(){
//     http.Handle("/", http.FileServer(http.Dir("..")))
//     //http.FileServer()方法返回的是fileHandler实例，
//     //fileHandler结构体实现了Handler接口中的ServerHTTP()方法。
// 	//http.Dir()方法会返回http.Dir类型用于将字符串路径转换为文件系统
//     http.Handle("/Equation", websocket.Handler(dataHandler))
//     fmt.Println("服务器即将开启，访问地址 http://localhost:8080")
//     err := http.ListenAndServe(":8080", nil)
//     if err != nil {
//         //打印错误并且退出应用程序
//         log.Fatal("http.ListenAndServe() error: ", err.Error())
//     }
// }



//处理`/Equation`页面发来的数据
// func dataHandler(sock *websocket.Conn){
//     var(
//         Mat cal.Equation //方程的变量个数和系数矩阵
//         err error //检查方程是否传送到后端
//         ans cal.Vec
//         ans_exit error
//         answer_str string 
//     )
//     for{
//         err = websocket.Message.Receive(sock, &Mat)
//         if err !=nil{
//             fmt.Printf("前端无法返回数据到后端：%s.\n", err.Error())
// 			//panic(err.Error())
//             break
//         }
//         ans,ans_exit=cal.Gauss(Mat)
//         answer_str=cal.Output(ans,ans_exit)
//         fmt.Printf("%s\n",answer_str)
//         err=websocket.Message.Send(sock, answer_str)
//         if err != nil {
// 			fmt.Printf("后端无法返回数据到前端： %s.\n", err.Error())
// 			break
// 		}
//     }
// }



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
