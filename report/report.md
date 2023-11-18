# 基于golang+Web（html，css,javascript,Qjuery）的多平台线性非齐次方程组低阶计算器 
>电子科技大学go语言与区块链技术应用(桂勋)半期设计
>
> 姓名：谢卿云
>
>学号：2022010910017
>
>专业：计算机科学与工程（"互联网+"复合型精英人才双学位培养计划）
> 
>声明：
>
>本项目遵循署名-非商业性使用-相同方式共享 4.0 国际 (CC BY-NC-SA 4.0) 协议，在`https://github.com/Kytolly/Equation`上已开源，相应介绍也同步于个人博客`kytolly.github.io`，项目有所疏漏在所难免，欢迎探讨交流和改进,但由于项目具有考核性质，本人也反对任何形式的无脑抄袭，若因为抄袭本项目或者某个部分而导致考试不通过甚至更严重的后果，本人概不负责，并声明本人并没有任何抄袭的代码的行为，并且本人提交的作业也并非开源的版本


## 一、项目简介与设计思想
本项目实现了一个基于高斯消元法求解线性非齐次方程组的交互式计算器，前端使用`HTML+CSS+Javascript`，后端使用`Golang`，前后端间通过传递`json`文件进行通信和交互。

利用`net/http`包，在前端打开位于`../front`文件夹中的前端文件，监听端口`:8080`，注册`dataHandler`作为访问`/`的`http`服务请求的处理器；

在前端设置一个下拉列表`<input type="number">`，用于限制用户的输入规模，另外设置两个按钮`button`具备`onclick`的`HTML DOM`事件属性，用于触发用户的输入方程组的信息与提交方程组的信息；

服务器将检查用户的合法输入，并将输入的数据转化为`json`格式，传入结构体变量`Mat`,后端利用`back/pck/cal`包中函数`func Gauss(Mat Equation)(x Vec,err error)`和`func Output(ans Vec,ans_exit error)(x string)`得到传回前端的字符串`ans_str`,在后端转化为`json`格式`ans_json`后传回前端，在前端相应的位置进行打印，这样便完成该计算机的初步模型设计；


## 二、系统架构与功能描述

