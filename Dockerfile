FROM golang:1.22.5-bullseye as development

RUN go install github.com/cosmtrek/air@v1.49.0

# Set working directory
WORKDIR /app

# Copy only necessary files to cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy project files
COPY . /app

# Command for Air
CMD ["air"]


FROM golang:1.22.5-bullseye as production

WORKDIR /app
COPY go.sum go.mod /app/

RUN go mod download
COPY . /app

RUN CGO_ENABLED=0 go build -o /bin/app ./cmd/

#FROM scratch
#COPY --from=0 /bin/app /bin/app
ENTRYPOINT ["/bin/app"]

