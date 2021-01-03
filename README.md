# Cross Platform Desktop Apps With Golang

With golang we can build apps that run on any OS. 

We will use a desktop framework called [wails](https://github.com/wailsapp/wails) for this talk.

## Prerequisites

* Go >= 1.13
* Npm
* Docker

### Additional installations
#### Windows 
* Install [gcc](https://jmeubank.github.io/tdm-gcc/)

#### MacOS
* Install xcode


### WAILS
Install wails using ``` go get -u github.com/wailsapp/wails/cmd/wails```


* Run `wails init` to initialize the project and follow the project setup prompt.
* Build the project using `wails build`
* Pull docker image for cross compilation ```docker pull wailsapp/xgo```
* Use ```wails build -x <os specification>``` to build for any os. 
 ```  
- darwin/amd64
- linux/amd64
- linux/arm-7
- windows/amd64 
 ```
