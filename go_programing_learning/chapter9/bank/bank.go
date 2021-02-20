/**
 * @Author $
 * @Description //TODO $
 * @Date $ $
 * @Param $
 * @return $
 **/
package bank

var (
	deposits = make(chan int)
	balances = make(chan int)
	withdraws = make(chan int)
)

func Deposit(amount int) {
	deposits <- amount
}

func Balance() int {
	return <-balances
}

func Withdraw(amount int) bool {
	if b := <-balances; amount > b {
		return false
	}
	withdraws <- amount
	return true
}

func teller() {
	balance := 0
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case amount := <-withdraws:
			balance -= amount
		case balances <- balance:
		}
	}
}

func init() {
	go teller()
}
