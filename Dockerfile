FROM node:14.18.1-bullseye-slim
WORKDIR /app

COPY package*.json .
RUN npm install
COPY . .

CMD [ "node", "src/server.ts" ]
