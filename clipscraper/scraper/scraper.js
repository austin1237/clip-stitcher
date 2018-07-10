const path = require('path')
const fs = require ('fs')
const tar = require('tar')
const puppeteer = require('puppeteer')

const chromeVer = "67.0.3361.0"
const compressedChromePath = `./chrome/chrome-${chromeVer}.tar.gz`
const tmpPath = path.join(path.sep, 'tmp');

class Scraper {
    constructor(browser) {
        if (typeof browser === 'undefined') {
            throw new Error('Cannot be called directly due to the async nature of puppeteer');
        }
        this._browser = browser
    }

    static async build (){
        try{
            await setupLocalChrome();
            const browser = await puppeteer.launch({
                executablePath: tmpPath + `/headless_shell`,
                headless:true,
                slowMo: 10,
                args: [
                    '--no-sandbox', 
                    '--single-process', 
                    '--disable-gpu',
                ]
            });
            return Promise.resolve(new Scraper(browser));
        } catch(e){
            return Promise.reject(e)
        }
    }

    async getVidSrcFromUrl (url) {
       const page = await this._goToUrl(url)
       return this._evaluateSrcOnPage(page, 10)
    }

    async _goToUrl (slug){
        const page = await this._browser.newPage();
        const timeout = 30000;
        const networkIdleTimeout = 1100;
        const url = `https://player.twitch.tv/?branding=false&clip=${slug}&externalfullscreen=true&muted=true&origin=https%3A%2F%2Fclips.twitch.tv&player=clips-viewing&playsinline=true`
        await Promise.all([
          page.waitForNavigation({ timeout, waitUntil: 'load' }),
          page.goto(url)
        ]);
    
        return Promise.resolve(page)
    }

    async _evaluateSrcOnPage (page, retriesLeft){
        const src = await page.evaluate((selector) => {
            const anchors_node_list = document.querySelectorAll(selector);
            const anchors = [...anchors_node_list];
            return anchors.map(link => link.src);
        }, 'video');
    
        if ((!src[0] || src[0].length) && retriesLeft >0) {
            return this._evaluateSrcOnPage(page, retriesLeft - 1)
        }
        return Promise.resolve(src[0])
    }
}

// To avoid the 50mb lambda upload limit chrome is unzipped during execution into /tmp as /tmp/headless_shell
// The default chromium install that came with puppeteer didn't work so I used the one found here instead https://github.com/sambaiz/puppeteer-lambda-starter-kit
const setupLocalChrome = () => {
    return new Promise((resolve, reject) => {
      fs.createReadStream(compressedChromePath)
      .on('error', (err) => reject(err))
      .pipe(tar.x({
        C: tmpPath,
      }))
      .on('error', (err) => reject(err))
      .on('end', () => resolve());
    });
};


module.exports = Scraper;