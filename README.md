### Installation

1.  Create a new repository based on this template by clicking the Use this template button at the top of this page.

2. Clone this repository to your local machine
    ```bash
    git clone https://github.com/Sw0xy/go-rest-api-template.git
    cd ./go-rest-api-template
    ```

    Create a .env file in the root directory based on the .env.example template and fill in your configuration details.

### Install Go packages:
```bash
go mod tidy
```


### Run 
```bash
go run cmd/main.go
```

### Project Structure
```
├── api/
│   ├── controller/
│   ├── middleware/
│   └── route/
├── bootstrap/
├── models/
├── cmd/
│   └── main.go
├── domain/
├── repository/
├── usecase/
├── utils/
```

### About Me

Hello, I am Emre, a fullstack developer from Turkey. I am interested in fullstack development and cyber security. You can find me on [LinkedIn](https://www.linkedin.com/in/ihsan-emre/).