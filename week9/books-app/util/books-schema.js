const mysql = require('mysql2');

const pool = mysql.createPool({
    host: 'localhost',
    user: 'root',
    database: 'books',
    password: 'LegBeforeWicket#2022'
})

module.exports = pool.promise();