本项目的文件关系可以用如下树状图表示：
```txt
./
├── back             后端
│   ├── go.mod       go.mod   
│   ├── main.go      运行封装的函数（Server,Debug）
│   └── pck          用到的包
│       ├── cal      计算模块和数据结构定义模块
│       │   └── cal.go
│       ├── launch
│       │   ├── Debug.go  测试cal.go
│       │   └── launch.go  启动服务程序
│       └── websocket  实际上并未用到
│           ├── client.go
│           ├── dial.go
│           ├── dial_test.go
│           ├── exampledial_test.go
│           ├── examplehandler_test.go
│           ├── hybi.go
│           ├── hybi_test.go
│           ├── server.go
│           ├── websocket.go
│           └── websocket_test.go
├── front
│   └── index.html  前端传输文件
├── README.md       README
├── release         发行版，适合多操作系统
│   ├── build
│   │   └── run
│   │       ├── Analysis-00.toc
│   │       ├── base_library.zip
│   │       ├── COLLECT-00.toc
│   │       ├── EXE-00.toc
│   │       ├── localpycs
│   │       │   ├── pyimod01_archive.pyc
│   │       │   ├── pyimod02_importers.pyc
│   │       │   ├── pyimod03_ctypes.pyc
│   │       │   └── struct.pyc
│   │       ├── PKG-00.toc
│   │       ├── PYZ-00.pyz
│   │       ├── PYZ-00.toc
│   │       ├── run
│   │       ├── run.pkg
│   │       ├── warn-run.txt
│   │       └── xref-run.html
│   ├── dist
│   │   └── run
│   │       ├── _internal
│   │       │   ├── base_library.zip
│   │       │   ├── libcrypto.so.1.1
│   │       │   ├── lib-dynload
│   │       │   │   ├── array.cpython-39-x86_64-linux-gnu.so
│   │       │   │   ├── binascii.cpython-39-x86_64-linux-gnu.so
│   │       │   │   ├── _bisect.cpython-39-x86_64-linux-gnu.so
│   │       │   │   ├── _blake2.cpython-39-x86_64-linux-gnu.so
│   │       │   │   ├── _bz2.cpython-39-x86_64-linux-gnu.so
│   │       │   │   ├── _codecs_cn.cpython-39-x86_64-linux-gnu.so
│   │       │   │   ├── _codecs_hk.cpython-39-x86_64-linux-gnu.so
│   │       │   │   ├── _codecs_iso2022.cpython-39-x86_64-linux-gnu.so
│   │       │   │   ├── _codecs_jp.cpython-39-x86_64-linux-gnu.so
│   │       │   │   ├── _codecs_kr.cpython-39-x86_64-linux-gnu.so
│   │       │   │   ├── _codecs_tw.cpython-39-x86_64-linux-gnu.so
│   │       │   │   ├── _contextvars.cpython-39-x86_64-linux-gnu.so
│   │       │   │   ├── _csv.cpython-39-x86_64-linux-gnu.so
│   │       │   │   ├── _datetime.cpython-39-x86_64-linux-gnu.so
│   │       │   │   ├── _decimal.cpython-39-x86_64-linux-gnu.so
│   │       │   │   ├── grp.cpython-39-x86_64-linux-gnu.so
│   │       │   │   ├── _hashlib.cpython-39-x86_64-linux-gnu.so
│   │       │   │   ├── _heapq.cpython-39-x86_64-linux-gnu.so
│   │       │   │   ├── _lzma.cpython-39-x86_64-linux-gnu.so
│   │       │   │   ├── math.cpython-39-x86_64-linux-gnu.so
│   │       │   │   ├── _md5.cpython-39-x86_64-linux-gnu.so
│   │       │   │   ├── _multibytecodec.cpython-39-x86_64-linux-gnu.so
│   │       │   │   ├── _pickle.cpython-39-x86_64-linux-gnu.so
│   │       │   │   ├── _posixsubprocess.cpython-39-x86_64-linux-gnu.so
│   │       │   │   ├── _random.cpython-39-x86_64-linux-gnu.so
│   │       │   │   ├── resource.cpython-39-x86_64-linux-gnu.so
│   │       │   │   ├── select.cpython-39-x86_64-linux-gnu.so
│   │       │   │   ├── _sha1.cpython-39-x86_64-linux-gnu.so
│   │       │   │   ├── _sha256.cpython-39-x86_64-linux-gnu.so
│   │       │   │   ├── _sha3.cpython-39-x86_64-linux-gnu.so
│   │       │   │   ├── _sha512.cpython-39-x86_64-linux-gnu.so
│   │       │   │   ├── _socket.cpython-39-x86_64-linux-gnu.so
│   │       │   │   ├── _statistics.cpython-39-x86_64-linux-gnu.so
│   │       │   │   ├── _struct.cpython-39-x86_64-linux-gnu.so
│   │       │   │   ├── unicodedata.cpython-39-x86_64-linux-gnu.so
│   │       │   │   └── zlib.cpython-39-x86_64-linux-gnu.so
│   │       │   ├── liblzma.so.5
│   │       │   ├── libpython3.9.so.1.0
│   │       │   └── libz.so.1
│   │       └── run
│   ├── GaussCalculator_ubuntu20.04_version1.0
│   ├── run.bat
│   ├── run.py
│   └── run.spec
└── report
    └── report.md

15 directories, 79 files

```


目前允许的功能有：

* 输入的方程阶数不超过6
* 输入的系数值不超过10位（包括小数点前后的所有数字）
* 能对方程解以行向量的形式返回；
* 对于有无数个解和无解的特殊情况能单独返回相应的提示；
* 能对通信间可能出现的错误进行了返回，并在控制台程序提供了详尽的日志查询


## 三、关键技术与算法论述
### 自定义数据结构
#### cal.Equation
由一个数据域`Value_num int`和两个指针域`Coefficient_Matrix [][]float64`和`Expansion []float64`构成，分别使用`json:"n"  json:"matrix"  json:"expansion"`来解析；

#### cal.Vec
由一个数据域`Len int`和一个指针域`Value []float64`构成，由于不在前端出现，因此不提供json的类型

### 高斯消元法模块
实际上采用的是高斯—若尔当消元法：

消元法理论的核心主要如下：

-   两方程互换，解不变；

-   一方程乘以非零数 $k$，解不变；

-   一方程乘以数 $k$ 加上另一方程，解不变。

