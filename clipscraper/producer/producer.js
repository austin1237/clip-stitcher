const AWS = require('aws-sdk');



class consumer {
    constructor(queueUrl,  snsClient){
        this._queueARN = queueARN
        this._sqsClient = snsClient
    }

    static async build (queueName, queEndpoint){
        let snsConfig = {apiVersion: '2010-03-31'};
        if (queEndpoint !== "") {
            snsConfig.endpoint = new AWS.Endpoint(queEndpoint);
            const snsClient = new AWS.SQS(sqsConfig);
            const queueARN = `${queEndpoint}/queue/${queueName}`
            return Promise.resolve(new consumer(queueUrl, sqsClient))
        }
    }
}