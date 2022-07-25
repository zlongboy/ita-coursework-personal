const mysql = require('mysql2');
const config = require('../config');

const pool = mysql.createPool({
    host: 'localhost',
    user: 'root',
    database: 'books',
    //password: config.DB_SECRET
    password: 'LegBeforeWicket#2022'
})

module.exports = pool.promise();