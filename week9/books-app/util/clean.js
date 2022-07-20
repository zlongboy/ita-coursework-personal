const Book = require('../models/book');

exports.author = (input) => {
    return input.toLowerCase().replaceAll(' ', '-');
};

exports.volumes = (volumes) => {
    const results = [];
    volumes.forEach(e => {
        const volume = new Book(e.id, e.volumeInfo.title, e.volumeInfo.authors[0], e.volumeInfo.publisher)
        results.push(this);
    });
    console.log(results);
};
