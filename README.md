# Web+GaussCalculator
## UESTC-go语言和区块链技术-期中设计

### 这是老师发的项目要求：
Go 设计一个基于 Web 的非齐次线性
方程组求解器，以 Web 方式输入非齐次方程
组信息，传入到 Go 服务后端进行计算，计
算结果返回 Web 前端。形成期中报告，14 周，
周日以前提交。

    报告内容包括以下部分：

    （1） 设计思想

    （2） 系统架构

    （3） 关键技术和算法论述

    （4） 运行效果截图

    （5） 附件代码
    
    提 交 PDF 文 档 ， 报 告 提 交 到
    52524599@qq.com，文件名：Go 期中报告-
    名字-学号.pdf
    发现互相抄袭，双方一律 0 分处理！


听起来还是挺难的，但是实际做起来思路还是挺清晰的；

先说说如何运行这个http服务：

可能需要的依赖：go、python、pip、pyinstaller

>对于linux平台，在/release下打开终端，输入`./GaussCaculator_ubuntu20.04_version1.0`,进入`http://localhost:8080`即可启动运行程序;

>对于windows11,点击`run.bat`即可启动程序；

>通用办法：在/release/dist目录下寻找适合操作系统的可执行文件，这是基于pyinstaller生成，运行/release下的`run.py`

经过测试，即便挂着代理服务应该也能访问，但防火墙需要关掉；


* 不仅仅局限于n个n元非齐次线性方程的求解，更一般意义上的m维n元非齐次线性方程能否稍作修改求解；
* 对于方程有无穷解的情况能否进一步优化，显式表达出来；
* 对于前端输入的并不是特定的数而是计算式组成的字符串，能否计算出来并且正常输入到后端，同时进行严格的输入检查；
* 对于前端页面能否进一步美化，前端能否实现多个URL跳转；
* 对于多平台的适应（Linux Ubuntu20.02,Windows11）,脚本执行文件和命令行的编写；
