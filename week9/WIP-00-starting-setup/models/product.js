const fs = require('fs');
const { getDefaultSettings } = require('http2');
const path = require('path');
const rootDir = require('../util/path')

//const products = [];

module.exports = class Product {
    constructor(t) {
        this.title = t;
    }

    save() {
        //products.push(this);
        //const p = path.join(path.dirname(process.main.filename));
        //TODO: filename is busted, using manual path for now
        const p = path.join(__dirname,'../', 'data', 'products.json');
        fs.readFile(p, (err, fileContent) => {
            let products = []
            if (!err) {
                products = JSON.parse(fileContent);
            }
            products.push(this);
            //need to use arrow function here or this refers to the function not the class.
            fs.writeFile(p, JSON.stringify(products), (err) => {
                console.log(err);
            });
        });
    }

    static fetchAll(cb) {
        const p = path.join(__dirname,'../', 'data', 'products.json');
        fs.readFile(p, (err, fileContent) => {
            if (err) {
                cb([])
            }
            cb(JSON.parse(JSON.stringify(fileContent)));
        });
    }
};