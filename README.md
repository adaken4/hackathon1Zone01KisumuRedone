# Peer2Peer Lending on Blockchain and Smart Contracts

## Overview

This project is a Peer2Peer Lending platform built on a blockchain network using smart contracts. The goal is to create a decentralized lending system where users can lend and borrow funds directly from each other, leveraging blockchain technology for security and transparency.

## Features

- **Decentralized Lending**: Facilitate direct lending and borrowing between users.
- **Smart Contracts**: Manage lending terms, repayments, and enforcement through blockchain-based contracts.
- **User Authentication**: Secure sign-in and verification using blockchain identities.
- **Loan Terms Management**: Flexible loan terms including interest rates, durations, and collateral.
- **Payment Tracking**: Real-time tracking of loan repayments and outstanding balances.
- **Reputation System**: User ratings and feedback to build trust within the platform.

## Technologies Used

- **Blockchain Platform**: [Ethereum/BSC/Other]
- **Smart Contracts**: [Solidity] (Planned for future implementation)
- **Backend**: [Go]
- **Frontend**: [HTML/CSS] (JavaScript planned for future enhancement)
- **Database**: [PostgreSQL/MongoDB/Other]
- **Wallet Integration**: [MetaMask/Other]

## Roadmap

### Phase 1: Backend Development (Current)

1. **Set Up Go Backend**
   - Develop RESTful APIs to handle user interactions and blockchain communication.
   - Implement logic for managing loan requests, repayments, and user data.

2. **Blockchain Integration**
   - Use Go to interact with smart contracts deployed on a blockchain network.
   - Leverage libraries like `web3.go` for Ethereum interactions.

### Phase 2: Frontend Development (Current)

1. **Build UI with HTML/CSS**
   - Create and customize HTML/CSS templates for the user interface.
   - Ensure responsiveness and basic styling.

2. **Basic Interactivity**
   - Implement minimal JavaScript for dynamic content and basic user interactions (planned for next phase).

### Phase 3: Smart Contracts (Future)

1. **Smart Contract Development**
   - Write and deploy smart contracts using Solidity to handle lending logic and enforce terms.
   - Use resources like the [Solidity documentation](https://docs.soliditylang.org/) and [Remix IDE](https://remix.ethereum.org/) for development.

2. **Integration with Backend**
   - Update Go backend to interact with the deployed smart contracts.

### Phase 4: JavaScript Integration (Future)

1. **Frontend Enhancements**
   - Add advanced functionality and interactivity using JavaScript.
   - Enhance user experience with dynamic updates and client-side logic.

2. **Client-Side Interaction**
   - Use JavaScript to handle interactions with the backend API and smart contracts.

## Getting Started

### Prerequisites

- [Go](https://golang.org/) (for backend development)
- [HTML/CSS](https://developer.mozilla.org/en-US/docs/Web) (for frontend development)
- [MetaMask](https://metamask.io/) or other Ethereum wallet
- [Ganache](https://www.trufflesuite.com/ganache) (for local blockchain simulation)

### Installation

1. **Clone the Repository**

    ```bash
    git clone https://github.com/adaken4/peerfund-p2p-lending.git
    cd peer2peer-lending
    ```

2. **Install Dependencies**

    - **Backend**

      ```bash
      cd backend
      go mod download
      ```

    - **Frontend**

      ```bash
      cd frontend
      # Ensure HTML/CSS files are correctly set up
      ```

3. **Set Up Environment Variables**

    Create a `.env` file in the root directory and add the following variables:

    ```
    INFURA_PROJECT_ID=your_infura_project_id
    PRIVATE_KEY=your_private_key
    ```

4. **Deploy Smart Contracts** (Planned for future)

    Follow Solidity tutorials and use Remix IDE for deployment.

5. **Start the Backend Server**

    ```bash
    cd backend
    go run main.go
    ```

6. **Start the Frontend Application**

    ```bash
    cd frontend
    # Serve static files using a simple HTTP server or integrate with Go backend
    ```

### Usage

1. **Create an Account**: Sign up and create a blockchain wallet using MetaMask.
2. **Deposit Funds**: Deposit cryptocurrency into your wallet to start lending or borrowing.
3. **Create Loan Request**: Set loan terms and request funds.
4. **Lend Funds**: Review and lend funds to other users based on their loan requests.
5. **Manage Loans**: Track loan status, repayments, and balances through the platform dashboard.

### Features to Implement

- **Advanced Security Features**: Enhance smart contracts with additional security measures.
- **Frontend Improvements**: Add interactivity and dynamic content using JavaScript.
- **Additional Blockchain Integration**: Explore integration with other blockchain platforms.
- **Reputation and Feedback System**: Develop and integrate a robust reputation system.

## Contributing

We welcome contributions to improve the platform. To contribute:

1. **Fork the Repository**
2. **Create a New Branch**
3. **Commit Your Changes**
4. **Push to Your Fork**
5. **Submit a Pull Request**

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contact

For any questions or feedback, please reach out to:

- **Kennedy Ada**: adakennedy6@gmail.com
