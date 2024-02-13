package slice

import "fmt"

func MainSlice() {
	// var users0 []string
	// var users1 = []string{"Tom", "Alice", "Kate"}
	// или так
	users := []string{"Tom", "Alice", "Kate"}
	secondUsers := []string{"Tom2", "Alice2", "Kate2"}
	initialUsers := [8]string{"Bob", "Alice", "Kate", "Sam", "Tom", "Paul", "Mike", "Robert"}
	users1 := initialUsers[2:6] // с 3-го по 6-й

	// for _, value := range users {
	// fmt.Println((value))
	// }

	// users = append(users, "Bob")

	// for _, value := range users {
	// 	fmt.Println(value)
	// }

	for _, value := range secondUsers {
		users = append(users, value)
	}

	// for _, value := range users {
	// 	fmt.Println(value)
	// }
	fmt.Println(users1)
	fmt.Println(users)

}

func UserDel() {
	users := []string{"Bob", "Alice", "Kate", "Sam", "Tom", "Paul", "Mike", "Robert"}
	var n = 3
	users = append(users[:n], users[n+1:]...)
	fmt.Println(users)
}
