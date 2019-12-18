package main

import "log"

func sum(x,y int64,ch chan int64)  {
	ch<-x+y
}

func main1()  {
	sliceAll:=[][]int64{}
	for i:=3;i>0 ;i--  {
		ch1:=make(chan int64,10)
		a:=[]int64{}
		for j:=i;j>0 ;j--  {
			go sum(int64(j),int64(j+1),ch1)
			n1:=<-ch1
			a= append(a, n1)
			log.Println(n1)
		}
		sliceAll= append(sliceAll, a)
	}
	log.Println(sliceAll)
}
func main()  {
	var user struct{
		name string
		pwd string
	}
	user.name="asd"
	user.pwd="222"
	log.Println(user)
}