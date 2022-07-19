require('dotenv').config();
let parseArgs = require('minimist');
let PROTO_PATH = './proto/squidgame.proto';
let grpc = require('@grpc/grpc-js');
let protoLoader = require('@grpc/proto-loader');
let packageDefinition = protoLoader.loadSync(
    PROTO_PATH,
    {keepCase: true,
        longs: String,
        enums: String,
        defaults: true,
        oneofs: true
    });
let squidGame_proto = grpc.loadPackageDefinition(packageDefinition).squidgame;

let argv = parseArgs(process.argv.slice(2), {
    string: 'target'
});
let target;
if (argv.target) {
    target = argv.target;
} else {
    target = process.env.SERVER_IP + ':' + process.env.SERVER_PORT;
}
let client = new squidGame_proto.Matches(target,grpc.credentials.createInsecure());

module.exports = client;