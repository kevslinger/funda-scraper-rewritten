# Funda Scraper Rewritten

This repo is a rewrite of my original repo, [Funda Scraper](https://github.com/kevslinger/funda-scraper).
The idea is to use what I learned while writing the app the first time to improve the design, make the repo more usable and easier to read.
The redesign will come in phases and will start from the ground up, from a very simple app with virtuall no options to a more complex, while still easy to undertsand, with more functionality.

## Phase 1: Laying the Groundwork

In the initial implementation, we will set up the structure of the application.
Simply put, we need a `main` function and something for the `main` function to run.
The `main` function will be stored in `cmd/funda-scraper-rewritten/main.go`, and it will import nd use the `fundascraperrewritten` package (top-level `funda-scraper-rewritten.go`).
At this point, only the name gives us any indication of what this application will be doing.

## Phase 2: Implementing the Core Functionality

Once we have the core built, we will add the basic functionality: making a housing query on https://www.funda.nl.
The application will not accept any arguments, and will return the result of the query to the command line.

## Phase 3: Specifying the request

At this point, we can introduce command-line flags to allow the user to filter their search query with basic demands they may have for their house: e.g. a maximum price, a minimum number of bedrooms, and a minimum size (in square meters).

## Installation and Running

 Installation

```bash
go install github.com/kevslinger/funda-scraper-rewritten
```

Running

```bash
funda-scraper-rewritten
```

You can supply optional arguments `search-area`, `max-price`, `min-bedrooms`, `min-square-meters`. E.g.

```bash
funda-scraper-rewritten -search-area utrecht --maximum-price 575000 --minimum-bedrooms 3 --minimum-square-meters 100
```
