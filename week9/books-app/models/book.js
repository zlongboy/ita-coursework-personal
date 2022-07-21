const db = require('../util/books-schema');

//Can refactor here if desired using a helper function to get path/readfile (video 103)

module.exports = class Book {
    constructor(BID, t, a, p, c) {
        this.bookId = BID
        this.title = t
        this.author = a
        this.publisher = p
        this.country = c
    }

    save() {
        return db.execute('INSERT INTO books (bookId, title, author, publisher, country) VALUES (?, ?, ?, ?, ?)', [this.bookId, this.title, this.author, this.publisher, this.country]);
    }

    static fetchRecords() {
        return db.execute('SELECT * FROM books')
    }
};