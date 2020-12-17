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
        fcn : "storeCode"
    }

    fabric_handler.submitTransaction(txnData).then((result) => {
        console.log(result)
        res.status(200).end();
    })
})

app.get('/storecode', (req, res) => {
    console.log('####...... storecode');

    let spentityid = req.body.spentityid
    let idpentityid = req.body.idpentityid
    let spCode = req.body.spcode;
    let idpCode = req.body.idpcode;
    let spCheck = req.body.spcheck;
    let idpCheck = req.body.idpcheck;
    let author = req.body.author

    let txnData = {
        fcn : "storeCode",
        args: [spentityid, idpentityid, spCode, idpCode, spCheck, idpCheck, author]
    }

    fabric_handler.submitTransaction(txnData).then((result) => {
        console.log(result)
        res.status(200).send(result);
    })
})

app.get('/approval', (req, res) => {
    console.log('####...... approval');

    let author = req.query.author;

    let txnData = {
        fcn : "storeCode",
        args: [author]
    }

    fabric_handler.submitTransaction(txnData).then((result) => {
        console.log(result)
        res.status(200).send(result);
    })
})

///////////////////////////////////////
startServer();
///////////////////////////////////////