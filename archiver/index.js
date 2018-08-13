const DBClass = require('./db/db');
const eventAdapter = require("./adapter/eventAdapter.js");
const eventValidator = require("./validator/eventValidator.js");
const testEvent = require("./testEvent.js");

// ENV vars
const dbTableName = process.env.DB_TABLE;
const dbEndPoint = process.env.DB_ENDPOINT;
const appENV = process.env.APP_ENV;

main = async (event) => {
    const db = new DBClass(dbEndPoint);
    try {
        console.log("validating event")
        eventValidator.validateEvent(event)
        console.log("adapting event")
        let dbItem = eventAdapter.adaptEventToMessage(event, dbTableName)
        console.log("saving item to db")
        await db.saveMessage(dbItem)
        console.log("item saved in db")
    } catch(e) {
        console.log("exiting due to error")
        console.log(e)
        process.exit(1)
    }
}

exports.handler = async(event, context) => {
    if (appENV === "local"){
        console.log("local mode enabled")
        await main(testEvent)
    }else{
        await main(event)
    }
}

