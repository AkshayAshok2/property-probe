// create pup instance
const puppeteer = require('puppeteer');

(async () => {
  //set scrape url
  const url = 'https://alachua.realforeclose.com/index.cfm?zaction=AUCTION&zmethod=PREVIEW&AUCTIONDATE=02/07/2023';
  //launch and wait for page to load
  const browser = await puppeteer.launch();
  const page = await browser.newPage();
  //set useragent and load url
  await page.setUserAgent('Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3');
  await page.goto(url);
  //wait for auctions to show in html
  await page.waitForSelector('.AUCTION_ITEM.PREVIEW');
  const html = await page.content();
  //grab auctions with cheerio
  const cheerio = require('cheerio');
  const $ = cheerio.load(html);

  const auctionItems = $('.AUCTION_ITEM.PREVIEW');

  // Loop through each auction item and print out the details
  for (let i = 0; i < auctionItems.length; i++) {
    const auctionItem = auctionItems.eq(i);
    // FOR BACKEND TO ADD
    // by looping through the rows of the detail table we grab all the information needed, this information needs to be added to a data structure that lets us refrence each property as an object so we can then use the scraped data
    const tableRows = auctionItem.find('table tr');
    tableRows.each((index, row) => {
      const label = $(row).find('th').text().trim();
      const data = $(row).find('td').text().trim();
      console.log(label, data);
    });
    //space out each auction item
    console.log("\n");
  }
  // Add in to add to a text file to parse from node script-file.js >properties.txt 2>error-file.txt

  await browser.close();
})();
