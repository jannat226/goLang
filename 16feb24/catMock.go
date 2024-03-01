package main
import "fmt"
func main(){
	fmt.Println("Welcome!Choose your drink :\n1->Coke  \n2->Pepsie \n3->Water")
	var amount,choice int
	pAmount := 32
	cAmount := 30
	fmt.Println("Enter your choice")
	fmt.Scanln(&choice)
	if choice == 1{
		fmt.Println("Enter the amount %d",cAmount)
		fmt.Scanln(&amount)
		if amount == cAmount {
			fmt.Println("Enjoy your drink")
		}else{
			if amount >cAmount{
				var returnedMoney int = amount-cAmount
				fmt.Println("Your returned money is %d , Enjoy your drink!", returnedMoney)
			}else{
				var less = cAmount - amount
				fmt.Println("Your funds are less by  %d bucks", less)
			}
		}
	}else if choice == 2{
		fmt.Println("Enter the amount %d",pAmount)
		fmt.Scanln(&amount)
		if amount == pAmount {
			fmt.Println("Enjoy your drink")
		}else{
			if amount >pAmount{
				var returnedMoney int = amount-cAmount
				fmt.Println("Your returned money is %d ,Enjoy your drink!", returnedMoney)
			}else{
				var less = pAmount - amount
				fmt.Println("Your funds are less by  %d bucks", less)
			}
		}
	}else if choice == 3{
		fmt.Println("Water is free, Enjoy your drink")
	}
	else{
		fmt.Println("Entered choice is invalid")
		main()
	}
}