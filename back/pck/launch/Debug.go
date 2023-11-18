package launch

import (
	"back/pck/cal"
	"fmt"
)

// 在终端启动。
func Debug() {
	var (
		Mat cal.Equation
        ans cal.Vec
        ans_exit error
        //ans_str cal.Answer
        ans_str string
	)
	for {
		fmt.Print(">>> ")
		fmt.Scanf("%d", &Mat.Value_num)
		n:=Mat.Value_num
		Mat.Init(n)
		Mat.Print()
		for i:=0;i<n;i++{
			for j:=0;j<n;j++{
				fmt.Scan(&Mat.Coefficient_Matrix[i][j])
			}
			fmt.Scanf("%f",&Mat.Expansion[i])
		}
		Mat.Print()
		ans,ans_exit= cal.Gauss(Mat)
		ans_str=cal.Output(ans,ans_exit)
		fmt.Println(ans_str)
		fmt.Println("------------------------------")
	}
}