const eventAdapter = require("../adapter/eventAdapter.js");
const assert = require("chai").assert;

describe('messageAdapter', () => {

    describe('#adaptEventToMessage()', () =>{
        it("should transform the event properly", ()=>{
            let mockDbTable = "fakeTable"
            let mockEvent = { Records: 
                [ { messageId: 'c6cf86d7-45c0-4a6e-ba44-ca5e4643b69a',
                receiptHandle: 'AQEB3be/YeHxFhqKb6XawKZWMZjaHRMKMD4pMrloYDH6y8hLIIIXC2cyiYmpXWDmNG7DmzeN4DE14vCzRt7lAJvqXZu6r4jk2nf6K+cXj3TaxpYkXVGECfxEOIdc8s8wZCsALx2pEFxw0WfMxxxVpREpLNhFQuCPlqjMkk9/GrIqPnPq6oXgvqPVECqqe4ViJQs3n2s0e5rLoCndIIbKLr8qIDprimSoP1KlrAHAEQxisgnMxB4YpV72D6d12LW3py6ubEu7cnEjZgIhycojpZ/jsVD6vqGZY3clOO7mmmJHOj3GefLaO7g5uw8hPN5RK8glRX98B30n+Fh8nNoh3RCp4C/FeBS7ZMJAMI/dcL1Yre8aUOBt/xhoMahDdh5LRWz+urNWjajiENb7cJbZ3SiMcjsnRxdQmwC66OvQCd3fWTA=',
                body: '{\n "Type" : "Notification",\n "MessageId" : "6ed7efc3-4399-58ab-adeb-bc278b37cda6",\n "TopicArn" : "arn:aws:sns:us-east-1:359864072703:clip-links-sns-dev",\n "Message" : "{\\"videoLinks\\":[\\"https://link.mp4\\",\\"https://link.mp4\\",\\"https://link.mp4\\",\\"https://link.mp4\\",\\"https://link.mp4\\",\\"https://link.mp4\\",\\"https://link.mp4\\",\\"https://link.mp4\\",\\"https://link.mp4\\"],\\"videoDescription\\":\\" \\\\nTotal lengh of the video should be 5.633333333333333 long\\",\\"channelName\\":\\"testChannel\\"}",\n "Timestamp" : "2018-08-08T10:31:10.679Z",\n }',
                messageAttributes: {},
                md5OfBody: '1dfbed8850b81fd9340863ffdebc5812',
                eventSource: 'aws:sqs',
                eventSourceARN: 'arn:aws:sqs:us-east-1:359864072703:clip-links-sqs-dev-dead-letter',
                awsRegion: 'us-east-1' } ] 
            }
            let result = eventAdapter.adaptEventToMessage(mockEvent, mockDbTable)
            assert.equal(mockEvent.Records[0].messageId, result.messageId)
            assert.equal("2018-08-08T10:31:10.679Z", result.timestamp)
            assert.equal(mockEvent.Message, result.body)
            assert.equal(mockDbTable, result.tableName)
            assert.equal("clip-links-sqs-dev-dead-letter", result.queName)
        })
    })
})

