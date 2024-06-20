const promise = new Promise((resolve, reject) => {
    resolve("VERILER ALINDI");
    reject("BAGLANTI HATASI");
});

//console.log(promise);

promise.then(value => { console.log(value); })
       .catch(error => { console.log(error); });