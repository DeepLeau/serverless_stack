
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

---

# 🚀 Deployment on AWS

Follow these steps to deploy the application on AWS Lambda and DynamoDB:


## 🏗️ 1. Build the Go Application

Compile the Go application into a binary compatible with AWS Lambda's environment:

```bash
GOOS=linux GOARCH=amd64 go build -o main main.go
```

This creates an executable binary named `main`.

---

## 📦 2. Package the Application

Create a ZIP file containing the binary:

```bash
zip -j build/main.zip main
```

Ensure the `build/` directory exists to store the ZIP file.

---

## ☁️ 3. Upload to AWS Lambda

1. Go to the **AWS Management Console**:
   - Navigate to the **Lambda** service.
2. Create a new Lambda function:
   - Choose **Author from scratch**.
   - Set the runtime to **Go 1.x**.
3. Upload the ZIP file:
   - Under the **Function code** section, upload `main.zip`.
4. Set environment variables:
   - Add `AWS_REGION` and any other required variables (e.g., `TABLE_NAME`) under the **Environment variables** section.

---

## 🗄️ 4. Create the DynamoDB Table

1. In the **AWS Management Console**:
   - Navigate to the **DynamoDB** service.
2. Create a table:
   - Set the table name to `LambdaInGoUser`.
   - Use `email` as the partition key (type: `String`).

---

## 🌐 5. Set Up API Gateway

1. Go to **API Gateway**:
   - Create a new API or use an existing one.
2. Add endpoints:
   - Configure methods (`GET`, `POST`, `PUT`, `DELETE`) for the `/user` resource.
3. Integrate with Lambda:
   - Link the methods to your Lambda function.

---

## 🚦 6. Test the Application

Use tools like **Postman** or `curl` to test your API endpoints:

- **Create User**: `POST /user`
- **Get User**: `GET /user`
- **Update User**: `PUT /user`
- **Delete User**: `DELETE /user`

---

## 🤝 Contribution

Contributions are welcome! To contribute:

1. Fork the repository.
2. Create a new feature branch.
3. Submit a pull request.

## 📜 License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
