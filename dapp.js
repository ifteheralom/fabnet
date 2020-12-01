const path = require('path');
const http = require('http');
// const hbs = require('hbs');
const socketIO = require('socket.io');
const express = require('express');
const bodyParser = require('body-parser');
const multer = require('multer');
const events = require('events');
const _ = require('lodash');
const yargs = require('yargs');

// const publicPath = path.join(__dirname + '/public');
const port = process.env.PORT || 3000;

const upload = multer();
const app = express();

// const eventEmitter = new events.EventEmitter();

// hbs.registerPartials(__dirname + '/views/partials');
// app.set('view engine', hbs);
// app.use(express.static(publicPath));

app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: true }));
app.use(upload.array());

let startServer = () => {
	let server = http.createServer(app);
	// let io = socketIO(server);
	server.listen(port, function () {
		console.log(`Server is up on: http://localhost:${server.address().port}`);
	});
}

// const argv = yargs.argv;
// var command = argv._[0];

const fabric_handler = require('./fabric-handler');


app.get('/fabric', (req, res) => {
    console.log('####......New fabric request');

    let txnData = {

    }

    fabric_handler.submitTransaction(txnData).then((result) => {
        console.log(result)
        res.status(200).end();
    })
})

///////////////////////////////////////
startServer();
///////////////////////////////////////