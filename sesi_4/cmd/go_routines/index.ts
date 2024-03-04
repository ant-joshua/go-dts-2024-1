async function crawlGoogle() {
  return new Promise((resolve, reject) => {
    setTimeout(() => {
      console.log("Crawling Google");
      resolve("berhasil");
    }, 1000);
  });
}

async function crawlBing() {
  return new Promise((resolve, reject) => {
    setTimeout(() => {
      console.log("Crawling Bing");
      resolve("berhasil crawl bing");
    }, 1000);
  });
}

async function crawlYahoo() {
  return new Promise((resolve, reject) => {
    setTimeout(() => {
      console.log("Crawling Yahoo");
      resolve("berhasil crawl yahoo");
    }, 1000);
  });
}

async function doAction() {
  const t0 = performance.now();
  //   const google = await crawlGoogle();

  //   const bing = await crawlBing();

  //   const yahoo = await crawlYahoo();

  await Promise.all([crawlGoogle(), crawlBing(), crawlYahoo()]);

  const t1 = performance.now();

  console.log(`Call to doSomething took ${t1 - t0} milliseconds.`);
}

doAction();
