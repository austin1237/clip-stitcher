let validateMessage = (message) =>{
    if (typeof message.Message !== 'string'  || message.Message.length === 0){
        throw "invalid messsage body found in message";
    }

    if (typeof message.MessageId !== 'string'  || message.MessageId.length ===  0){
        throw new Error("invalid messageId found in message");
    }

    if (typeof message.Timestamp !== 'string'  || message.Timestamp.length === 0){
        throw "invalid timestamp found in message";
    }

}

exports.validateMessage = validateMessage