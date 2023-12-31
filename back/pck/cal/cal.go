package cal

import(
	"fmt"
    "math"
    "errors"
)

//定义数据结构
// type Answer struct{
//     ans_str string `json:"result"`
// }
type Vec struct {
    Len int
    Value []float64
}

type Equation struct {
    Value_num int `json:"n"`   //方程维数
    Coefficient_Matrix [][]float64 `json:"matrix"` //系数矩阵
    Expansion []float64  `json:"expansion"`//拓展向量
}

//定义方法

func(s *Equation) Init(n int){//初始化
    s.Value_num=n
    s.Coefficient_Matrix =make([][]float64,s.Value_num)
    for i:=0;i<s.Value_num;i++ {
        s.Coefficient_Matrix[i]= make([]float64,s.Value_num)
        for j:=0;j<s.Value_num;j++{
            s.Coefficient_Matrix[i][j]=0
        }

    }
    s.Expansion =make([]float64,s.Value_num)
    for j:=0;j<s.Value_num;j++{
        s.Expansion[j]=0
    }
}

func(a *Vec) Init(n int){
    a.Len =n
    a.Value =make([]float64,a.Len)
    for j:=0;j<n;j++{
        a.Value[j]=0
    }
}

func(b Vec)Print(){
    fmt.Printf("向量的维数：%v\n",b.Len)
    fmt.Printf("向量的分量为\n")
    for i:=0;i<b.Len;i++{
        fmt.Printf("%14.6v",b.Value[i])
    }
    fmt.Println()
}
func(s Equation)Print(){//显式打印出来
    fmt.Printf("方程的维数：%v\n",s.Value_num)
    fmt.Printf("方程的系数矩阵为：\n")
    for i:=0;i<s.Value_num;i++ {
        for j:=0;j<s.Value_num;j++{
            fmt.Printf("%4.6v",s.Coefficient_Matrix[i][j])
        }
        fmt.Println()
    }
    fmt.Printf("方程的拓展向量为：\n")
    for i:=0;i<s.Value_num;i++ {
        fmt.Printf("%4.6v",s.Expansion[i])
    }
    fmt.Println()
}

func Gauss(s Equation)(x Vec,err error){
    n:=s.Value_num
    x.Init(n);

    for i:=0;i<n;i++{//交换行列
        head:=i
        for k:=i+1;k<n;k++{
            if math.Abs(s.Coefficient_Matrix[k][i]) > math.Abs(s.Coefficient_Matrix[head][i]) {
                head=k
            }
        }

        s.Coefficient_Matrix[i],s.Coefficient_Matrix[head]=s.Coefficient_Matrix[head],s.Coefficient_Matrix[i]
        s.Expansion[i],s.Expansion[head]=s.Expansion[head],s.Expansion[i]
        if s.Coefficient_Matrix[i][i]==0 {
            continue;
        }

        for k:=i+1;k<n;k++{//高斯消元
            factor:=s.Coefficient_Matrix[k][i]/s.Coefficient_Matrix[i][i]
            s.Expansion[k]-=factor*s.Expansion[i]
            for j:=i;j<n;j++ {
                s.Coefficient_Matrix[k][j]-=factor*s.Coefficient_Matrix[i][j]
            }
        }
    }

    for i:=n-1;i>=0;i--{//化成对角矩阵
        if s.Coefficient_Matrix[i][i]!=0{
            x.Value[i]=s.Expansion[i]/s.Coefficient_Matrix[i][i]
            for j:=i;j>=0;j--{
                s.Expansion[j]-=x.Value[i]*s.Coefficient_Matrix[j][i]
            }
        } else{
            for j:=0;j<n;j++{
                x.Value[i]=0
            }
            if s.Expansion[i]==0{
                err=errors.New("方程有无数个解")
            }else{
                err=errors.New("方程无解")
            }
            return x,err
        }
    }
    return x,nil
}

func Output(ans Vec,ans_exit error)(x string){//传回前端的字符串
    if ans_exit==nil{
        x="讷讷～你要的解向量： "
        for i:=0;i<ans.Len;i++{
            x+=fmt.Sprintf("%.2f",ans.Value[i])+" "
        }
    }else if ans_exit.Error()=="方程有无数个解"{
        x="残念～方程有无数个解呢！"
    }else if ans_exit.Error()=="方程无解"{
        x="达咩～方程无解哦，请再好好地检查对我的输入"
    }else{
        x="苦路西～发生了未知错误"
    }
    return x
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