const zlib = require('zlib');
const testLog = require('testLog');
const logAdapter = require('./adapters/logAdapter.js');
const logService = require('./services/logService.js');

// env vars
const appENV = process.env.APP_ENV;
const logDnaKey = process.env.LOG_DNA_KEY;

main = async (cloudWatchLog) => {
    // validate here
    let payload = logAdapter.transformCWToLD(cloudWatchLog);
    return logService.sendLogsOff(payload, logDnaKey);
}

exports.handler = async(event, context) => {
    if (appENV === "local"){
        console.log("local mode enabled")
        await main(testLog);
    }else{
        console.log("about to payload")
        let payload = new Buffer(event.awslogs.data, 'base64');
        console.log("about to gunzip");
        let crLog = await decompressEvent(payload);
        await main(crLog);
    }
}

let decompressEvent = (payload) =>{
    return new Promise((resolve, reject)=>{
        zlib.gunzip(payload, (err, result) =>{
            if (err){
                return reject(e);
            }
            result = JSON.parse(result.toString('ascii'));
            let resultStr = JSON.stringify(result, null, 2);
            console.log("Event Data:", resultStr);
            return resolve(result)
        })
    })
}



