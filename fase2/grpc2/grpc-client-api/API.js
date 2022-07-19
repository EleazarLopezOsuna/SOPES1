require('dotenv').config();
const express = require('express');
const app = express();
let morgan = require('morgan');
let cors = require('cors');

const port = process.env.API_PORT || 3000;

app.use(express.json());
app.use(morgan('dev'));
app.use(cors());

app.use('/match', require('./routes/matches.js'));

app.listen(port, () => {
    console.log('Server in port: ', port);
})