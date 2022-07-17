import mysql from 'mysql';

const pool = mysql.createPool({
    host: 'localhost',
    user: 'root',
    database: 'simple-mvc-practice',
    //password: '';
});

export default pool.promise();