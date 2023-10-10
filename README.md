# Go Receipts - Receipt Processor

This project processes receipts and calculates points based on given rules. 

## Getting Started

These instructions will guide you on how to run the service locally and with Docker.

### Prerequisites

- [Go](https://golang.org/dl/) (at least version 1.16)
- [Docker](https://docs.docker.com/get-docker/)

### Setup & Running (without Docker)

1. Clone the repository:

```bash
git clone https://github.com/salomonj11/Go-Receipts.git
cd Go-Receipts
```

2. Install the required dependencies:

```bash
go mod download
```

3. Run the application:

```bash
go run main.go
```

The application will start, and you can access the API endpoints at `http://localhost:8080`.

### Docker Setup and Running

#### Building the Docker Image

Build the Docker image for the service:

```bash
docker build -t go-receipts .
```

This will package the application into a Docker container with all necessary dependencies.

#### Running the Service in a Docker Container

To run the service inside the Docker container:

```bash
docker run -p 8080:8080 go-receipts
```

The service will be accessible on your host machine at `http://localhost:8080`.

### Accessing the Application

Once the Docker container is running, you can send requests to `http://localhost:8080/receipts/process` to process receipts and `http://localhost:8080/receipts/{id}/points` to retrieve points for a particular receipt ID.

### Stopping the Service

To stop the Docker container:

1. Find the container ID using:

```bash
docker ps
```

2. Stop the container using the container ID:

```bash
docker stop <container_id>
```

## API Usage Example

### Example Request

You can use the below example to test the `Process Receipt` endpoint:

```json
{
    "retailer": "M&M Corner Market",
        "purchaseDate": "2022-03-20",
        "purchaseTime": "14:33",
        "items": [
        {
            "shortDescription": "Gatorade",
            "price": "2.25"
        },{
            "shortDescription": "Gatorade",
            "price": "2.25"
        },{
            "shortDescription": "Gatorade",
            "price": "2.25"
        },{
            "shortDescription": "Gatorade",
            "price": "2.25"
        }
        ],
        "total": "9.00"
}
```

### Expected Response

For the given request, you should expect a total of 109 points:

- 50 points - total is a round dollar amount
- 25 points - total is a multiple of 0.25
- 14 points - retailer name (M&M Corner Market) has 14 alphanumeric characters (note: '&' is not alphanumeric)
    - 10 points - 2:33pm is between 2:00pm and 4:00pm
- 10 points - 4 items (2 pairs @ 5 points each)

    The total comes up to 109 points.



