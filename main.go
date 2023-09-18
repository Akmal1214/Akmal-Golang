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

package main 

import "fmt"

func main() {

	 if i := recover(); i == nil{
		fmt.Println("Ну вот так вот")
	 }
	
    fmt.Println("Приложения успешно выполнилось")
}

