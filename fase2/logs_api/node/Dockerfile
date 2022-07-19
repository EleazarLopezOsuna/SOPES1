FROM node:16-alpine
ENV NODE_ENV=production

WORKDIR /app

COPY . .

RUN yarn install --production
ENV HOSTREDIS 34.125.126.185
ENV HOSTTIDIS 34.125.195.218
ENV PORTDB 80

CMD [ "node", "server.js" ]
EXPOSE 5000