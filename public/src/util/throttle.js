// calls func at most once per wait ms (leading and trailing), with the latest arguments
export default function throttle(func, wait) {
  let timeout = null;
  let lastArgs = null;
  return function(...args) {
    if (timeout) {
      lastArgs = args;
      return;
    }
    func.apply(this, args);
    timeout = setTimeout(() => {
      timeout = null;
      if (lastArgs) {
        const a = lastArgs;
        lastArgs = null;
        func.apply(this, a);
      }
    }, wait);
  };
}
