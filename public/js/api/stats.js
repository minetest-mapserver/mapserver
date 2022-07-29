
export const get = () => fetch("api/stats").then(r => r.json());
