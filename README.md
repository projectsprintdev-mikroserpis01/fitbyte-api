# Fitbyte API

![Go Version](https://img.shields.io/badge/Go-1.23+-00ADD8?style=flat&logo=go)

## Description

The application is built on [Go v1.23.4](https://tip.golang.org/doc/go1.22) and [PostgreSQL](https://www.postgresql.org/). It uses [Fiber](https://docs.gofiber.io/) as the HTTP framework and [pgx](https://github.com/jackc/pgx) as the driver and [sqlx](github.com/jmoiron/sqlx) as the query builder.

## Getting started

1. Ensure you have [Go](https://go.dev/dl/) 1.23 or higher and [Task](https://taskfile.dev/installation/) installed on your machine:

   ```bash
   go version && task --version
   ```

2. Create a copy of the `.env.example` file and rename it to `.env`:

   ```bash
   cp ./config/.env.example ./config/.env
   ```

   Update configuration values as needed.

3. Install all dependencies, run docker compose, create database schema, and run database migrations:

   ```bash
   task
   ```

4. Run the project in development mode:

   ```bash
   task dev
   ```

## Running the Application with Docker Compose

1. Start the Docker containers:

   ```sh
   task service:up
   ```

2. To stop the Docker containers:

   ```sh
   task service:down
   ```

## Checking Database Connection

To check the database connection, use the following command:

```sh
task service:db:connect
```

## Running Load Tests

1. Navigate to the folder test and clone the repository:

   ```sh
   git clone https://github.com/ProjectSprint/Batch3Project2TestCase.git
   ```

2. Install k6 (if you don't have it installed):

   Follow the instructions on the [k6 installation page](https://k6.io/docs/getting-started/installation/) to install k6 on your machine.

3. Navigate to the folder where this is extracted/cloned in terminal and run this for :

   ```sh
   BASE_URL=http://localhost:8080 make pull-test
   ```

4. Make sure that you have Redis installed and exposed on port 6379, then run:

   ```sh
   BASE_URL=http://localhost:8080 k6 run load_test.js
   ```


## Installing to Production

Before that, run VPN first:
```
sudo openvpn --config /path/to/your/config.ovpn
```
Replace `/path/to/your/config.ovpn` with the path to the `.ovpn` configuration file you received.

1. **Update Go modules**:

   Before building the production binary, make sure to update the Go modules by running:
   ```bash
   task
   ```

2. **Build the application for production**:

   Run the following command to build the production binary:
   ```bash
   task build
   ```

3. **Upload the binary to your EC2 instance using SCP**:

   Upload the built binary to your EC2 instance:
   ```bash
   scp -i /path/to/your-key.pem mybinary ubuntu@<EC2_PUBLIC_IP>:/home/ubuntu/
   ```

   Replace `/path/to/your-key.pem` with the path to your private key, `mybinary` with the name of your binary, and `<EC2_PUBLIC_IP>` with the public IP of your EC2 instance.

4. **Upload the config/.env to your EC2 instance using SCP**:

   Upload the `.env` configuration file (along with the `config/` directory) to your EC2 instance:
   ```bash
   scp -i /path/to/your-key.pem -r config ubuntu@<EC2_PUBLIC_IP>:/home/ubuntu/
   ```

   Replace `/path/to/your-key.pem` with the path to your private key and `<EC2_PUBLIC_IP>` with the public IP of your EC2 instance.

5. **Login to your EC2 instance**:

   SSH into your EC2 instance:
   ```bash
   ssh -i /path/to/your-key.pem ubuntu@<EC2_PUBLIC_IP>
   ```

6. **Make the binary executable**:

   If the binary isn't executable, run the following command to make it executable:
   ```bash
   chmod +x /home/ubuntu/mybinary
   ```

7. **Run the binary**:

   Run the binary to start the application:
   ```bash
   ./mybinary
   ```