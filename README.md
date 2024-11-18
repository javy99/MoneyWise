# MoneyWise

## Project Overview

MoneyWise is a personal finance manager that helps users manage their income, expenses, and investments efficiently. The app integrates with bank APIs, allows users to categorize transactions, and provides insightful visualizations to track financial goals.

## Tech Stack

- **Backend**: Go (Microservices Architecture)
- **Frontend**: ReactJS with TypeScript
- **Database**: TBD
- **Containerization**: Docker
- **Orchestration**: Kubernetes
- **HELM**: For Kubernetes deployment
- **CI/CD**: GitHub actions for continuous integration and deployment.

## Features

- **User Authentication**: Secure login system with JWT tokens.
- **Transaction Management**: Track and categorize transactions.
- **Budgeting**: Set up budgets for various categories.
- **Insights**: Graphs and reports to visualize financial data.
- **Bank API Integration**: Automatic syncing of transactions with bank APIs.

## Getting Started

### Prerequisites

### Setup Instructions

1. Clone the repository:

````bash
git clone https://github.com/javy99/MoneyWise.git
```

2. Navigate to the backend directory:
```bash
cd MoneyWise/backend
```

3. Initialize the Go module:
```bash
go mod init MoneyWise
```

4. Run Docker containers (for local development):
```bash
docker-compose up
```

5. To run the backend locally, use:
```bash
go run cmd/service-name/main.go
```

6. For the frontend:
```bash
cd frontend
npm install
npm run dev
```

CI/CD
We use GitHub Actions for continuous integration and deployment. The workflows are configured to:
- Run tests on pull requests.
- Deploy the app to a staging environment upon merging to the develop branch
- Deploy to production upon merging to the main branch.

For more details on the CI/CD process, check the .github.workflows folder in the repository

## Developers Guidelines

Please follow the CONTRIBUTING.md file for detailed instructions on how to contribute, including the branching strategy, commit message convention, and pull request guidelines.
