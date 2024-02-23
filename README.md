# wallet-api
A Go-based backend API with Gin framework for managing user wallets, transactions, and authentication. Features include user registration, wallet creation, fund credit/debit, and transaction history retrieval.


To run the project locally, you'll need to follow these steps:

Clone the Repository
git clone https://github.com/Themmythorpe/wallet-api.git



Set up Environment Variables


DB_USERNAME=
DB_PASSWORD=
DB_HOST=127.0.0.1
DB_PORT=3306
DB_NAME=wallet_api


Install Dependencies

go mod tidy


Run migrations


migrate -path db/migrations -database "mysql://username:password@tcp(hostname:port)/dbname" up

Build the Application


go build -o walletapp

Run the Application


./walletapp


1. **User Routes**:
   - **POST /users/register**: Register a new user with a unique email and password. Required parameters: email, password.
   - **POST /users/login**: Authenticate user login with email and password. Required parameters: email, password.

2. **Wallet Routes**  (Bearer token required):
   - **POST /wallets/create**: Create a new wallet for a user with a specific currency.
   - **POST /wallets/:wallet_id/credit**: Credit funds to a wallet specified by its ID.
   - **POST /wallets/:wallet_id/debit**: Debit funds from a wallet specified by its ID.
   - **GET /wallets/:wallet_id/transactions**: Retrieve transaction history for a wallet specified by its ID.

