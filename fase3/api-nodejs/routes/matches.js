const express = require('express');
let router = express.Router();
const client = require('../gRPC-Client-API')

router.get('/', function(req, res){
    const data_match = {
        gameId: req.query.game_id,
        numberPlayers: req.query.players
    }

    client.AddMatch(data_match, function(err, response){
        res.status(200).json({message: response.message});
    })
})

router.get('/hola', function(req, res){
    res.status(200).json({message: "Hola"});
})

module.exports = router;