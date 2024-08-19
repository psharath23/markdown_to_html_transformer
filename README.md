# Markdown to HTML transformer

## installation
go get .

## usage
put a html content in a file preferably in the current project root directory

run the command go run main.go "input markdown file name"

output html file will be created in the same folder with .html extension

## testing
for seeing the test coverage, please run below command which will show the ouput of test coverage
go test -v -coverprofile cover.out . 