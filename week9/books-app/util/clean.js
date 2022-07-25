const Book = require('../models/book');
const data = require('./data')

exports.author = (input) => {
    return input.toLowerCase().replaceAll(' ', '-').trim();
};

exports.volumes = (volumes) => {
    const results = [];
    volumes.forEach(e => {
        results.push(new Book(
            e.id, 
            e.volumeInfo.title, 
            e.volumeInfo.authors ? e.volumeInfo.authors[0] : '', 
            e.volumeInfo.publisher ? e.volumeInfo.publisher : '', 
            e.accessInfo.country,
            e.volumeInfo.publishedDate ? e.volumeInfo.publishedDate.substring(0, 4) : '',
            e.saleInfo.retailPrice ? e.saleInfo.retailPrice.amount : 0.00,   
        ));
    });
    
    let sql = results.map(el => `("${el.bookId}", "${el.title}", "${el.author}", "${el.publisher}", "${el.country}", "${el.year}", ${el.price})`);
    
    data.saveBooks(sql);
    
    return results
};
