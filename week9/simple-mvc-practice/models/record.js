import fs from 'fs';
import path from 'path';
import { __dirname, __filename } from '../app.js';

//Can refactor here if desired using a helper function to get path/readfile (video 103)

export class Record {
    constructor(name) {
        this.title = name; 
    }

    save() {
        const p = path.join(__dirname, 'data', 'records.json');
    
        fs.readFile(p, (err, fileContent) => {
            let records = [];
            if (!err) {
                records = JSON.parse(fileContent);
            }
            records.push(this);
            fs.writeFile(p, JSON.stringify(records), (err) => {
                console.log(err);
            });
        });
    }

    static fetchRecords(callback) {
        const p = path.join(__dirname, 'data', 'records.json');

        fs.readFile(p, (err, fileContent) => {
            if (err) {
                callback([])
            }
            callback(JSON.parse(fileContent));
        })
    }
}