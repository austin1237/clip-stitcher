FROM node:10.8.0-alpine

WORKDIR /app

COPY package*.json ./

RUN npm install

COPY . .


CMD [ "npm", "start"]

