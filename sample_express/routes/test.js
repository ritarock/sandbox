const express = require('express');
const router = express.Router();

router.get('/', function(req, res, next) {
  res.render('test', { test: 'TestPage' });
});

module.exports = router;
