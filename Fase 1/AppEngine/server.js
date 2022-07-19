const express = require("express");
const http = require("http");
const socketIo = require("socket.io");
const {MongoClient} = require("mongodb");

const portListen = process.env.PORT || 5000;
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
    }
  });

let interval;

io.on("connection", (socket) => {
    console.log("New client connected");
    if (interval) {
        clearInterval(interval);
    }
    interval = setInterval(() => getApiAndEmit(socket), 10000);
    socket.on("disconnect", () => {
        console.log("Client disconnected");
        clearInterval(interval);
    });
});

const usr      = "sistemasOperativos1"
const pwd      = "1234"
const host     = "34.67.195.168"
const port     = 27017
const database = "projectDatabase"

const uri = "mongodb://" + usr + ":" + pwd + "@" + host + ":" + port + "/admin"

const client = new MongoClient(uri)

async function getLogs(client, type){
    const cursor = await client.db(database).collection(database).find({
        '$or': [
            {
                'endpoint': type + '1'
            }, {
                'endpoint': type + '2'
            }
        ]
    })
    return await cursor.toArray()
}

const getApiAndEmit = async socket => {
    await client.connect()
    // Emitting a new message. Will be consumed by the client
    var n = await getLogs(client, '/ram');
    console.log( JSON.stringify(n))
    socket.emit("ramLog", n);
    socket.emit("procLog", await getLogs(client, '/procesos'));
};

server.listen(portListen, () => console.log(`Listening on port ${portListen}`));