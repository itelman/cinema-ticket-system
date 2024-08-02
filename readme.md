# Cinema Ticket System

This application manages information about movie tickets. It allows adding movies, registering users, purchasing tickets, and canceling ticket purchases.

## Installation

### Prerequisites

- Go 1.22.5

### Installing

Instructions on how to install the application.

#### From Source

1. Clone the repository:

   ```sh
   git clone https://github.com/itelman/cinema-ticket-system.git
   ```

2. Navigate to the project directory:

   ```sh
   cd cinema-ticket-system
   ```

3. Run the application:

   ```sh
   go run main.go
   ```

## Usage

- You will encounter the main menu of the application:

   ```sh
   Здравствуйте, у вас есть следующие доступные функции:
   
   1. Добавить новый фильм
   2. Показать все доступные фильмы
   3. Добавить нового пользователя
   4. Купить билет
   5. Отменить покупку билета
   6. Выход
   
   Выберите опцию: 
   ```

- You will be prompted to select one of the options by typing the number.

### Examples

- Select the first option, and you will be prompted to type a movie name in order to add it to the system. You will receive the ID of the movie, which can be used to buy a ticket.

   ```sh
   Выберите опцию: 1

   Введите название фильма (или введите -1 чтобы вернуться в меню): The Inception
   Фильм добавлен. ID фильма: 1
   ```

## Features

- Adding a movie to the system.
- Viewing all available movies in the system.
- Adding a user to the system.
- Buying a ticket in the system.
- Cancelling a ticket in the system.