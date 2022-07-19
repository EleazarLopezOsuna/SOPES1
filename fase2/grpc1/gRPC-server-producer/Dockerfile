FROM node:alpine

WORKDIR /app

COPY ["package.json", "package-lock.json*", "./"]

RUN npm install

COPY . .

EXPOSE 50051

CMD [ "node", "gRPC-server.js" ]