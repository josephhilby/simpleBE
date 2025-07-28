# SimpleBE

<!-- Overview -->

## Project Overview

This repository contains a simple API in Golang. It is intended to serve as a template to build upon as I become more proficient with this language.

The Frontend to this project is SimpleFE. It is a Next.js app.

<!-- SETUP -->

## Project Setup & Initialization

### Prerequisites

Ensure the following tools are installed:

* Docker `v28.1.1`, or later
* Docker Compose `v2.36.0`, or later

### Initial Setup

1. Clone your forked repository and navigate to the project root:
    ```bash
    git clone <your-forked-repo-url>
    cd simpleBE
    ```

2. Make the control script executable:
    ```bash
    chmod +x docker-control.sh
    ```

3. Build and start the containers:
    ```bash
    ./docker-control.sh build
    ```

<div align="center">

💡 Run `./docker-control.sh help` to view additional commands.

</div>


4. Visit the API at: <http://localhost:8080/api/hello>


<br />

<!-- API ENDPOINTS -->

