import express from 'express';

import { getAllRecords } from '../controllers/main.js';

export const router = express.Router();

router.get('/', getAllRecords);

