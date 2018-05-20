const AWS = require('aws-sdk');

class consumer {
    constructor(queueUrl,  sqsClient){
        this._queueUrl = queueUrl
        this._sqsClient = sqsClient
    }

    static async build (queueName, queEndpoint){
        let sqsConfig = {apiVersion: '2012-11-05'};
        if (queEndpoint !== "") {
            sqsConfig.endpoint = new AWS.Endpoint(queEndpoint);
            const sqsClient = new AWS.SQS(sqsConfig);
            const queueUrl = `${queEndpoint}/queue/${queueName}`
            return Promise.resolve(new consumer(queueUrl, sqsClient))
        }
        try{
            const sqsClient = new AWS.SQS(sqsConfig);
            const query = {QueueName: queueName}
            const data = await sqsClient.getQueueUrl(query).promise();
            return Promise.resolve(new consumer(data.QueueUrl, sqsClient))
        } catch (e) {
            return Promise.reject(e)
        }
    }

    async getMessage(){
        const query = {
            QueueUrl: this._queueUrl,
        }
        try{
            const data = await this._sqsClient.receiveMessage(query).promise();
            const snsWrapper = JSON.parse(data.Messages[0].Body)
            let message ={}
            message.receiptHandle = data.Messages[0].ReceiptHandle
            message.body = JSON.parse(snsWrapper.Message)
            return Promise.resolve(message)
        }catch (e){
            return Promise.reject(e)
        }   
    }

    deleteMessage(message){
        const query = {
            QueueUrl: this._queueUrl,
            ReceiptHandle: message.ReceiptHandle
        }
        return this._sqsClient.deleteMessage(query).promise();
    }
}

module.exports = consumer