
export const get = () => fetch("api/config").then(r => r.json());
