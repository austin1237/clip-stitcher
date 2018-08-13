let validateEvent = (event) =>{

    if (!Array.isArray(event.Records) || event.Records.length === 0){
        throw new Error("invalid records found in message");
    }

    let record = event.Records[0];
    if (typeof record.body !== 'string'  || record.body.length === 0){
        throw new Error("invalid messsage body found in message");
    }

    if (!record.body.includes('"Timestamp" :')){
        throw new Error("Timestamp not included in body of the message");
    }

    if (!record.body.includes('"Message" :')){
        throw new Error("message not included in body of the message");
    }

    if (typeof record.messageId !== 'string'  || record.messageId.length ===  0){
        throw new Error("invalid messageId found in message");
    }
}

exports.validateEvent = validateEvent