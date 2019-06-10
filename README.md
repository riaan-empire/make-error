# make-error
Generate errors to test monitoring

# Requirements
 - Docker
# Setup
 - `$ git clone git@github.com:EmpireLife/make-error.git`
 - `$ cd make-error`
 - `$ docker build -t make-error .`
 - `$ docker run -d --name make-error -p 8080:8080 make-error:latest`

 # Usage
Make multiple consecutive calls to the api to receive random error/success calls back.

### Bash
`$ cmd="curl localhost:8080"; for i in $(seq 10); do $cmd; sleep 1; done`

### ZSH
`$ repeat 10 {curl localhost:8080; sleep 1}`