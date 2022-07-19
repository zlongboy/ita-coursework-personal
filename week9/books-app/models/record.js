const fs = require('fs');

const db = require('../util/database');

//Can refactor here if desired using a helper function to get path/readfile (video 103)

module.exports = class Record {
    constructor(name) {
        this.title = name; 
    }

    save() {
        return db.execute('INSERT INTO records (title) VALUES (?)', [this.title]);
    }

    static fetchRecords(callback) {
        return db.execute('SELECT * FROM records')
    }
};