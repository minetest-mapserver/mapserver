import { getViewMediaBatch } from '../api.js';

// max number of files per request, see web/webview.go
const MAX_BATCH = 128;
// time in milliseconds to collect requests before they are sent
const BATCH_DELAY = 20;

// filename -> promise of Blob or null, shared across views
const cache = new Map();
// filename -> { resolve, reject } of the requests that are not sent yet
let pending = new Map();
let timer = null;

async function flush(){
  clearTimeout(timer);
  timer = null;
  const batch = pending;
  pending = new Map();

  try {
    const list = await getViewMediaBatch([...batch.keys()]);
    for (const entry of list || []){
      const request = batch.get(entry.filename);
      if (request){
        // atob accepts the unpadded base64 of the backend
        request.resolve(new Blob([Uint8Array.from(atob(entry.data), c => c.charCodeAt(0))]));
        batch.delete(entry.filename);
      }
    }
    // the remaining files do not exist
    batch.forEach(request => request.resolve(null));
  } catch (e){
    // don't keep failed requests, they are retried with the next request of the file
    batch.forEach((request, filename) => {
      cache.delete(filename);
      request.reject(e);
    });
  }
}

// texture or model as blob, null if it does not exist
// requests are collected and sent in batches to the backend
export function getViewMedia(filename){
  let promise = cache.get(filename);
  if (!promise){
    promise = new Promise((resolve, reject) => pending.set(filename, { resolve, reject }));
    cache.set(filename, promise);

    if (pending.size >= MAX_BATCH){
      flush();
    } else if (timer === null){
      timer = setTimeout(flush, BATCH_DELAY);
    }
  }
  return promise;
}
