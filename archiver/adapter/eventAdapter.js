let adaptEventToMessage = (event, tableName) =>{
    let record = event.Records[0]
    
    let body = record.body
    // removes trailing commas from a json object
    let regex = /\,(?=\s*?[\}\]])/g;
    body = body.replace("\'", "")
    body = body.replace("\n", "")
    body = body.replace(regex, ''); 
    body = JSON.parse(body)
    let queArr= record.eventSourceARN.split(":")
    let queName = queArr[queArr.length -1]
    let message = {
        queName: queName,
        message: body.Message,
        timestamp:  body.Timestamp,
        messageId: record.messageId,
        tableName: tableName,
    }
    return message
}

exports.adaptEventToMessage = adaptEventToMessage