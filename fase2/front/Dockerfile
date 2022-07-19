FROM node:16-alpine as nodey
WORKDIR /app
COPY . .
RUN yarn install
ENV REACT_APP_IP_LOGS 34.125.51.22:8080
ENV REACT_APP_IP_SOCKET 34.125.51.22:5000
RUN yarn build

### STAGE 2: Run ###
FROM nginx:1.21.6-alpine
COPY --from=nodey /app/build/ /usr/share/nginx/html
EXPOSE 80
