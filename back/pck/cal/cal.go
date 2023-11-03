package pck

import(
	
)

type equ struct {
    value_num int //方程维数
    cft [][]int  //系数矩阵
}


//初始化
func(s equ) Init(){
    s.cft =make([][]int,s.value_num)
    for i:=0;i<s.value_num;i++ {
        s.cft[i]= make([]int,s.value_num)
    }
}

func cal_2(n int,A [][]int,Y []int) []float64{//处理二阶的简单版本
    X:= make([]float64,n,2*n);
    var un float64;
    un=float64(A[0][0]*A[1][1]-A[1][0]*A[0][1]);
    X[0]=(float64(Y[0]*A[1][1]-Y[1]*A[0][1]))/un;
    X[1]=(float64(Y[1]*A[0][0]-Y[0]*A[1][0]))/un;
    return X;
}