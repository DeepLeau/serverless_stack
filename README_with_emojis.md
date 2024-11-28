
# 🚀 AWS Lambda User Management with DynamoDB

This project demonstrates a serverless application for user management, built using AWS Lambda, DynamoDB, and Go. The application allows CRUD operations on user data stored in DynamoDB.

## ✨ Features

- **🛠️ User CRUD Operations**:
  - ➕ Create a user (`POST`).
  - 🔍 Retrieve a user or list of users (`GET`).
  - ✏️ Update user information (`PUT`).
  - ❌ Delete a user (`DELETE`).
- **☁️ Serverless Framework**: Designed to work seamlessly on AWS Lambda.
- **✔️ Validation**: Includes email validation for user inputs.
- **⚠️ Error Handling**: Comprehensive error messages for edge cases.

## 🗂️ Project Structure

- **`main.go`**: Initializes the application and configures the AWS Lambda handler.
- **`handlers/`**: Contains the logic for handling API Gateway requests.
- **`dynamoapi.go`**: Interfaces with DynamoDB.
- **`user.go`**: Implements CRUD operations for users.
- **`validators/`**: Includes utility functions such as email validation.

## 🔧 Prerequisites

- 🐹 Go installed on your machine.
- 🛠️ AWS CLI configured with appropriate credentials.
- 🗄️ DynamoDB table created with the following details:
  - **Table Name**: `LambdaInGoUser`
  - **Primary Key**: `email` (String).

## 🛠️ Setup and Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/your-repo/serverless-user-management.git
   cd serverless-user-management
   ```

2. Install dependencies:

   ```bash
   go mod tidy
   ```

3. Set the AWS Region:

   Ensure the `AWS_REGION` environment variable is set, e.g.:

   ```bash
   export AWS_REGION=us-west-2
   ```

4. Deploy the Lambda function (if using a framework like SAM or Serverless).

## 📒 Usage

### 📋 API Endpoints

1. **➕ Create User**
   - **Method**: `POST`
   - **Endpoint**: `/user`
   - **Request Body**:
     ```json
     {
       "email": "example@example.com",
       "firstName": "John",
       "lastName": "Doe"
     }
     ```

2. **🔍 Get User**
   - **Method**: `GET`
   - **Endpoint**: `/user?email=example@example.com`

3. **✏️ Update User**
   - **Method**: `PUT`
   - **Endpoint**: `/user`
   - **Request Body**:
     ```json
     {
       "email": "example@example.com",
       "firstName": "Jane",
       "lastName": "Doe"
     }
     ```

4. **❌ Delete User**
   - **Method**: `DELETE`
   - **Endpoint**: `/user?email=example@example.com`

## ⚠️ Error Handling

Common error responses include:
- ❌ `400 Bad Request`: Invalid input or missing required parameters.
- ❓ `404 Not Found`: Resource not found.
- ⚙️ `500 Internal Server Error`: Server-side issues.

## 🤝 Contribution

Contributions are welcome! To contribute:

1. Fork the repository.
2. Create a new feature branch.
3. Submit a pull request.

## 📜 License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
