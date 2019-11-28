const superagent = require('superagent');

const getVersionURL = process.env.API_SERVER + "/api/getVersion";

console.log("Starting to poll " + getVersionURL);

setInterval(() => {
    superagent.get(getVersionURL).end((err,res)=>{
        try {
            console.log(res.text);
        } catch (err) {
            console.error('Error');
        }
    })
}, 1000);