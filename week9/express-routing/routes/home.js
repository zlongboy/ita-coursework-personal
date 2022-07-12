import path from 'path';
import express from 'express';
import { rootDir } from '../util/path.js';

const router = express.Router();
const __dirname = path.dirname(rootDir);

router.get('/', (req, res) => {
    res.sendFile(path.join(__dirname, 'views', 'home.html'))
});

export { router };