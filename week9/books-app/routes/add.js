const express = require('express');

const recordsController = require('../controllers/main');

const router = express.Router();

router.get('/add-author', recordsController.getAuthor);
router.post('/add-author', recordsController.postAuthor);

module.exports = router;
