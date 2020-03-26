package tcp

import (
	"fmt"
	"net"
	"os"
	"sort"
	"time"
)

func ConnectPort(port chan int,result *[]int){
	for p := range port{
		url := fmt.Sprintf("scanme.nmap.org:%d",p)
		_,err := net.Dial("tcp",url)
		if err != nil {
			continue
		}
		*result = append(*result,p)
	}
}

func main(){
	t := time.Now()
	ports := make(chan int,8192)
	result := make([]int,0)
	for i:=1;i<=cap(ports);i++{
		go ConnectPort(ports,&result)
	}
	for i:=1;i<=65535 ;i++  {
		ports <- i
		fmt.Fprintf(os.Stdout,"Port:%d is scaning\r",i)
	}
	close(ports)
	sort.Ints(result)
	for _,value := range result{
		fmt.Printf("Port:%d is Open\n",value)
	}
	fmt.Println(time.Since(t))
}