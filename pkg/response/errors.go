package response

import "fmt"

func RecordExistsError() {
	fmt.Println("Билет, пользователь или фильм с такими данными уже существует. Пожалуйста, введите другие данные.\n")
}

func BadRequestError() {
	fmt.Println("Неверный запрос. Пожалуйста, попробуйте еще раз.\n")
}

func NotFoundError() {
	fmt.Println("Билет, пользователь или фильм не найдены. Пожалуйста, попробуйте еще раз, введя корректные данные.\n")
}
