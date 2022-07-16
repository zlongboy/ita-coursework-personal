import express from 'express';

import { getAddRecords, postAddRecords } from '../controllers/main.js';

export const router = express.Router();

router.get('/add-record', getAddRecords);
router.post('/add-record', postAddRecords);


