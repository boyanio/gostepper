# Go Stepper

Go Stepper collects how many steps you have done each day and generates statistics for you. I use this project when I teach Go as it shows the basics of the language.

## Docker

```
docker build . -t gostepper
docker run -it --rm -p 8080:8080 gostepper
```