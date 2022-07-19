var PROTO_PATH = './proto/fase2.proto';

var grpc = require('@grpc/grpc-js');
var protoLoader = require('@grpc/proto-loader');
var packageDefinition = protoLoader.loadSync(
  PROTO_PATH,
  {keepCase: true,
    longs: String,
    enums: String,
    defaults: true,
    oneofs: true
  });
  var fase2_proto = grpc.loadPackageDefinition(packageDefinition).proto;
  
  var games = require("./juegos");

  const dotenv = require('dotenv').config();
  var amqp = require("amqplib/callback_api");
  const settingsRabbit = {
    "protocolo": process.env.Prot_Rabbit,
    "hostname":  process.env.Host_Rabbit,
    "port": process.env.Port_Rabbit
  }
  
/**
 * Implements the SenConfigurationPlay RPC method.
 */
function SenConfigurationPlay(call, callback) {
    let {gameId, player} = call.request;
    let {game1,game2,game3,game4,game5} = games;
    let ganador = 0;
    let msj = ` Game ID ${gameId} con ${player} jugadores`;
    let gameName = "";

    switch(gameId) {
      case 1:
        ganador = game1(player);
        gameName = "Game 1";
        break;
      case 2:
        ganador = game2(player);
        gameName = "Game 2";
        break;
      case 3:
        ganador = game3();
        gameName = "Game 3";
        break
      case 4:
        ganador = game4(player);
        gameName = "Game 4";
        break
      default:
        ganador = game5(player);
        gameName = "Game 5";
    }

    let msg = {
        "game_id":gameId.toString(),
        "players": player.toString(),
        "game_name": gameName,
        "winner": ganador.toString(),
        "queue": "RabbitMQ"
    };

    console.log(">> gRPC-Server:" + msj + ` ganador: juador ${ganador}`);

    RabbitMQ(msg);
    
    callback(null, {message: msj});
}


/*************** RabbitMQ ****************/
const RabbitMQ = (msg) => {
  amqp.connect(`${settingsRabbit.protocolo}://${settingsRabbit.hostname}:${settingsRabbit.port}`, function(error0, connection) {

    if (error0) {
      throw error0;
    }

    connection.createChannel(function(error1, channel) {
      if (error1) {
        throw error1;
      }
      var queue = 'fase2';
  
      channel.assertQueue(queue, {
        durable: false
      });
  
      channel.sendToQueue(queue, Buffer.from(JSON.stringify(msg)));
      console.log(" [x] Sent", JSON.stringify(msg));
    });

    setTimeout(function() {
      connection.close();
    }, 500);

  });
}

/**
 * Starts an RPC server that receives requests for the RunPlay service at the
 * sample server port
 */
function main() {
  const urlServer = process.env.Host_grpcServer + ":" + process.env.Port_grpcServer
  var server = new grpc.Server();
  server.addService(fase2_proto.RunPlay.service, {SenConfigurationPlay: SenConfigurationPlay});
  server.bindAsync(urlServer, grpc.ServerCredentials.createInsecure(), () => {
    server.start();
    console.log(">> gRPC-Server: on port", process.env.Port_grpcServer);
  });
}

main();