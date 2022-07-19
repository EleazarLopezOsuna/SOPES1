FROM node:16-alpine as nodey
WORKDIR /app
COPY . .
RUN yarn install
ENV REACT_APP_IP_BACKONE http://35.212.231.127:80/
ENV REACT_APP_IP_BACKTWO http://35.212.231.127:80/
ENV REACT_APP_IP_SOCKET https://tallerdocker-341423.uc.r.appspot.com/
RUN yarn build

### STAGE 2: Run ###
FROM nginx:1.21.6-alpine
COPY --from=nodey /app/build/ /usr/share/nginx/html
EXPOSE 80
