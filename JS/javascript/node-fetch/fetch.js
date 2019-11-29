// ref: https://developer.mozilla.org/ja/docs/Web/API/Fetch_API/Using_Fetch
// ref: https://morizyun.github.io/javascript/node-js-npm-library-node-fetch.html
// ref: https://qiita.com/ryohji/items/93f5050b9af6fc15693c

const fetch = require('node-fetch')

const url = 'http://localhost:8081/fetch';

fetch(url) 
  .then(function(resp) {
    return resp.json();
  })
  .then(function(myJ) {
    console.log(JSON.stringify(myJ));
  });
