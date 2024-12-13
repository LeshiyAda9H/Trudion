# Trudion

Trudion is an application designed to help users find study partners efficiently. It streamlines the process of connecting with like-minded learners to achieve academic goals together.

---

## Getting Started

Follow the instructions below to set up and run Trudion using Docker.

---

## Prerequisites

- Ensure you have Docker and Docker Compose installed on your system.
  For installation instructions, refer to the [official Docker documentation](https://docs.docker.com/get-docker/).

---

## Steps to Run the Application

1. **Navigate to the Application Directory**

   Open your terminal and move to the `./app` directory where the Docker setup files are located:

   ```bash
   cd ./app
   ```

2. **Build and Start the Application**

   Use the following command to build and start the application:

   ```bash
   docker compose up --build
   ```

   This command will:
    - Build the application using the Docker configuration.
    - Start the application and its dependencies.

---

## Notes

- For **Windows** and **macOS**, the above command might work as is if Docker Desktop is properly configured. If you encounter issues, refer to the [Docker Compose documentation](https://docs.docker.com/compose/).

- Ensure no other applications are using the same ports as Trudion.

---

## Troubleshooting

If you experience any issues:

1. Check the Docker logs:
   ```bash
   docker compose logs
   ```

2. Rebuild the application:
   ```bash
   docker compose up --build --force-recreate
   ```

3. If the issue persists, consult the application's documentation or contact support.

---

## Contribution

Feel free to contribute to Trudion by submitting pull requests or reporting issues in the repository.

---

## License

This project is licensed under the MIT License.
