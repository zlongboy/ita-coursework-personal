import express from 'express';
import path from 'path';
import { fileURLToPath } from 'url'
import bodyParser from 'body-parser';

//TODO: import pool from './util/database.js' 

const app = express();
const port = 3000;

export const __filename = fileURLToPath(import.meta.url);
export const __dirname = path.dirname(__filename);

app.set('view engine', 'pug');
app.set('views', 'views');
app.use(bodyParser.urlencoded({ extended: true }));
//Load public files here (e.g. CSS - video 75)

db.execute('SELECT * FROM records')
    .then(result => {
        console.log(result);
    })
    .catch(err => {
        console.log(err);
    });

import { router as addRoutes } from './routes/add.js';
import { router as displayRoutes } from './routes/display.js';
import { get404 as errorController } from './controllers/errors.js';

app.use(addRoutes);
app.use(displayRoutes);
app.use(errorController);

app.listen(port);

