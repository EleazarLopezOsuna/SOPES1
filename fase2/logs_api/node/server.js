const express = require("express");
const http = require("http");
const socketIo = require("socket.io");
require('dotenv').config();
const redis = require("redis");

const portListen =  5000;
const index = require("./routes/index");

const app = express();
app.use(index);
app.use(function (req, res, next) {
    res.setHeader('Access-Control-Allow-Origin', '*');
    res.setHeader('Access-Control-Allow-Methods', 'GET, POST, OPTIONS, PUT, PATCH, DELETE');
    res.setHeader('Access-Control-Allow-Headers', 'X-Requested-With,content-type');
    res.setHeader('Access-Control-Allow-Credentials', false);
    next();
});

const server = http.createServer(app);

const io = require('socket.io')(server, {
    cors: {
      origin: '*',
      methods: ["GET", "POST"]

    }
  });

let interval;

io.on("connection", (socket) => {
    console.log("New client connected");

    if (interval) {
        clearInterval(interval);
    }
    interval = setInterval(() => getApiAndEmit(socket), 4000);
    
    socket.on("disconnect", () => {
        console.log("Client disconnected");
        clearInterval(interval);
    });
});

async function emitRedis (){

    const host     = `${process.env.HOSTREDIS}`
    const port     = `${process.env.PORTDB}`
    let res
    try{
        var client = redis.createClient(
            {
                url: "redis://"+host+":"+port
            }
        );
        (async () => {
            client.on('error', (err) => {
              console.log('Redis Client Error', err);
            });
            client.on('ready', () => console.log('Redis is ready'));
            await client.connect();
        })();
        
        if (res) { res = null}
        res =  await client.lRange("logskafka", 0, -1);
        
    } catch {
        try{
            await client.quit();
        } catch {}
    }finally{
        if (client){
            try {
                await client.quit();
            }catch (err){
            }
        }
        
    }

    return res
}

async function emitTidis (){

    const host     = `${process.env.HOSTTIDIS}`
    const port     = `${process.env.PORTDB}`
    let res
    try{
        var client = redis.createClient(
            {
                url: "redis://"+host+":"+port
            }
        );
        (async () => {
            client.on('error', (err) => {
              console.log('Tidis Client Error', err);
            });
            client.on('ready', () => console.log('Tidis is ready'));
            await client.connect();
        })();
        
        if (res) { res = null}
        res =  await client.lRange("logskafka", 0, -1);
        
    } catch {
        try{
            await client.quit();
        } catch {}
    }finally{
        if (client){
            try {
                await client.quit();
            }catch (err){
            }
        }
        
    }

    return res
}

const getApiAndEmit = async socket => {

    socket.emit("redis", await emitRedis());
    socket.emit("tidis", await emitTidis());
};

server.listen(portListen, () => console.log(`Listening on port ${portListen}`));