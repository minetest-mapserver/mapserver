import { ParsedMapblock, MAPBLOCK_VOLUME } from '@buckaroobanzay/luanti-viewer';
import { getViewChunk } from '../api.js';

// edge length of a chunk in mapblocks, see types/chunk.go
const CHUNK_SIZE = 5;
// number of chunks kept in memory, the viewer keeps its own copy of the loaded mapblocks
const MAX_CHUNKS = 64;

async function inflate(bytes){
  const stream = new Blob([bytes]).stream().pipeThrough(new DecompressionStream("deflate"));
  return new Uint8Array(await new Response(stream).arrayBuffer());
}

// mapdata: base64 of the zlib-compressed payload, see util/serialize.go
// 4096 * uint16 (big endian) content ids, 4096 * param1, 4096 * param2
async function parseViewBlock(block){
  const data = await inflate(Uint8Array.from(atob(block.mapdata), c => c.charCodeAt(0)));
  if (data.length != MAPBLOCK_VOLUME * 4){
    throw new Error(`unexpected mapdata length: ${data.length}`);
  }

  const view = new DataView(data.buffer, data.byteOffset, data.byteLength);
  const nodeIds = new Uint16Array(MAPBLOCK_VOLUME);
  for (let i = 0; i < MAPBLOCK_VOLUME; i++){
    nodeIds[i] = view.getUint16(i * 2, false);
  }

  const nodeNames = new Map();
  Object.entries(block.blockmapping || {}).forEach(([id, name]) => nodeNames.set(+id, name));
  // ids without a mapping entry are rendered as unknown nodes
  for (const id of nodeIds){
    if (!nodeNames.has(id)){
      nodeNames.set(id, "unknown");
    }
  }

  return new ParsedMapblock(
    nodeIds,
    data.slice(MAPBLOCK_VOLUME * 2, MAPBLOCK_VOLUME * 3),
    data.slice(MAPBLOCK_VOLUME * 3, MAPBLOCK_VOLUME * 4),
    nodeNames
  );
}

const blockKey = pos => `${pos.x}/${pos.y}/${pos.z}`;

async function loadChunk(x, y, z){
  const blocks = new Map();
  const list = await getViewChunk(x, y, z);
  await Promise.all((list || []).map(async block => {
    blocks.set(blockKey(block.block_pos), await parseViewBlock(block));
  }));
  return blocks;
}

// loads mapblocks in batches (chunks) from the backend
export function createViewChunkLoader(){
  // chunk key -> promise of (mapblock key -> ParsedMapblock), in least-recently-used order
  const chunks = new Map();

  return async function getMapblock(pos){
    const x = Math.floor(pos.x / CHUNK_SIZE);
    const y = Math.floor(pos.y / CHUNK_SIZE);
    const z = Math.floor(pos.z / CHUNK_SIZE);
    const key = `${x}/${y}/${z}`;

    let chunk = chunks.get(key);
    chunks.delete(key);
    if (!chunk){
      chunk = loadChunk(x, y, z);
      // don't keep failed requests, they are retried with the next mapblock
      chunk.catch(() => chunks.get(key) === chunk && chunks.delete(key));
    }
    chunks.set(key, chunk);

    if (chunks.size > MAX_CHUNKS){
      chunks.delete(chunks.keys().next().value);
    }

    return (await chunk).get(blockKey(pos)) || null;
  };
}
