const express = require('express');
const router = express.Router();

router.get('/', function(req, res, next){
  res.render('sample', { sample: 'SAMPLE' });
});

module.exports = router;
