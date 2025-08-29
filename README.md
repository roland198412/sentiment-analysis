# Sentiment Analysis with Go and MySQL

This is a Go-based project that performs sentiment analysis on text data retrieved from a MySQL database. It uses the **VADER sentiment analysis model** (via the `govader` library) to analyze the sentiment of data stored in the database.

## Features

- Connects to a MySQL database using the `go-sql-driver/mysql` package.
- Extracts data from a specified database table and column.
- Analyzes sentiment (positive, neutral, and compound scores) using the `govader` library.
- Outputs sentiment scores to the console.

---

## Prerequisites

Before running the project, ensure the following software and libraries are installed:

### Software:

- **Go SDK** (1.24+)
- **MySQL server** (configured and running)

### Libraries:

The project requires the following Go libraries, which can be installed using the `go get` command:

1. MySQL driver for Go:
   ```bash
   go get -u github.com/go-sql-driver/mysql
   ```
2. GoVader (for sentiment analysis):
   ```bash
   go get github.com/jonreiter/govader
   ```

---

## Environment Variables

Set up the following environment variables to configure the database connection and query details:

- `DB_USER`: MySQL database username.
- `DB_PASS`: MySQL database password.
- `DB_HOST`: MySQL server address (e.g., `127.0.0.1:3306` for localhost).
- `DB_NAME`: Name of the database to connect to.
- `TABLE_NAME`: Name of the table containing the text data.
- `TABLE_COLUMN_NAME`: Column within the table that contains the text data to be analyzed.

---

## How It Works

1. The project retrieves table data from the MySQL database using the `sql.Open` and `Query` methods.
2. The `govader` library is used to calculate sentiment scores (Compound, Positive, Neutral).
3. Sentiment scores are printed to the console for each row of data.

### Code Structure:

- **Database Connection**:
  Configured with environment variables. Connection handled with `database/sql` and `go-sql-driver/mysql`.
- **Sentiment Analysis**:
  Leverages the `govader` library, which implements the VADER sentiment analysis algorithm.
- **Functions**:
    - `main`: Application entry point—initializes database connection, retrieves data, and performs sentiment analysis.
    - `getColumnData`: Fetches data from a specified column in the database for sentiment analysis.

---

## How to Run

1. Clone the repository:

   ```bash
   git clone <repository-url>
   cd <repository-directory>
   ```

2. Set the necessary environment variables. For example:

   ```bash
   export DB_USER="your-username"
   export DB_PASS="your-password"
   export DB_HOST="127.0.0.1:3306"
   export DB_NAME="your-database-name"
   export TABLE_NAME="your-table-name"
   export TABLE_COLUMN_NAME="your-column-name"
   ```

3. Run the application:

   ```bash
   go run main.go
   ```

4. View the sentiment analysis results in the terminal, including compound, positive, and neutral scores.

---

## Example Output

When the program runs successfully, you will see output similar to the following:
