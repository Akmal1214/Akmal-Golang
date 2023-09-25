/*package main 

import "fmt"

func main() {
	fmt.Println("Z cnfye Golang ninja")
} */

/*package main 

import "fmt"

func main() {
	message := "Z cnfye Golang ninja"
	fmt.Println(message)
} */

/*package main 

import "fmt"

func main() {
	var message string
	message = "Z cnfye Golang ninja"
	fmt.Println(message)
} */

/*package main 

import (
	"errors"
	"fmt" 
	"log"
)

func main() {
	message, err := enterTheClub(66)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}



func enterTheClub(age int) (string, error) {
	if age >= 18 && age < 45 {
		return "Входи", nil
	} else if age >= 45 && age < 65 {
		return "Вам точно понравиться это музыка", nil
	} else if age >= 65 {
		return "Вам не понравиться это музыка", errors.New("You are too old")
	}

	return "Тебе нет 18-ти", errors.New("You are too young")
	
}*/

/*package main 

import (
	"errors"
	"fmt" 
	
)

func main() {
	fmt.Println(prediction("jhbjhb"))
}

func prediction(dayOfWeek string) (string, error) {
	switch dayOfWeek {
	case "пн":
		return "Хорошего тебе начало недели", nil
	case "вт":
		return "Прекрасного тебе вторника", nil
	case "ср": 
	    return "Замечателтное тебе среды", nil
	case "чт":
		return "Четверг - это маленькая пятница!", nil
	default: 
	return "Невалидный день недели", errors.New("Unknown day of the week")
	}
}*/

/*package main

import "fmt"

func main() {
	inc := increment()

	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())

	fmt.Println(incremetn2())
	fmt.Println(incremetn2())
	fmt.Println(incremetn2())
	fmt.Println(incremetn2())
	fmt.Println(incremetn2())
}

func increment() func() int{
	count := 0
	return func() int {
		count++
		return count
	} 
}

func incremetn2() int {
	count := 0
	count++ 
	return count
}*/

/*package main

import "fmt"

func init() {
msg := "функция init"

fmt.Println(msg)
}

func main() {
message := "функция main"

fmt.Println(message)
}*/

/*package main

import "fmt"

func main() {
	message := "Скоро я стану golang нинзя"
	printMessage(&message)

	fmt.Println(message)
}

func printMessage(message *string) {
	*message += ("Из функции printMessage()")
	fmt.Println(message)
}*/


/* package main 

import (
	"fmt"
	"sort"
)

func main() {
	slice := []int{3,5,7,2,8}
	slice = append(slice, 45)
	slice[2] = 135
	sort.Ints(slice)
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
}*/

/*package main 

import "fmt"

func main() {

	 if i := recover(); i == nil{
		fmt.Println("Ну вот так вот")
	 }
	
    fmt.Println("Приложения успешно выполнилось")
}*/

/*package main 

import "fmt"

func main() {
	users := map[string]int {
		"Аслан": 13,
		"Акмаль": 15,
		"Аскар": 29,
	}

	 users["Акмаль"] = 16
	 delete(users, "Аслан")

	fmt.Println(users)
}*/

/*package main 

import "fmt"

type User struct {
	name string
	age int64
	password string
}

func main() {
	user := User{name:"Tokuchi", age:27, password: "pass"}
	user.name = "Jhon"
	fmt.Println(user.nameUser())
}

func (g User) nameUser() string{
	return g.name
}*/

/*package main 

import "fmt"

type NumberStruct struct{ // создаём структуру 
	num1 int 
	num2 int
}

type Number interface { // создаём интерфейс 
	Sum() int 
	Multiply() int 
	Division() float64
	Substract() int
}

func (n NumberStruct) Sum() int{ // создаем метод для структуры, унаследуем метод интерфейса для структуры 
	return n.num1 + n.num2
}

func main(){
	var i NumberStruct // создаем переменную которая будет вместить наш интерфейс 
	numbers := NumberStruct{15, 5} // создаем нашу структуру на основе чисел
	i = numbers // в интерфейс мы помещаем нашу структуру 
	fmt.Println(i.Sum()) 
}*/

/*package main 

import "fmt"

func main() {
	for i := 0; i <= 10; i++{
		if i == 8{
			continue
		}
		fmt.Println(i)
	}
}*/

/*package main

import "fmt"

func main() {
	numbers := [3][5]int {{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(numbers[2][2])
}*/

/* ------------------ДОМАШНИЕ ЗАДАНИЯ---------------*/

/*package main

import (
	"fmt"
	"time"
	"github.com/zhashkevych/scheduler"
	"cache"
)

type Example struct {
	Name string 
	age int 
}


func main() {
	newExample := &Example {
		Name: "Akmal",
		Age:  15,
	}

	c := cache.New(3*time.Minute, 5*time.Minute)

	abcd, found := c.Get("abcd")

	if found {
		a:= abcd.([]Example)
		a = append(a, *newExample)
		c.Set("abcd", a, cache.NoExpiration)
		fmt.Println("Ну короче, да")
	}else {
		c.Set("abcd", newExample, cache.NoExpiration)
		fmt.Println("Ну короче, нет")
	}

	abcd, found = c.Get("abcd")
	if found {
		fmt.Println(abcd)
	}
}*/

/*package main 

import (
	"fmt"
	"github.com/jellydator/ttlcache/v2"
)

func main() {
	cache := cache.New()

	cache.Set("Akmal", 14)
	userId := cache.Get("userId")

	fmt.Println(userId)

	cache.Delete("userId")

	fmt.Println(userId)
}*/

package main

import "fmt"

func main() {
fmt.Println("Hello world")
}
