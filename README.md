# Bills API

This is a sample API for managing bills. It provides endpoints to perform CRUD operations on bills and calculate expenses.

## Prerequisites

Before running the application, make sure you have the following installed:

- Go (Golang) 1.16 or higher


## Getting Started

1. Clone the repository:

   ```bash
   git clone https://github.com/seifguerbouj/bills.git
   cd bills
   go run main.go

## Swagger

1. Access swagger from http://localhost:8080/docs/index.html
## API Endpoints

### Get all bills

- **Method:** `GET`
- **URL:** `/bills`
- **Description:** Get all bills.
- **Response:** Status code `200 OK` with an array of bill objects.

### Create a new bill

- **Method:** `POST`
- **URL:** `/bill`
- **Description:** Create a new bill.
- **Request Body:** JSON object with the following properties:
  - `title` (string, required): Title of the bill.
  - `author` (string, required): Author of the bill.
  - `amount` (number): Amount of the bill.
  - `date` (string): Date of the bill.
- **Response:** Status code `200 OK` with the created bill object.

### Calculate total expenses

- **Method:** `GET`
- **URL:** `/expenses`
- **Description:** Calculate total expenses.
- **Response:** Status code `200 OK` with the total expenses.

### Get bills for a specific month

- **Method:** `GET`
- **URL:** `/bill/{date}`
- **Description:** Get bills for a specific month.
- **URL Parameter:** `date` (string, required): Date in the format MM.
- **Response:** Status code `200 OK` with an array of bill objects.

### Update a bill

- **Method:** `PUT`
- **URL:** `/bill/{id}`
- **Description:** Update a bill with the specified ID.
- **URL Parameter:** `id` (string, required): ID of the bill to update.
- **Request Body:** JSON object with the following properties:
  - `title` (string, required): Updated title of the bill.
  - `author` (string, required): Updated author of the bill.
  - `amount` (number): Updated amount of the bill.
  - `date` (string): Updated date of the bill.
- **Response:** Status code `200 OK` with the updated bill object.

### Delete a bill

- **Method:** `DELETE`
- **URL:** `/bill/{id}`
- **Description:** Delete a bill with the specified ID.
- **URL Parameter:** `id` (string, required): ID of the bill to delete.
- **Response:** Status code `200 OK` with the deleted bill object.

