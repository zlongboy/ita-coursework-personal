const db = require('../util/books-config');

exports.saveBooks = (valuesSQL) => {
    db.execute('INSERT IGNORE INTO books.volumes (Id, Title, Author, Publisher, Country, Year, Price) VALUES' + valuesSQL);
}

exports.fetchBook = async (bookId) => {
   const [rows, fields] = await db.execute(`SELECT * FROM books.volumes WHERE Id = '${bookId}'`);
   //console.log(rows);
   return rows;
}