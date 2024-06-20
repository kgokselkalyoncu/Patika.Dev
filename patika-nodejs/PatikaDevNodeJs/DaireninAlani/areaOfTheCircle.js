'use strict';

const radius = process.argv.slice(2);
var area = Math.PI * Math.pow(radius, 2);
console.log(`Yaricapi ${radius} olan dairenin alani: ${area}`);