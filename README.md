# Request Limiter with Redis Middleware for Go Fiber

This middleware for Go Fiber enables request rate limiting based on client's IP address using Redis as a backend storage. It's useful for limiting the number of requests from a client within a specified time window.

## Installation

To use this middleware, follow the steps below:

1. **Clone the Repository**: 
   ```sh
   git clone https://github.com/bishnudev1/request_limiter.git

2. Create a .env File:
Create a file named .env in the root directory of your project and define your Redis connection details:
    ```sh
     REDIS_ADDR=redis://<REDIS_HOST>:<REDIS_PORT>
     REDIS_PASSWORD=<REDIS_PASSWORD>

3. Install Dependencies:
Use go get to install the required dependencies:
    ```sh
    go get

4. In the routes/product_route.go define your maxRequest and timeOut
    ```sh
    middlewares.RateLimitMiddleware(timeOut, maxRequest,)

5. Test your own Request Limiter API with
     ```
     go run server.go



## Feel free to contribute more features. I'll convert it into a package soon so that we can use it with any project with ease.
