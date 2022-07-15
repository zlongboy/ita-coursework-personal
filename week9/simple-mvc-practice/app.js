import express from 'express';
import path from 'path';
import bodyParser from 'body-parser';

const app = express();
const port = 3000;
app.set('view engine', 'pug');
app.set('views', 'views');

//add routes
app.get('/', (req, res, next) => {
    res.render('all-records');
});

app.get('/add-record', (req, res, next) => {
    res.render('add-record');
});

app.listen(port);