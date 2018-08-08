const SQSClass = require('./sqs/sqs');

// ENV VARS
const que_url =  process.env.QUE_URL;
const sqsEndpoint = process.env.SQS_ENDPOINT;

let mockMessage = {
    "Type" : "Notification",
    "MessageId" : "cb21a901-5504-5eee-8ca7-fae1810379c1",
    "TopicArn" : "arn:aws:sns:us-east-1:359864072703:clip-slugs-sns-dev",
    "Message" : "{\"videoSlugs\":[\"SmoggyEnergeticPeafowlBudBlast\",\"SparklyCarelessHorseFreakinStinkin\",\"SmilingRelatedWeaselEleGiggle\",\"CharmingTenderLaptopTriHard\",\"ResoluteSnappyWormWTRuck\",\"JazzyWildEggnogKappaPride\",\"NiceKitschyTriangleGingerPower\",\"AthleticHardHeronYouWHY\",\"SwissBombasticChimpanzeePunchTrees\"],\"videoDescription\":\"I'll be gentle | !prime @DrDisRespect\\n https://clips.twitch.tv/SmoggyEnergeticPeafowlBudBlast?tt_medium=clips_api\\u0026tt_content=url \\nRIP Shroud!!  https://clips.twitch.tv/SparklyCarelessHorseFreakinStinkin?tt_medium=clips_api\\u0026tt_content=url \\nOKAY?! https://clips.twitch.tv/SmilingRelatedWeaselEleGiggle?tt_medium=clips_api\\u0026tt_content=url \\nDrDisreporpoise  WARNING: VOLUME https://clips.twitch.tv/CharmingTenderLaptopTriHard?tt_medium=clips_api\\u0026tt_content=url \\nI'll be gentle | !prime @DrDisRespect https://clips.twitch.tv/ResoluteSnappyWormWTRuck?tt_medium=clips_api\\u0026tt_content=url \\nDoc finds out if it's Friday or Monday https://clips.twitch.tv/JazzyWildEggnogKappaPride?tt_medium=clips_api\\u0026tt_content=url \\nDrop dead gorgeous | !prime @DrDisRespect https://clips.twitch.tv/NiceKitschyTriangleGingerPower?tt_medium=clips_api\\u0026tt_content=url \\nBetter | !prime @DrDisRespect https://clips.twitch.tv/AthleticHardHeronYouWHY?tt_medium=clips_api\\u0026tt_content=url \\nLend me your strength brother!! https://clips.twitch.tv/SwissBombasticChimpanzeePunchTrees?tt_medium=clips_api\\u0026tt_content=url \\nTotal lengh of the video should be 4.466666666666667 long\",\"channelName\":\"drdisrespectlive\"}",
    "Timestamp" : "2018-08-05T10:30:31.209Z",
    "SignatureVersion" : "1",
    "Signature" : "cIu/d4XFTqh8hN1/Epsl+WkyZ3/wGraRIQzMwihb6SPB/oyPeu0tZuFR8cKJ5CS6EmK5nYrpC/SJzSgB+NqV3fpC1CxPMQQ8xKdFRqzkXtpDP36wdwmgdFS82EMsVONIMLGkpcX3xViWstKU6FsJG0gHYNtVf5xEdhq93LXh710bxwaN9e2oYXWf13f+OIafITFm+p0BYz3TECiOW8SwdxbY5hQ8zWcgdzUz6o8/maoYoZtTjmIPVyRLHlm5pJEajtj7NrMP51o//qQNcw+Umb8hqNLTq8p2FQHuwZ4L/20MkuILFAEZ4GRveUnF4mInxnxlCITHnmJ8VaF5r1h8dQ==",
    "SigningCertURL" : "https://sns.us-east-1.amazonaws.com/SimpleNotificationService-eaea6120e66ea12e88dcd8bcbddca752.pem",
    "UnsubscribeURL" : "https://sns.us-east-1.amazonaws.com/?Action=Unsubscribe&SubscriptionArn=arn:aws:sns:us-east-1:359864072703:clip-slugs-sns-dev:01d3fb4c-d773-47b3-8b3a-124138bc2c13"
}


main = async () => {
    console.log("mock data started")
    const sqs = new SQSClass(que_url, sqsEndpoint)
    try {
        console.log("Sending Mock Message to dead-letter")
        await sqs.sendMessage(JSON.stringify(mockMessage))
        console.log("mock message send to dead letter")
        process.exit(0)
    } catch(e) {
        console.log(e)
        process.exit(1)
    }
}

main();