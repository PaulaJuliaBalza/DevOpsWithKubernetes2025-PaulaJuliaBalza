const uuid = require('uuid');
const randomString = uuid.v4();

setInterval(() => {
  console.log(`${new Date().toISOString()}: ${randomString}`);
}, 5000);