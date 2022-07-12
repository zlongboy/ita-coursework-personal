import path from 'path';
import express, { application } from 'express';

import { __dirname } from './app.js'

const router = express.Router();

router.get('/', (req, res, next) => {
    //res.send('<h1>You bet this is the homepage!</h1>');
    res.sendFile(path.join(__dirname, 'views', 'shop.html'));
});

export { router }; 