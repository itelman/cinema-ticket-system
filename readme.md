# Cinema Ticket System

This application manages information about movie tickets. It allows adding movies, registering users, purchasing tickets, and cancelling ticket purchases.

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

- Select the third option, and you will be prompted to type a user in order to add it to the system. You will receive the ID of the user, which can be used later to purchase a ticket.

   ```sh
   Выберите опцию: 3

   Введите имя пользователя (или введите -1 чтобы вернуться в меню): Alice

   Пользователь добавлен. ID пользователя: 1
   ```

- Select the fourth option, and you will be prompted to type user ID and movie ID in order to purchase a ticket. You will receive the ID of the ticket, which can be used later to cancel the purchase.

   ```sh
   Выберите опцию: 4
   
   Введите ID пользователя (или введите -1 чтобы вернуться в меню): 1
   Введите ID фильма (или введите -1 чтобы вернуться в меню): 1

   Билет успешно оформлен. ID билета: 1
   ```

- Select the fifth option, and you will be prompted to type ticket ID in order to cancel the purchase.

   ```sh
   Выберите опцию: 5
   
   Введите ID билета (или введите -1 чтобы вернуться в меню): 1

   Покупка билета отменена.
   ```

## Features

- Adding a movie to the system.
- Viewing all available movies in the system.
- Registering a user in the system.
- Purchasing a ticket in the system.
- Cancelling ticket purchases in the system.