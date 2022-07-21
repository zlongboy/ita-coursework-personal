const express = require('express');

const recordsController = require('../controllers/main');

const router = express.Router();

router.get('/', recordsController.getAllBooks);

module.exports = router;