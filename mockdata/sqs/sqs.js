const AWS = require('aws-sdk');

class sqs {
    constructor(queueUrl,  customEndpoint){
        let sqsConfig = {apiVersion: '2012-11-05'};
        if (customEndpoint) {
            sqsConfig.region = 'us-east-1';
            sqsConfig.endpoint = new AWS.Endpoint(customEndpoint);
            sqsConfig.accessKeyId = "xxxxxxxxxxxxxx"
            sqsConfig.secretAccessKey = "xxxxxxxxxxxxxx"
        }
        this._queueUrl = queueUrl
        this._sqsClient = new AWS.SQS(sqsConfig);
        this._retryCount = 60
    }

    async sendMessage(message){
        const params = {
            QueueUrl: this._queueUrl,
            MessageBody: message,
            DelaySeconds: 0
        }
        let err = {};
        for (let retry = 0; retry < this._retryCount; retry++) {
            try{
                await this._sqsClient.sendMessage(params).promise();
                return Promise.resolve();
            }catch (e){
                err = e;
                console.log("Error publishing message retrying")
                await sleep(1000)
            }
        }
        return Promise.reject(err) 
    }
}

function sleep(ms) {
    return new Promise(resolve => setTimeout(resolve, ms));
}


module.exports = sqs