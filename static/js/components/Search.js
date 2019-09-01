
export default {
  oncreate(vnode){
    console.log("oncreate", vnode);
  },

  view(vnode){
    return m("div", vnode.attrs.query)
  }
}
