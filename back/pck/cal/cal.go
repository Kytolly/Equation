package cal

import(
	"fmt"
)

//定义数据结构
type double float64

type Vec struct {
    len int
    value []double
}

type Equation struct {
    value_num int //方程维数
    Coefficient_Matrix [][]double  //系数矩阵
    Expansion []double
}

//定义方法

func(s Equation) Init(n int){//初始化
    s.value_num=n
    s.Coefficient_Matrix =make([][]double,s.value_num)
    for i:=0;i<s.value_num;i++ {
        s.Coefficient_Matrix[i]= make([]double,s.value_num)
    }
    s.Expansion =make([]double,s.value_num)
}

func(a Vec) Init(n int){
    a.len =n
    a.value =make([]double,a.len)
}


func(s Equation)Print(){//显式打印出来
    fmt.Printf("方程的维数：%v\n",s.value_num)
    fmt.Printf("方程的系数矩阵为：\n")
    for i:=0;i<s.value_num;i++ {
        for j:=0;j<s.value_num;j++{
            fmt.Printf("%14.6v",s.Coefficient_Matrix[i][j]);
        }
        fmt.Println()
    }
}
 func(b Vec)Print(){
    fmt.Printf("向量的维数：%v\n",b.len)
    fmt.Printf("向量的分量为\n")
    for i:=0;i<b.len;i++{
        fmt.Printf("%14.6v",b.value[i])
    }
    fmt.Println()
}

//利用高斯消元法求解方程，定义无解或者无穷解为错误类型
//后期可以优化无穷解的表示形式并且显式地传入前端
//errors.New("error info")传入错误的描述信息
func Gauss(s Equation,b Vec)(x Vec,err error){
    
}


// 随便手搓一个理解一下
// func cal_2(n int,A [][]int,Y []int) []float64{//处理二阶的简单版本
//     X:= make([]float64,n,2*n);
//     var un float64;
//     un=float64(A[0][0]*A[1][1]-A[1][0]*A[0][1]);
//     X[0]=(float64(Y[0]*A[1][1]-Y[1]*A[0][1]))/un;
//     X[1]=(float64(Y[1]*A[0][0]-Y[0]*A[1][0]))/un;
//     return X;
// }