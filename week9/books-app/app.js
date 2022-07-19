const express = require('express');
const bodyParser = require('body-parser');

const db = require('./util/database');
const errorController = require('./controllers/errors');
const getBooks = require('./integrations/books')

const app = express();
const port = 8080;

app.set('view engine', 'pug');
app.set('views', 'views');
app.use(bodyParser.urlencoded({ extended: true }));
//Load public files here (e.g. CSS - video 75)

const addRoutes = require('./routes/add');
const displayRoutes = require('./routes/display');

getBooks('michael-lewis')

app.use(addRoutes);
app.use(displayRoutes);

app.use(errorController.get404);

app.listen(port);

