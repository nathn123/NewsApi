# Readme

## To Run

To run this Project type `go run main.go` in the root directory

## Routes

| Route | Params | Description |
|-------|--------|-------------|
| all | none | this route will retrieve all newsfeeds |
| name | name | this route will retrieve one newsfeed by name using a query string |
| filter | provider, category | this route will retrieve multiple newsfeeds using provider and category as query strings |
| add | name, provider, category, url | this route will add a new newsfeed to the API |

## Known Issues
 * Errors may lead to the API stopping
 * validation on the url is only that it exists entering a non-url string will work and crash the API when requested
 * Stricter response codes could potentially be used for bad requests