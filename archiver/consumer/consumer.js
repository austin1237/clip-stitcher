const AWS = require('aws-sdk');

class consumer {
    constructor(queueUrl,  customEndpoint){
        let sqsConfig = {apiVersion: '2012-11-05'};
        if (customEndpoint) {
            sqsConfig.endpoint = new AWS.Endpoint(customEndpoint);
        }
        this._queueUrl = queueUrl
        this._sqsClient = new AWS.SQS(sqsConfig);
        this._retryCount = 60
    }

    async getMessage(){
        const params = {
            QueueUrl: this._queueUrl,
            WaitTimeSeconds: 20
        }
        let err = {}
        for (let retry = 0; retry < this._retryCount; retry++) {
            try{
                const data = await this._sqsClient.receiveMessage(params).promise();
                console.log('we got the data')
                if (!data || !Array.isArray(data.Messages) || !data.Messages.length){
                    const err = new Error(`no messages returned from ${this._queueUrl}`)
                    throw err
                }
                console.log("parsing message data")
                const message = JSON.parse(data.Messages[0].Body)
                message.receiptHandle = data.Messages[0].ReceiptHandle
                return Promise.resolve(message)
            }catch (e){
                console.log("Error getting message retrying")
                err = e;
                await sleep(1000)
            }
        }
        return Promise.reject(err) 
    }

    async deleteMessage(message){
        const params = {
            QueueUrl: this._queueUrl,
            ReceiptHandle: message.receiptHandle
        }
        let err = {};
        for (let retry = 0; retry < this._retryCount; retry++) {
            try{
                this._sqsClient.deleteMessage(params).promise();
                return Promise.resolve();
            }catch (e){
                err = e;
                console.log("Error deleting message retrying")
                await sleep(1000)
            }
        }
        return Promise.reject(err) 
    }
}

function sleep(ms) {
    return new Promise(resolve => setTimeout(resolve, ms));
}

module.exports = consumer