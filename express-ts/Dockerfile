FROM node:alpine AS build

WORKDIR /app

COPY package.json yarn.lock ./
RUN yarn install 

COPY . ./
CMD ["yarn", "run", "dev"]