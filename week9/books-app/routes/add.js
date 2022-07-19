const express = require('express');

const recordsController = require('../controllers/main');

const router = express.Router();

router.get('/add-record', recordsController.getAddRecords);
router.post('/add-record', recordsController.postAddRecords);

module.exports = router;
