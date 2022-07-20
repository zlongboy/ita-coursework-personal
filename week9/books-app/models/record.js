const db = require('../util/records-schema');

//Can refactor here if desired using a helper function to get path/readfile (video 103)

module.exports = class Record {
    constructor(name) {
        this.title = name; 
    }

    save() {
        return db.execute('INSERT INTO records (title) VALUES (?)', [this.title]);
    }

    static fetchRecords() {
        return db.execute('SELECT * FROM records')
    }
};