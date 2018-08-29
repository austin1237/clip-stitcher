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
    if (logTextLower.includes("error ") || logTextLower.includes("err ")){
        level = "ERROR";
    }

    if (logTextLower.includes("retry")){
        level = "WARN";
    }

    return level
}

exports.transformCWToLD = transformCWToLD
exports.dectectLevel = detectLevel