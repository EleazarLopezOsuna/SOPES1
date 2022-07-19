// Funcion que genera los numeros random
const generateRandom = (player_random) => {
    return parseInt(Math.random() * ((player_random + 1) - 1) + 1);
}

// Algoritmo  del juego ID 1, ganador el jugador random par
const game1 = (players) => {
    let player_random = generateRandom(players);

    if(player_random == 1){
        return 2;
    }
    if((player_random % 2) == 0){
        return player_random;
    }else{
        return (player_random + 1);
    }
}

// Algoritmo  del juego ID 2, ganador el jugador random impar
const game2 = (players) => {
    let player_random = generateRandom(players);

    if(player_random == 1){
        return 1;
    }
    if((player_random % 2) == 0){
        return (player_random - 1);
    }else{
        return player_random;
    }
}

// Algoritmo  del juego ID 3, ganador el jugador 1
const game3 = () => {
    return 1;
}

// Algoritmo  del juego ID 4, ganador el jugador de enmedio
const game4 = (players) => {
    return (players / 2).toFixed();
}

// Algoritmo  del juego ID 5, ganador el ultimo jugador
const game5 = (players) => {
    return players;
}

// Variable a exportar
const games =  {
    game1,
    game2,
    game3,
    game4,
    game5
}

module.exports = games;