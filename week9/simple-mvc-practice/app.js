const express = require('express');
const path = require('path');
const bodyParser = require('body-parser');

const app = express();
const port = 3000;

app.set('view engine', 'pug');
app.set('views', 'views');
app.use(bodyParser.urlencoded({ extended: true }));
//Load public files here (e.g. CSS - video 75)

const db = require('./util/database'); 

db.execute('SELECT * FROM records LIMIT 5')
    .then(result => {
        console.log(result);
    })
    .catch(err => {
        console.log(err);
    });

const addRoutes = require('./routes/add');
const displayRoutes = require('./routes/display');
const errorController = require('./controllers/errors');

app.use(addRoutes);
app.use(displayRoutes);

app.use(errorController.get404);

app.listen(port);

