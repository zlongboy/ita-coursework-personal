const express = require('express');

const recordsController = require('../controllers/main');

const router = express.Router();

router.get('/', recordsController.getSearchBooks);
router.post('/', recordsController.postSearchBooks);

module.exports = router;