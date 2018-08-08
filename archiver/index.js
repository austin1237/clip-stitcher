const ConsumerClass = require('./consumer/consumer');
const DBClass = require('./db/db');
const messageAdapter = require("./adapter/messageAdapter.js");
const validator = require("./validator/messageValidator.js");

// env vars
const consumerUrl = process.env.CONSUMER_URL;
const consumerEndpoint = process.env.CONSUMER_ENDPOINT;
const dbTableName = process.env.DB_TABLE;
const dbEndPoint = process.env.DB_ENDPOINT;

main = async () => {
    console.log("getting message from dead letter")
    const consumer = new ConsumerClass(consumerUrl, consumerEndpoint);
    const db = new DBClass(dbEndPoint);
    try {
        let message = await consumer.getMessage()
        console.log(`message recevied from dead letter`);
        validator.validateMessage(message)
        let dbItem = messageAdapter.adaptQueToDb(consumerUrl, dbTableName, message)
        await db.saveMessage(dbItem)
        console.log("message saved in db")
        await consumer.deleteMessage(message)
        console.log("message deleted")
    } catch(e) {
        console.log("exiting due to error")
        console.log(e)
        process.exit(1)
    }
}

exports.handler = async(event, context) => {
    await main()
 }