<h1 align="center">File Tagger</h1>

## Table of contents
- [Table of contents](#table-of-contents)
- [Description](#description)
- [Usage](#usage)
- [Author](#author)

## Description
This project is a simple Go application that generates random user data and tag sensitive information in a text file. It consists of two main functionalities:

- **Random User Generator**: This functionality generates random user data and saves it to a text file. Each user entry includes a unique ID, username, email, password, first name, last name, phone number, address, street number, and zipcode.

- **File Tagger**: The file tagger reads an existing text file and tags sensitive information by replacing occurrences of "user" and "password" with "[RISK] user" and "[RISK] password," respectively.

## Usage

1. Clone the repository and navigate to his directory

If you have **Go 1.21** installed:

2. Generate random user data and tag sensitive information in the file by running:

        go run main.go
        
Alternatively, you can run the application in a **Docker** container.

1. Build the Docker image:

        docker build -t tagfile .

2. Run the Docker container:

        docker run -v $(pwd):/app tagfile

## Author
| [<img src="https://avatars.githubusercontent.com/u/105685220?v=4" width=115><br><sub>Pedro Vinicius Messetti</sub>](https://github.com/pedromessetti) |
| :---------------------------------------------------------------------------------------------------------------------------------------------------: |
