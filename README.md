# Ice Breaker GO
## Project Overview

This project is a command-line tool that aims to fetch and parse LinkedIn and Twitter profile information using Google search engine. It utilizes Langchai
n's LLM agents for text processing and Selenium WebDriver for web scraping.

## Table of Contents
1. [Installation](#installation)
2. [Usage](#usage)
3. [Components Overview](#components-overview)
   * [GoogleSearchLinkedIn](#googlesearchlinkedin)
   * [GoogleSearchTwitter](#googlesearchtwitter)
4. [Project Structure](#project-structure)

## Installation
To install the project, follow these steps:
1. Clone or download the repository to your local machine.
2. Install required packages by running `go get` in the terminal.

```bash
$ git clone https://github.com/yourusername/repository_name.git
$ cd repository_name
$ go get -u github.com/tmc/langchaingo
$ go get -u github.com/dragaoazuljr/ice-breaker-go
$ go get -u github.com/selenium/go-selenium
```

## Usage
To use the project, execute the `main.go` file in your terminal:

```bash
$ go run cmd/app/main.go <input_string>
```
Replace `<input_string>` with a person's name or any search query related to the desired profile information. The tool will return formatted LinkedIn and 
Twitter links, along with extracted data for each profile in a structured format.

## Components Overview
### GoogleSearchLinkedIn
This is a custom tool in `tools/google_search_linkedin.go` file that extends Langchain's LLM agent to find LinkedIn profiles using the Google search engin
e.

### GoogleSearchTwitter
Similarly, this is another custom tool in `tools/google_search_twitter.go` file that extends Langchain's LLM agent to find Twitter profiles using the Goog
le search engine.

## Project Structure
The project consists of multiple packages:
- `scrapper`: Contains the Selenium WebDriver functionality for web scraping LinkedIn and Twitter profile pages.
- `tools`: Contains the custom LLM tools for fetching information from LinkedIn and Twitter profiles using Google search engine.
- `main.go`: The entry point of the project, which initializes, calls, and processes the data from both custom LLM agents.

