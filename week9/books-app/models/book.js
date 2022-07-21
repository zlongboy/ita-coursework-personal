const db = require('../util/books-schema');

module.exports = class Book {
    constructor(BID, t, a, p, c, y, usd) {
        this.bookId = BID
        this.title = t
        this.author = a
        this.publisher = p
        this.country = c
        this.year = y
        this.price = usd
    }
};