算法的流程可以简单描述如下：

找到第i次迭代的列主元（这里我们找到绝对值最大的通过交换到第i行，**避免因为精度问题而造成的误差**）$\to$ 消去其余列主元 $\to$ 迭代结束；

自上往下回带的同时化为对角矩阵$\to$化为对角矩阵后得出方程的解向量（对比列主元和拓展向量的分量，可能存在着无解或者无穷解）；

算法复杂度为$O(N^3)$,并且存在精度误差的缺点；

可能的解有三种情况:
* 有唯一解，此时返回`x Vec`记录方程阶数和解向量，以及`err=nil error`;
* 有无数解，此时返回`x Vec`记录方程阶数和默认的零向量，以及`err=Errors.New(”方程有无数解“) error`;
* 无解，此时返回`x Vec`记录方程阶数和零向量，以及`err=Errors.New(”方程无解“) error`;

### Output模块
根据Gauss消元法算法模块得到的`Vec error`,利用`fmt.Sprintf()`函数转化为字符串，方便利用并解析为`json`传回前端;

### dataHandler处理器模块
`func dataHandler(w http.ResponseWriter, r *http.Request)`流程如下：

1. 判断`http`服务的请求方法`r.Method`，如果为`GET`,打开前端页面`../front/index.html`，或者打开404页面；
如果为`POST`,进入下一步；默认分支为控制台提示无效请求并记录日志；
2. 控制台提示获得`POST`请求，读取并处理`r.body`，若是`closed`的状态，控制台提示无效读取并记录日志；
3. 解析`body`到`&Mat`,若解析成功，控制台提示成功解析；
4. 处理得到的`json`，得到`ans_str string`;
5. 将`ans_str`解析为`ans_json`传回前端，若传回失败，记录日志；

### $.ajax前端模块
准备一个JSON对象，其中包含n（矩阵的尺寸），matrix（系数矩阵）和expansion（扩展列）。然后，使用jQuery的Ajax方法向本地运行的服务器（地址为http://localhost:8080）发送一个POST请求。请求成功时，解析返回的JSON数据并更新页面；请求失败时，显示错误信息。

## 四、运行效果与运行环境
下面是本项目的编译环境与依赖包
>OS: Ubuntu 22.04.3 LTS
>
> IDE: Microsoft Visual Studio Code
>
> Go version: 1.18
>
> Dependency: 
> * "fmt"
> * "net/http"
> * "log"
>* "io/ioutil"
>*  "encoding/json"
> * "math"
> * "errors"


下面是程序运行的运行环境
>对于linux平台，在/release下打开终端，输入`./GaussCaculator_ubuntu20.04_version1.0`,进入`http://localhost:8080`即可启动运行程序;

>对于windows11,点击`run.bat`即可启动程序；

>通用办法：在/release/dist目录下寻找适合操作系统的可执行文件，这是基于pyinstaller生成，运行/release下的`run.py`

经过测试，即便挂着代理服务应该也能访问，但防火墙需要关掉；

以下是项目的运行效果截图：

## 五、总体评价与未来展望
总体评价：

本项目体量小，代码结构具有代表性，可学习性高，对网络编程、计算机网络与协议、golang基础语法与简单数据结构与算法的综合性高，适合对于利用golang开发网站的初学者，有利于对大规模项目的文件管理与基本代码编写风格（如面向对象、函数式编程）与调试方法（如日志追踪）有进一步的理解；

但是对课业紧张的大二计算机专业生，加上对前端语言与框架、网络通信协议，Linux操作系统的不熟悉，本项目的学习曲线较为陡峭，带来相当大的压力，并且容易让初学者陷入无穷无尽的细节当中，加上本项目并没有体现golang作为轻量化后端语言在高并发编程和高性能的垃圾回收与编译机制上的优势，这一定程度上让项目意义失色不少；


对于项目本身的展望：

参考CASIO计算器上对联立方程的功能实现
* 前后通信上，对于方程有无穷解的情况能否进一步优化，显式表达出来；
* 算法上，组合一个网页计算器，对于前端输入的并不是特定的数而是计算式组成的字符串，能否计算出来并且正常输入到后端，同时进行严格的输入检查；
* 前端上，对于前端页面能否进一步美化，能否实现多个URL跳转；
