import SearchStore from './SearchStore.js';
import { getMapObjects } from '../api.js';

export default {

    search: function(){
      SearchStore.show = true;
      this.fetchData();
    },

    fetchData: function(){
      SearchStore.result = [];

      if (!SearchStore.query){
        return;
      }

      SearchStore.busy = true;



    },

    clear: function(){
      SearchStore.result = [];
      SearchStore.show = false;
    }
};
