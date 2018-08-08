let adaptQueToDb = (queUrl,dbTableName, message) =>{
    let queName = queUrl.split("/")
    queName = queName[queName.length -1]
    let dbItem = {
            queName: queName,
            tableName: dbTableName,
            timestamp: message.Timestamp,
            messageId: message.MessageId,
            message: message.Message
    }
    return dbItem
}

exports.adaptQueToDb = adaptQueToDb