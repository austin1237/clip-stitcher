const axios = require('axios');
const zlib = require('zlib');
const testLog = require('testLog');

// env vars
const appENV = process.env.APP_ENV;
const logDnaKey = process.env.LOG_DNA_KEY;

main = async (cloudWatchLog) => {
    // validate here
    let payload = transformCWToLD(cloudWatchLog);
    return sendLogsOff(payload, logDnaKey);
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

let sendLogsOff = async (payload, logDnaKey) => {
    const axiosClient = axios.create({
        baseURL: 'https://logs.logdna.com/logs/ingest?hostname=aws',
        timeout: 10000,
        headers: {'apikey': logDnaKey}
    });
    return axiosClient.post("", payload)
}

let transformCWToLD = (logs) => {
    let transformedLd= {};
    let ldLines = [];
    let cwLines = logs.logEvents;
    let appStrArr = logs.logGroup.split("/");
    let appName = appStrArr[appStrArr.length -1];
    cwLines.forEach(function(cwLine) {
        // dunno what to do with cwLine.timestamp
       let level = detectLevel(cwLine.message);
       let ldLine = {
            line: cwLine.message,
            app: appName,
            level: level,
            meta: {}
        }
        ldLines.push(ldLine);
    });
    transformedLd.lines = ldLines;
    return transformedLd
}

let detectLevel = (logText) => {
    let level = "INFO";
    let logTextLower = logText.toLowerCase();
    if (logTextLower.includes("error")){
        level = "ERROR";
    }

    if (logTextLower.includes("retry")){
        level = "WARN";
    }

    return level
}