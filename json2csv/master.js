const Json2csvParser = require('json2csv').Parser;
const fs = require("fs");

async function main() {
 const csv = await getCsv();
 console.log(csv);
 // await outPut(csv);
}

async function getCsv() {
 const fields = ['jan', 'price', 'description'];

 const products = [
   {
     "jan": "001",
     "name": "span",
     "price": 2000,
     "description": "spam, spam"
   },
   {
     "jan": "102",
     "name": "hoge",
     "price": 2000,
     "description": "hoge, hoge"
   },
   {
     "jan": "201",
     "name": "fuga",
     "price": 3000,
     "description": "fuga, fuga"
   },
 ];
 const json2csvParser = new Json2csvParser({ fields, header: true});
 const report = json2csvParser.parse(products);
 retrun (report);
}

async function outPut(csv) {
 fs.writeFile('./out.csv', csv);
}

main();
