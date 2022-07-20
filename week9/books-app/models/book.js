const db = require('../util/books');

//Can refactor here if desired using a helper function to get path/readfile (video 103)

module.exports = class Book {
    constructor(BID, t, a, p) {
        this.bookId = BID
        this.title = t
        this.author = a
        this.publisher = p
    }

    save() {
        return db.execute('INSERT INTO books (bookId, title, author, publisher) VALUES (?, ?, ?, ?)', [this.bookId, this.title, this.author, this.publisher]);
    }

    static fetchRecords() {
        return db.execute('SELECT * FROM books')
    }
};