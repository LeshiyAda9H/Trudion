FROM node:latest

WORKDIR /app/ui

COPY package.json package-lock.json ./

RUN npm install

EXPOSE 5173

COPY . .