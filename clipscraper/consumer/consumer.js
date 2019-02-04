const AWS = require('aws-sdk');

class consumer {
    constructor(queueUrl,  customEndpoint, waitTime, retryCount){
        let sqsConfig = {apiVersion: '2012-11-05'};
        if (customEndpoint) {
            sqsConfig.endpoint = new AWS.Endpoint(customEndpoint);
        }
        this._queueUrl = queueUrl
        this._sqsClient = new AWS.SQS(sqsConfig);
        this._retryCount = retryCount
        this._waitTime = waitTime
    }

    async getMessage(){
        const params = {
            QueueUrl: this._queueUrl,
            WaitTimeSeconds: this._waitTime
        }
        let err = {}
        for (let retry = 0; retry < this._retryCount; retry++) {
            try{
                const data = await this._sqsClient.receiveMessage(params).promise();
                if (!data || !Array.isArray(data.Messages) || !data.Messages.length){
                    
                    const noMsgErr = new Error(`no messages returned from ${this._queueUrl}`);
                    throw noMsgErr;
                }
                const snsWrapper = JSON.parse(data.Messages[0].Body)
                let message = JSON.parse(snsWrapper.Message)
                message.receiptHandle = data.Messages[0].ReceiptHandle
                return Promise.resolve(message)
            }catch (e){
                err = e;
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
                await this._sqsClient.deleteMessage(params).promise();
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

module.exports = consumer