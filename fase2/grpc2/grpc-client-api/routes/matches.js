const express = require('express');
let router = express.Router();
const client = require('../gRPC-Client-API')

router.post('/addMatch', function(req, res){
    const data_match = {
        gameId: req.body.game_id,
        numberPlayers: req.body.players
    }

    client.AddMatch(data_match, function(err, response){
        res.status(200).json({message: response.message});
    })
})

router.get('/hola', function(req, res){
    res.status(200).json({message: "Hola"});
})

module.exports = router;