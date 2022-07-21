const e = require('express');
const Book = require('../models/book');

exports.author = (input) => {
    return input.toLowerCase().replaceAll(' ', '-').trim();
};

exports.volumes = (volumes) => {
    const results = [];
    volumes.forEach(e => {
        results.push(new Book(
            e.id, 
            e.volumeInfo.title, 
            e.volumeInfo.authors ? e.volumeInfo.authors[0] : null, 
            e.volumeInfo.publisher ? e.volumeInfo.publisher : null, 
            e.accessInfo.country));
    });
    console.log(results);
};