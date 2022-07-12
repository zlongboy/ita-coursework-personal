import path from 'path';
import express from 'express';
import { rootDir } from '../util/path.js';

const router = express.Router();
const __dirname = path.dirname(rootDir);

router.get('/users', (req, res) => {
    res.sendFile(path.join(__dirname, 'views', 'users.html'));
});

router.post('/users', (req, res) => {
    res.redirect('/')
});

export { router };