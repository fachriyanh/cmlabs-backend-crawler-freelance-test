# cmlabs-backend-crawler-freelance-test

This is a web crawler application written in GoLang that can crawl websites and generate HTML files.

### Prerequisites

Make sure you have GoLang installed on your machine. You can download it from the official website: [https://golang.org/](https://golang.org/)

### Installation

Clone the repository to your local machine:

```bash
git clone https://github.com/your-username/cmlabs-backend-crawler-freelance-test.git

cd cmlabs-backend-crawler-freelance-test
```

### Dependencies

This project uses Go modules for dependency management. Before running the application, ensure that you have the required dependencies. You can install them using:

```bash
go mod tidy
```

### Usage

1. Edit the websiteList map in the main.go file to add the websites you want to crawl:

```bash
websiteList := map[string]string{
	"example1": "https://www.example1.com",
	"example2":  "https://www.example2.com",
	// Add more websites here
}
```

2. Run the crawler:

```bash
go run main.go
```

Happy web crawling!
