const mysql = require('mysql2');

const connection = mysql.createConnection({
  host: 'localhost',
  user: 'root',
  password: 'password'
  database: 'test'
});

connection.query(`CREATE TABLE IF NOT EXISTS contacts(id NUMBER, name VARCHAR2(50), email VARCHAR2(100) )`
  ,
  function(err) {
  if(err){
      console.log("Error!");
  }
  }
);
