import ws from '../service/ws.js';

export default {
    data: function(){
        return {
            info: null
        };
    },
    methods: {
        infoListener: function(info){
            this.info = info;
        }
    },
    computed: {
        lagColor: function(){
            if (this.info.max_lag > 0.8)
                return "orange";
            else if (this.info.max_lag > 1.2)
                return "red";
            else
                return "green";
        },
        time: function() {
            const min = Math.floor((this.info.time % 1000) / 1000 * 60);
            return Math.floor(this.info.time/1000) + ":" + (min >= 10 ? min : "0" + min);
        }
    },
    created: function() {
        // bind infoListener to this
        this.infoListener = this.infoListener.bind(this);
        ws.addListener("minetest-info", this.infoListener);
    },
    beforeUnmount: function() {
        ws.removeListener("minetest-info", this.infoListener);
    },
    template: /*html*/`
    <div v-if="info">
        <span v-if="info.players">
            <span class="fa fa-users"></span> {{ info.players.length }}
        </span>
        <span class="fa fa-wifi" v-bind:style="{ 'color': lagColor }"></span> {{ parseInt(info.max_lag*1000) }} ms
        <span class="fa fa-clock"></span> {{ time }}
        <span v-if="info.time < 5500 || info.time > 19000" class="fa fa-moon" style="color: blue;"></span>
        <span v-else class="fa fa-sun" style="color: orange;"></span>
    </div>
    `
};