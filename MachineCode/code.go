package main

import "fmt"

var reg [10] int

func subfun(a,b int){
	op:=a
	if(op==3){
		loadregister(b)
	}else if(op==5 || op==6 || op==7){
		operation(a,b)
	}	
}

func loadregister(val int){
	fmt.Println("Register inputed")
	rval:= val%10000
	regnum:= rval/1000
	rvalu:= val%1000
	reg[regnum]=rvalu
}

func operation(obj,opval int){
	fmt.Println("Operations:\n")
	regtemp:=opval%10000
	reg1:=regtemp/1000
	reg2:=opval%10
	if(obj==5){
		reg[reg1]=reg[reg1]+reg[reg2]	
	}else if(obj==6){
		reg[reg1]=reg[reg1]-reg[reg2]	
	}else if(obj==7){
		reg[reg1]=reg[reg1]/reg[reg2]	
	}else if(obj==8){
		reg[reg1]=reg[reg1]*reg[reg2]	
	}
}

func display(ival,dval int){
	fmt.Printf("10%d:%d",ival,dval)
	for i:= 0;reg[i]!=0;i++{
		t:=(30+i)*1000+reg[i]
		fmt.Printf("\nRegister value:%d\n",t)
	}
}

func main(){	
	var str [10] int
	str[0]=30012
	str[1]=31010
	str[2]=50001
	str[3]=32006
	str[4]=50002

	for i:= 0;str[i]!=0;i++{
		k:=str[i]/10000
		subfun(k,str[i])
		display(i,str[i])
	}	
}
/*
000000       Halt
03rnnn       Load register r with the number nnn.
04r00s       Load register r with the memory word addressed by register s.
05r00s       Add contents of register s to register r
06r00s       Sub contents of register s from register r
07r00s       Mul contents of register r by register s
08r00s       Div contents of register r by register s
*/
