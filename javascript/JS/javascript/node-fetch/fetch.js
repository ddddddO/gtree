// ref: https://developer.mozilla.org/ja/docs/Web/API/Fetch_API/Using_Fetch
// ref: https://morizyun.github.io/javascript/node-js-npm-library-node-fetch.html
// ref: https://qiita.com/ryohji/items/93f5050b9af6fc15693c

const fetch = require('node-fetch')

const url = 'http://localhost:8081/fetch';

async function fetchData() {
  const rslt = await fetch(url) 
    .then(function(resp) {
      let tmp = resp.json();
      // console.log(tmp);
      return tmp;
    })
    .then(function(myJ) {
      return JSON.stringify(myJ);
    });


  parseJSON(rslt);
}

function parseJSON(json) {
  const parsed = JSON.parse(json)

  // 以下でJSONキーの配列取得
  console.log(Object.keys(parsed))
  // 以下でJSON値の配列取得
  console.log(Object.values(parsed))
}

fetchData();
