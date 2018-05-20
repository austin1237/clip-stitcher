const ScraperClass = require('./scraper/scraper');
const ConsumerClass = require('./consumer/consumer');

// ENV vars
const consumerName = process.env.CONSUMER_NAME;
const consumerEndpoint = process.env.CONSUMER_ENDPOINT;
const producerName =  process.env.PRODUCER_NAME;
const producerEndpoint = process.env.PRODUCER_ENDPOINT;

main = async () => {
    
    initPromises = [
      ConsumerClass.build(consumerName, consumerEndpoint),
      ScraperClass.build()
    ]
    const results = await Promise.all(initPromises)
    const consumer = results[0]
    const scraper = results[1]
    const message = await consumer.getMessage()
    let srcPromises = [];
    message.body.videoSlugs.forEach((slug)=> {
        srcPromises.push(scraper.getVidSrcFromUrl(slug))
    });
    srcs = await Promise.all(srcPromises)
    console.log(srcs)
}

exports.handler = async(event, context) => {
   console.log("testinz")
   await main()
   console.log("all day")
}

