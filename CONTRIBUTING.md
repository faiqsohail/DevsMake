# Contributing

It's awesome you're considering to contribute to this project, there are many ways to contribute. 
As DevsMake is an open-source project your contribution helps the project thrive and grow.

### Reporting Bugs

When reporting a Bug please be sure to follow the [Bug Report Template](https://github.com/faiqsohail/DevsMake/issues/new?assignees=&labels=&template=bug_report.md&title=) 
and provide as much detail as possible.


### Suggesting Enhancements

To suggest a feature or enhancement use the [Feature Request Template](https://github.com/faiqsohail/DevsMake/issues/new?assignees=&labels=&template=feature_request.md&title=),
be sure your request is inline with the scope of the project.

### Writing Code

You may write code to improve the project in any way deemed fit, improving the current code base to make it efficient, fixing reported bugs, implementing feature requests and more! 

Be sure to follow good coding style and be consistent with the current code.

To get started:
1) Clone the Project
2) Create a branch for your enhancement
3) Commit changes & Push
4) Finally submit a Pull Request!

### Setting up Development Environment

Required tools: [Docker](https://www.docker.com/products/docker-desktop), [Node](https://nodejs.org/en/download/), [Go](https://golang.org/dl/), [golang-migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate), [git](https://git-scm.com/downloads)

1) Clone the repository ```git clone git@github.com:faiqsohail/DevsMake.git $GOPATH/src/devsmake```
2) Change current directory to project root ```cd $GOPATH/src/devsmake```
3) Install frontend modules ```npm --prefix ./frontend install ./frontend```
4) Build containers and run: 
```
docker-compose build
docker-compose up -d
```
5) Run any migrations locally ```migrate -path migrations -database "mysql://devsmake:devsmake@tcp(localhost:3306)/devsmake" up```
6) Visit in your browser ```http://localhost:5001``` 🥳
