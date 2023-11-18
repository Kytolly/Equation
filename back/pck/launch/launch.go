package launch

import(
    "back/pck/cal"
	"fmt"
	"net/http"
    "log"
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
            if r.URL.Path == "/" {//获取前端页面front
                http.ServeFile(w, r, "../front/index.html")
            } else {
                http.NotFound(w, r)
            }
            fmt.Println("Received a GET request")

        case "POST":
            fmt.Println("Received a POST request")
            
            body, err := ioutil.ReadAll(r.Body)//读取和处理HTTP请求中的body部分
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
            ans,ans_exit= cal.Gauss(Mat)//后端处理数据得到解
            ans_str=cal.Output(ans,ans_exit)
            fmt.Println(ans_str)

            ans_json, err = json.Marshal(ans_str)//处理解为json格式
            if err != nil {
                http.Error(w, "Error encoding result to JSON", http.StatusInternalServerError)
                return
            }

            w.Header().Set("Content-Type", "application/json")
            w.Write(ans_json)//发送回前端

        default:
            http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
    }
}

//8080端口启动服务器
func Server(){
    fs := http.FileServer(http.Dir("../front"))
	http.Handle("/front/", http.StripPrefix("/front", fs))
    http.HandleFunc("/", dataHandler)
    fmt.Println("The service is going to launch: http://localhost:8080")
	err:=http.ListenAndServe(":8080", nil)
    if err != nil {
        //打印错误并且退出应用程序
        log.Fatal("http.ListenAndServe() error: ", err.Error())
    }
}

