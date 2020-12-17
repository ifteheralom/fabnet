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


app.get('/', (req, res) => {
    console.log('####......New fabric request');

    // let txnData = {
    //     fcn : "storeCode"
    // }

    // fabric_handler.submitTransaction(txnData).then((result) => {
    //     console.log(result)
    //     res.status(200).end();
    // })

    res.status(200).send("Fabric DApp is Running");
})

//////////////////////////////////////////////////////////
app.post('/storetallist', (req, res) => {
    console.log('####...... storetallist');

    let entityId = req.body.entityId
    let tal = req.body.tal

    let txnData = {
        fcn : "StoreTalList",
        args: [entityId, tal]
    }

    fabric_handler.submitTransaction(txnData).then((result) => {
        console.log(result)
        res.status(200).send("success");
    })
})

app.get('/tallistfetch', (req, res) => {
    console.log('####...... tallistfetch');

    let entityId = req.query.entityId

    let txnData = {
        fcn : "TalListReturn",
        args: [entityId]
    }

    fabric_handler.submitTransaction(txnData).then((result) => {
        console.log(result.toString())
        let resultJson = JSON.parse(result.toString())
        let arr = []
        for (var i in resultJson)
            {
                var name = resultJson[i].Tal;
                arr.push(name)
            }
        res.status(200).send(arr)
    })
})


app.post('/storecode', (req, res) => {
    console.log('####...... storecode');

    let spentityid = req.body.spentityid
    let idpentityid = req.body.idpentityid
    let spCode = req.body.spcode;
    let idpCode = req.body.idpcode;
    let spCheck = req.body.spcheck;
    let idpCheck = req.body.idpcheck;
    let author = req.body.author

    let txnData = {
        fcn : "NewStoreCode",
        args: [spentityid, idpentityid, spCode, idpCode, spCheck, idpCheck, author]
    }

    fabric_handler.submitTransaction(txnData).then((result) => {
        console.log("success")
        res.status(200).send("success");
    })
})

app.get('/codefetch', (req, res) => {
    console.log('####...... codefetch');
    
    let spentityid = req.query.spentityid;
    let idpentityid = req.query.idpentityid;
    let author = req.query.author;
    let code = req.query.code;

    let txnData = {
        fcn : "NewCode",
        args: [spentityid, idpentityid, author, code]
    }

    fabric_handler.submitTransaction(txnData).then((result) => {
        console.log(result)
        res.status(200).send(JSON.stringify(result.toString()));
    })
})

app.get('/approval', (req, res) => {
    console.log('####...... approval');
    
    let author = req.query.author;

    let txnData = {
        fcn : "Approval",
        args: [author]
    }

    fabric_handler.submitTransaction(txnData).then((result) => {
        console.log(result)
        res.status(200).send(result.toString());
    })
})

app.get('/removeapproval', (req, res) => {
    console.log('####...... removeapproval');
    
    let spentityid = req.query.spentityid;
    let idpentityid = req.query.idpentityid;

    let txnData = {
        fcn : "RemoveApproval",
        args: [spentityid, idpentityid]
    }

    fabric_handler.submitTransaction(txnData).then((result) => {
        console.log(result)
        res.status(200).send("success");
    })
})

app.get('/deletetal', (req, res) => {
    console.log('####...... deletetal');
    
    let entityid = req.query.entityid
    let tal = req.query.tal

    let txnData = {
        fcn : "TalListDelete",
        args: [entityid, tal]
    }

    fabric_handler.submitTransaction(txnData).then((result) => {
        console.log(result)
        res.status(200).send("success");
    })
})

///////////////////////////////////////
startServer();
///////////////////////////////////////