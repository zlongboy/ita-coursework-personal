const mysql = require('mysql2');
const booksDB = require('../config');

const pool = mysql.createPool({
    host: 'localhost',
    user: 'root',
    database: 'books',
    password: booksDB.SECRET
})

module.exports = pool.promise();