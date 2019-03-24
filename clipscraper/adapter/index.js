let adaptEventToMessage = (event, tableName) =>{
    let record = event.Records[0]
    
    let body = record.body
    // removes trailing commas from a json object
    let regex = /\,(?=\s*?[\}\]])/g;
    body = body.replace("\'", "")
    body = body.replace("\n", "")
    body = body.replace(regex, ''); 
    message = JSON.parse(body);
    return message
}

exports.adaptEventToMessage = adaptEventToMessage