FROM golang:1.22.1-bookworm as builder

RUN mkdir /home/recipe_calculator && cd /home/recipe_calculator

COPY . .

RUN make build 

FROM builder 

WORKDIR /home/recipe_calculator
ENTRYPOINT [ "bin/recipe_calculator" ]