const express = require('express');

const recordsController = require('../controllers/main');

const router = express.Router();

router.get('/add-record', recordsController.getAddRecords);
router.post('/add-record', recordsController.postAddRecords);
router.get('/add-book', recordsController.searchAuthor);

module.exports = router;
