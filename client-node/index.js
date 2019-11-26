const superagent = require('superagent');

setInterval(() => {
    superagent.get('http://localhost:3000/api/getVersion').end((err,res)=>{
        try {
            console.log(res.text);
        } catch (err) {
            console.error('Error');
        }
    })
}, 1000);