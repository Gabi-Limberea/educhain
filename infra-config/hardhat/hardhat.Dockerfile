FROM node:24-slim

WORKDIR /hardhat

RUN npm install --save-dev "hardhat@^2.23.0"

CMD ["npx", "hardhat", "node"]

